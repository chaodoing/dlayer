package generator

import (
	"strings"

	"gorm.io/gen"
	"gorm.io/gen/field"
)

// FieldMapping 描述单个字段名到 Go 类型的映射。
type FieldMapping struct {
	// FieldName 是数据库列名，例如 deleted_at。
	FieldName string `json:"field_name" yaml:"field_name" toml:"field_name"`
	// GoType 是生成到模型结构体中的 Go 类型，例如 gorm.DeletedAt。
	GoType string `json:"go_type" yaml:"go_type" toml:"go_type"`
	// ImportPath 是 GoType 需要的导入路径；gorm.io/gorm 等已在模板中内置时可省略。
	ImportPath string `json:"import_path" yaml:"import_path" toml:"import_path"`
	// GenType 是查询层 field.NewXxx 使用的类型名；为空时由 gorm.io/gen 自动推导。
	GenType string `json:"gen_type" yaml:"gen_type" toml:"gen_type"`
}

// defaultFieldMappings 返回内置字段映射，可被配置文件覆盖。
func defaultFieldMappings() map[string]FieldMapping {
	return map[string]FieldMapping{
		"deleted_at": {
			FieldName:  "deleted_at",
			GoType:     "gorm.DeletedAt",
			ImportPath: "gorm.io/gorm",
		},
	}
}

// effectiveFieldMappings 合并默认映射与用户配置，用户配置优先。
func effectiveFieldMappings(cfg Config) map[string]FieldMapping {
	merged := defaultFieldMappings()
	for key, mapping := range cfg.FieldMappings {
		normalized := normalizeFieldKey(key)
		if mapping.FieldName == "" {
			mapping.FieldName = normalized
		} else {
			mapping.FieldName = normalizeFieldKey(mapping.FieldName)
		}
		merged[normalized] = mapping
	}
	return merged
}

// fieldMappingModelOpt 根据字段名映射覆盖模型字段类型。
func fieldMappingModelOpt(cfg Config) gen.ModelOpt {
	mappings := effectiveFieldMappings(cfg)
	return gen.FieldModify(func(f gen.Field) gen.Field {
		if f == nil {
			return f
		}
		mapping, ok := mappings[normalizeFieldKey(f.ColumnName)]
		if !ok {
			return f
		}
		f.Type = mapping.GoType
		if mapping.GenType != "" {
			f.CustomGenType = mapping.GenType
		}
		return f
	})
}

// nullablePointerModelOpt 按 gorm not null 标签决定字段是否使用指针类型。
func nullablePointerModelOpt() gen.ModelOpt {
	return gen.FieldModify(func(f gen.Field) gen.Field {
		if f == nil || f.Relation != nil || isFixedModelFieldType(f.Type) {
			return f
		}
		baseType := strings.TrimPrefix(f.Type, "*")
		if !canUsePointer(baseType) {
			return f
		}
		if gormFieldNotNull(f) {
			f.Type = baseType
			return f
		}
		f.Type = "*" + baseType
		return f
	})
}

// gormFieldNotNull 判断字段是否应按非空列生成值类型。
func gormFieldNotNull(f gen.Field) bool {
	if f == nil {
		return false
	}
	if _, ok := f.GORMTag[field.TagKeyGormPrimaryKey]; ok {
		return true
	}
	_, ok := f.GORMTag[field.TagKeyGormNotNull]
	return ok
}

// isFixedModelFieldType 判断是否为不应再套指针的固定模型类型。
func isFixedModelFieldType(goType string) bool {
	switch strings.TrimPrefix(goType, "*") {
	case "gorm.DeletedAt":
		return true
	default:
		return false
	}
}

// fieldMappingImportPaths 收集字段映射所需的额外 import 路径。
func fieldMappingImportPaths(cfg Config) []string {
	seen := make(map[string]struct{})
	paths := make([]string, 0, len(cfg.FieldMappings))
	for _, mapping := range effectiveFieldMappings(cfg) {
		path := strings.TrimSpace(mapping.ImportPath)
		if path == "" {
			continue
		}
		if _, ok := seen[path]; ok {
			continue
		}
		seen[path] = struct{}{}
		paths = append(paths, path)
	}
	return paths
}

// normalizeFieldMappings 统一字段映射 key 与 FieldName。
func normalizeFieldMappings(mappings map[string]FieldMapping) {
	for key, mapping := range mappings {
		normalized := normalizeFieldKey(key)
		if mapping.FieldName == "" {
			mapping.FieldName = normalized
		} else {
			mapping.FieldName = normalizeFieldKey(mapping.FieldName)
		}
		if normalized != key {
			delete(mappings, key)
		}
		mappings[normalized] = mapping
	}
}

// normalizeFieldKey 统一数据库列名大小写与空白。
func normalizeFieldKey(value string) string {
	return strings.ToLower(strings.TrimSpace(value))
}
