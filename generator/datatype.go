package generator

import (
	"strings"

	"gorm.io/gorm"
)

// goTypeForColumn 根据数据库列类型推导验证结构中使用的 Go 类型。
func goTypeForColumn(cfg Config, column gorm.ColumnType) string {
	if mapping, ok := lookupCustomTypeMapping(cfg, column); ok {
		return mapping.GoType
	}

	dbType := normalizeDBType(column.DatabaseTypeName())
	columnType, _ := column.ColumnType()
	columnType = strings.ToLower(columnType)
	switch {
	case strings.Contains(dbType, "json"):
		return "datatypes.JSON"
	case strings.Contains(dbType, "uuid"):
		return "string"
	case isArrayDBType(dbType, columnType):
		return "[]string"
	case strings.Contains(dbType, "tinyint") && strings.Contains(columnType, "(1)"):
		return "bool"
	case strings.Contains(dbType, "tinyint"):
		return unsignedType(columnType, "int8", "uint8")
	case strings.Contains(dbType, "uint8"):
		return "uint8"
	case strings.Contains(dbType, "smallint"):
		return unsignedType(columnType, "int16", "uint16")
	case strings.Contains(dbType, "uint16"):
		return "uint16"
	case strings.Contains(dbType, "mediumint"):
		return unsignedType(columnType, "int32", "uint32")
	case strings.Contains(dbType, "uint32"):
		return "uint32"
	case strings.Contains(dbType, "bigint"), strings.Contains(dbType, "bigserial"):
		return unsignedType(columnType, "int64", "uint64")
	case strings.Contains(dbType, "uint64"):
		return "uint64"
	case strings.Contains(dbType, "smallserial"):
		return "int16"
	case strings.Contains(dbType, "serial"):
		return "int"
	case strings.Contains(dbType, "int"):
		return unsignedType(columnType, "int", "uint")
	case strings.Contains(dbType, "number"), strings.Contains(dbType, "numeric"), strings.Contains(dbType, "decimal"):
		return "float64"
	case strings.Contains(dbType, "bool"):
		return "bool"
	case strings.Contains(dbType, "float"):
		return "float32"
	case strings.Contains(dbType, "double"), strings.Contains(dbType, "real"):
		return "float64"
	case strings.Contains(dbType, "money"):
		return "string"
	case strings.Contains(dbType, "date"), strings.Contains(dbType, "time"):
		return "time.Time"
	case strings.Contains(dbType, "blob"), strings.Contains(dbType, "binary"), strings.Contains(dbType, "bytea"):
		return "[]byte"
	case strings.Contains(dbType, "char"), strings.Contains(dbType, "string"), strings.Contains(dbType, "text"), strings.Contains(dbType, "clob"), strings.Contains(dbType, "xml"), strings.Contains(dbType, "varchar2"), strings.Contains(dbType, "nvarchar2"):
		return "string"
	case isNetworkDBType(dbType), isGeometryDBType(dbType), strings.Contains(dbType, "interval"):
		return "string"
	default:
		return "string"
	}
}

// canUsePointer 判断可空列是否适合用指针表达；切片类型本身已经可为 nil。
func canUsePointer(goType string) bool {
	return !strings.HasPrefix(goType, "[]")
}

// importForType 根据字段类型返回生成文件需要补充的 import 路径。
func importForType(cfg Config, goType string) string {
	if mapping, ok := lookupCustomTypeMappingByGoType(cfg, goType); ok {
		return mapping.ImportPath
	}

	switch strings.TrimPrefix(goType, "*") {
	case "time.Time":
		return "time"
	case "datatypes.JSON":
		return "gorm.io/datatypes"
	default:
		return ""
	}
}

// mergeDataTypeMap 将内置类型映射和用户自定义映射合并，自定义映射优先。
func mergeDataTypeMap(base map[string]func(gorm.ColumnType) string, custom map[string]TypeMapping) map[string]func(gorm.ColumnType) string {
	if len(custom) == 0 {
		return base
	}

	merged := make(map[string]func(gorm.ColumnType) string, len(base)+len(custom))
	for key, value := range base {
		merged[key] = value
	}
	for key, mapping := range custom {
		goType := mapping.GoType
		merged[normalizeDBType(key)] = func(gorm.ColumnType) string {
			return goType
		}
	}
	return merged
}

// lookupCustomTypeMapping 查找列对应的用户自定义类型映射。
func lookupCustomTypeMapping(cfg Config, column gorm.ColumnType) (TypeMapping, bool) {
	if len(cfg.TypeMappings) == 0 {
		return TypeMapping{}, false
	}

	dbType := normalizeDBType(column.DatabaseTypeName())
	if mapping, ok := cfg.TypeMappings[dbType]; ok {
		return mapping, true
	}

	columnType, _ := column.ColumnType()
	for _, key := range typeLookupKeys(columnType) {
		if mapping, ok := cfg.TypeMappings[key]; ok {
			return mapping, true
		}
	}
	return TypeMapping{}, false
}

// lookupCustomTypeMappingByGoType 用于为自定义 Go 类型补充 import。
func lookupCustomTypeMappingByGoType(cfg Config, goType string) (TypeMapping, bool) {
	if len(cfg.TypeMappings) == 0 {
		return TypeMapping{}, false
	}

	goType = strings.TrimPrefix(goType, "*")
	for _, mapping := range cfg.TypeMappings {
		if strings.TrimPrefix(mapping.GoType, "*") == goType {
			return mapping, mapping.ImportPath != ""
		}
	}
	return TypeMapping{}, false
}

// typeValidateRules 将数据库类型映射为 gookit/validate 的基础类型规则。
func typeValidateRules(column gorm.ColumnType) []string {
	dbType := normalizeDBType(column.DatabaseTypeName())
	columnType, _ := column.ColumnType()
	columnType = strings.ToLower(columnType)

	switch {
	case strings.Contains(dbType, "json"):
		return nil
	case isArrayDBType(dbType, columnType):
		return []string{"array"}
	case strings.HasPrefix(columnType, "enum("):
		if values := enumToIn(columnType); values != "" {
			return []string{"in:" + values, "string"}
		}
		return []string{"string"}
	case strings.Contains(dbType, "bool"), strings.Contains(columnType, "tinyint(1)"):
		return []string{"bool"}
	case strings.Contains(columnType, "unsigned"):
		return []string{"uint"}
	case strings.Contains(dbType, "int"), strings.Contains(dbType, "serial"):
		return []string{"int"}
	case strings.Contains(dbType, "float"), strings.Contains(dbType, "double"), strings.Contains(dbType, "decimal"), strings.Contains(dbType, "numeric"), strings.Contains(dbType, "number"), strings.Contains(dbType, "real"):
		return []string{"float"}
	case strings.Contains(dbType, "char"), strings.Contains(dbType, "string"), strings.Contains(dbType, "text"), strings.Contains(dbType, "clob"), strings.Contains(dbType, "uuid"), strings.Contains(dbType, "xml"), strings.Contains(dbType, "money"), strings.Contains(dbType, "interval"):
		return []string{"string"}
	case isNetworkDBType(dbType), isGeometryDBType(dbType):
		return []string{"string"}
	default:
		return nil
	}
}

// normalizeDBType 统一处理不同数据库返回的大小写、Nullable/LowCardinality 等类型包装。
func normalizeDBType(raw string) string {
	value := strings.ToLower(strings.TrimSpace(raw))
	for {
		open := strings.IndexByte(value, '(')
		if open <= 0 || !strings.HasSuffix(value, ")") {
			return value
		}
		wrapper := strings.TrimSpace(value[:open])
		if wrapper != "nullable" && wrapper != "lowcardinality" {
			return value
		}
		value = strings.TrimSpace(value[open+1 : len(value)-1])
	}
}

// typeLookupKeys 为 ColumnType 生成宽松匹配 key，兼容 varchar(32)、Nullable(Int64) 等写法。
func typeLookupKeys(raw string) []string {
	raw = strings.ToLower(strings.TrimSpace(raw))
	if raw == "" {
		return nil
	}

	keys := []string{normalizeDBType(raw)}
	if open := strings.IndexByte(raw, '('); open > 0 {
		keys = append(keys, normalizeDBType(raw[:open]))
	}
	if strings.Contains(raw, " unsigned") {
		keys = append(keys, strings.TrimSpace(strings.ReplaceAll(raw, " unsigned", "")))
	}
	return dedupeStrings(keys)
}

// isArrayDBType 判断 PostgreSQL 数组或其他数据库数组列类型。
func isArrayDBType(dbType, columnType string) bool {
	return strings.Contains(dbType, "array") ||
		strings.HasPrefix(dbType, "_") ||
		strings.HasSuffix(dbType, "[]") ||
		strings.Contains(columnType, "[]")
}

// isNetworkDBType 判断 PostgreSQL 网络地址类型。
func isNetworkDBType(dbType string) bool {
	switch dbType {
	case "inet", "cidr", "macaddr", "macaddr8":
		return true
	default:
		return false
	}
}

// isGeometryDBType 判断常见空间/几何类型。
func isGeometryDBType(dbType string) bool {
	switch dbType {
	case "geometry", "geography", "point", "line", "lseg", "box", "path", "polygon", "circle":
		return true
	default:
		return false
	}
}

// unsignedType 根据列定义中的 unsigned 标记选择有符号或无符号 Go 类型。
func unsignedType(columnType, signed, unsigned string) string {
	if strings.Contains(columnType, "unsigned") {
		return unsigned
	}
	return signed
}

// isSizedString 判断列是否为可从长度信息生成 maxLen 规则的字符类型。
func isSizedString(column gorm.ColumnType) bool {
	dbType := normalizeDBType(column.DatabaseTypeName())
	return (strings.Contains(dbType, "char") || strings.Contains(dbType, "string")) && !strings.Contains(dbType, "text")
}

// datetimeType 根据列可空性为时间列选择 time.Time 或 *time.Time。
func datetimeType(column gorm.ColumnType) string {
	nullable, _ := column.Nullable()
	if column.Name() == "deleted_at" || nullable {
		return "*time.Time"
	}
	return "time.Time"
}

// mysqlDataTypeMap 覆盖 gorm.io/gen 对 MySQL 类型到 Go 类型的映射。
var mysqlDataTypeMap = map[string]func(gorm.ColumnType) string{
	"tinyint":            func(gorm.ColumnType) string { return "int8" },
	"tinyint unsigned":   func(gorm.ColumnType) string { return "uint8" },
	"bool":               func(gorm.ColumnType) string { return "bool" },
	"boolean":            func(gorm.ColumnType) string { return "bool" },
	"smallint":           func(gorm.ColumnType) string { return "int16" },
	"smallint unsigned":  func(gorm.ColumnType) string { return "uint16" },
	"mediumint":          func(gorm.ColumnType) string { return "int32" },
	"mediumint unsigned": func(gorm.ColumnType) string { return "uint32" },
	"int":                func(gorm.ColumnType) string { return "int" },
	"integer":            func(gorm.ColumnType) string { return "int" },
	"int unsigned":       func(gorm.ColumnType) string { return "uint" },
	"bigint":             func(gorm.ColumnType) string { return "int64" },
	"bigint unsigned":    func(gorm.ColumnType) string { return "uint64" },
	"serial":             func(gorm.ColumnType) string { return "uint64" },
	"year":               func(gorm.ColumnType) string { return "int16" },
	"float":              func(gorm.ColumnType) string { return "float32" },
	"double":             func(gorm.ColumnType) string { return "float64" },
	"decimal":            func(gorm.ColumnType) string { return "float64" },
	"numeric":            func(gorm.ColumnType) string { return "float64" },
	"real":               func(gorm.ColumnType) string { return "float64" },
	"char":               func(gorm.ColumnType) string { return "string" },
	"varchar":            func(gorm.ColumnType) string { return "string" },
	"nchar":              func(gorm.ColumnType) string { return "string" },
	"nvarchar":           func(gorm.ColumnType) string { return "string" },
	"tinytext":           func(gorm.ColumnType) string { return "string" },
	"text":               func(gorm.ColumnType) string { return "string" },
	"mediumtext":         func(gorm.ColumnType) string { return "string" },
	"longtext":           func(gorm.ColumnType) string { return "string" },
	"clob":               func(gorm.ColumnType) string { return "string" },
	"enum":               func(gorm.ColumnType) string { return "string" },
	"set":                func(gorm.ColumnType) string { return "string" },
	"binary":             func(gorm.ColumnType) string { return "[]byte" },
	"varbinary":          func(gorm.ColumnType) string { return "[]byte" },
	"tinyblob":           func(gorm.ColumnType) string { return "[]byte" },
	"blob":               func(gorm.ColumnType) string { return "[]byte" },
	"mediumblob":         func(gorm.ColumnType) string { return "[]byte" },
	"longblob":           func(gorm.ColumnType) string { return "[]byte" },
	"bit":                func(gorm.ColumnType) string { return "[]byte" },
	"json":               func(gorm.ColumnType) string { return "datatypes.JSON" },
	"date":               func(gorm.ColumnType) string { return "time.Time" },
	"datetime":           datetimeType,
	"timestamp":          datetimeType,
	"time":               func(gorm.ColumnType) string { return "time.Time" },
	"geometry":           func(gorm.ColumnType) string { return "[]byte" },
	"point":              func(gorm.ColumnType) string { return "[]byte" },
	"linestring":         func(gorm.ColumnType) string { return "[]byte" },
	"polygon":            func(gorm.ColumnType) string { return "[]byte" },
	"multipoint":         func(gorm.ColumnType) string { return "[]byte" },
	"multilinestring":    func(gorm.ColumnType) string { return "[]byte" },
	"multipolygon":       func(gorm.ColumnType) string { return "[]byte" },
}

// postgresDataTypeMap 覆盖 gorm.io/gen 对 PostgreSQL/GaussDB 类型到 Go 类型的映射。
var postgresDataTypeMap = map[string]func(gorm.ColumnType) string{
	"smallint":                    func(gorm.ColumnType) string { return "int16" },
	"int2":                        func(gorm.ColumnType) string { return "int16" },
	"integer":                     func(gorm.ColumnType) string { return "int" },
	"int":                         func(gorm.ColumnType) string { return "int" },
	"int4":                        func(gorm.ColumnType) string { return "int" },
	"bigint":                      func(gorm.ColumnType) string { return "int64" },
	"int8":                        func(gorm.ColumnType) string { return "int64" },
	"smallserial":                 func(gorm.ColumnType) string { return "int16" },
	"serial":                      func(gorm.ColumnType) string { return "int" },
	"bigserial":                   func(gorm.ColumnType) string { return "int64" },
	"real":                        func(gorm.ColumnType) string { return "float32" },
	"float4":                      func(gorm.ColumnType) string { return "float32" },
	"double precision":            func(gorm.ColumnType) string { return "float64" },
	"float8":                      func(gorm.ColumnType) string { return "float64" },
	"numeric":                     func(gorm.ColumnType) string { return "float64" },
	"decimal":                     func(gorm.ColumnType) string { return "float64" },
	"money":                       func(gorm.ColumnType) string { return "string" },
	"character varying":           func(gorm.ColumnType) string { return "string" },
	"character":                   func(gorm.ColumnType) string { return "string" },
	"varchar":                     func(gorm.ColumnType) string { return "string" },
	"char":                        func(gorm.ColumnType) string { return "string" },
	"text":                        func(gorm.ColumnType) string { return "string" },
	"bytea":                       func(gorm.ColumnType) string { return "[]byte" },
	"uuid":                        func(gorm.ColumnType) string { return "string" },
	"json":                        func(gorm.ColumnType) string { return "datatypes.JSON" },
	"jsonb":                       func(gorm.ColumnType) string { return "datatypes.JSON" },
	"date":                        func(gorm.ColumnType) string { return "time.Time" },
	"timestamp":                   datetimeType,
	"timestamp without time zone": datetimeType,
	"timestamp with time zone":    datetimeType,
	"time":                        func(gorm.ColumnType) string { return "time.Time" },
	"time without time zone":      func(gorm.ColumnType) string { return "time.Time" },
	"time with time zone":         func(gorm.ColumnType) string { return "time.Time" },
	"interval":                    func(gorm.ColumnType) string { return "string" },
	"boolean":                     func(gorm.ColumnType) string { return "bool" },
	"bool":                        func(gorm.ColumnType) string { return "bool" },
	"bit":                         func(gorm.ColumnType) string { return "string" },
	"bit varying":                 func(gorm.ColumnType) string { return "string" },
	"varbit":                      func(gorm.ColumnType) string { return "string" },
	"xml":                         func(gorm.ColumnType) string { return "string" },
	"inet":                        func(gorm.ColumnType) string { return "string" },
	"cidr":                        func(gorm.ColumnType) string { return "string" },
	"macaddr":                     func(gorm.ColumnType) string { return "string" },
	"macaddr8":                    func(gorm.ColumnType) string { return "string" },
	"tsvector":                    func(gorm.ColumnType) string { return "string" },
	"tsquery":                     func(gorm.ColumnType) string { return "string" },
	"point":                       func(gorm.ColumnType) string { return "string" },
	"line":                        func(gorm.ColumnType) string { return "string" },
	"lseg":                        func(gorm.ColumnType) string { return "string" },
	"box":                         func(gorm.ColumnType) string { return "string" },
	"path":                        func(gorm.ColumnType) string { return "string" },
	"polygon":                     func(gorm.ColumnType) string { return "string" },
	"circle":                      func(gorm.ColumnType) string { return "string" },
	"array":                       func(gorm.ColumnType) string { return "[]string" },
}

// sqliteDataTypeMap 覆盖 SQLite 常见类型和类型亲和性的 Go 类型映射。
var sqliteDataTypeMap = map[string]func(gorm.ColumnType) string{
	"integer":   func(gorm.ColumnType) string { return "int64" },
	"int":       func(gorm.ColumnType) string { return "int" },
	"bigint":    func(gorm.ColumnType) string { return "int64" },
	"smallint":  func(gorm.ColumnType) string { return "int16" },
	"tinyint":   func(gorm.ColumnType) string { return "int8" },
	"boolean":   func(gorm.ColumnType) string { return "bool" },
	"bool":      func(gorm.ColumnType) string { return "bool" },
	"real":      func(gorm.ColumnType) string { return "float64" },
	"double":    func(gorm.ColumnType) string { return "float64" },
	"float":     func(gorm.ColumnType) string { return "float32" },
	"numeric":   func(gorm.ColumnType) string { return "float64" },
	"decimal":   func(gorm.ColumnType) string { return "float64" },
	"text":      func(gorm.ColumnType) string { return "string" },
	"varchar":   func(gorm.ColumnType) string { return "string" },
	"char":      func(gorm.ColumnType) string { return "string" },
	"clob":      func(gorm.ColumnType) string { return "string" },
	"blob":      func(gorm.ColumnType) string { return "[]byte" },
	"json":      func(gorm.ColumnType) string { return "datatypes.JSON" },
	"date":      func(gorm.ColumnType) string { return "time.Time" },
	"datetime":  datetimeType,
	"timestamp": datetimeType,
	"time":      func(gorm.ColumnType) string { return "time.Time" },
}

// sqlserverDataTypeMap 覆盖 SQL Server 常见类型到 Go 类型的映射。
var sqlserverDataTypeMap = map[string]func(gorm.ColumnType) string{
	"bit":              func(gorm.ColumnType) string { return "bool" },
	"tinyint":          func(gorm.ColumnType) string { return "uint8" },
	"smallint":         func(gorm.ColumnType) string { return "int16" },
	"int":              func(gorm.ColumnType) string { return "int" },
	"bigint":           func(gorm.ColumnType) string { return "int64" },
	"real":             func(gorm.ColumnType) string { return "float32" },
	"float":            func(gorm.ColumnType) string { return "float64" },
	"decimal":          func(gorm.ColumnType) string { return "float64" },
	"numeric":          func(gorm.ColumnType) string { return "float64" },
	"money":            func(gorm.ColumnType) string { return "float64" },
	"smallmoney":       func(gorm.ColumnType) string { return "float64" },
	"char":             func(gorm.ColumnType) string { return "string" },
	"varchar":          func(gorm.ColumnType) string { return "string" },
	"text":             func(gorm.ColumnType) string { return "string" },
	"nchar":            func(gorm.ColumnType) string { return "string" },
	"nvarchar":         func(gorm.ColumnType) string { return "string" },
	"ntext":            func(gorm.ColumnType) string { return "string" },
	"uniqueidentifier": func(gorm.ColumnType) string { return "string" },
	"xml":              func(gorm.ColumnType) string { return "string" },
	"date":             func(gorm.ColumnType) string { return "time.Time" },
	"time":             func(gorm.ColumnType) string { return "time.Time" },
	"datetime":         datetimeType,
	"datetime2":        datetimeType,
	"smalldatetime":    datetimeType,
	"datetimeoffset":   datetimeType,
	"binary":           func(gorm.ColumnType) string { return "[]byte" },
	"varbinary":        func(gorm.ColumnType) string { return "[]byte" },
	"image":            func(gorm.ColumnType) string { return "[]byte" },
	"rowversion":       func(gorm.ColumnType) string { return "[]byte" },
	"timestamp":        func(gorm.ColumnType) string { return "[]byte" },
	"json":             func(gorm.ColumnType) string { return "datatypes.JSON" },
	"geography":        func(gorm.ColumnType) string { return "string" },
	"geometry":         func(gorm.ColumnType) string { return "string" },
	"hierarchyid":      func(gorm.ColumnType) string { return "string" },
}

// clickhouseDataTypeMap 覆盖 ClickHouse 常见类型到 Go 类型的映射。
var clickhouseDataTypeMap = map[string]func(gorm.ColumnType) string{
	"int8":        func(gorm.ColumnType) string { return "int8" },
	"int16":       func(gorm.ColumnType) string { return "int16" },
	"int32":       func(gorm.ColumnType) string { return "int32" },
	"int64":       func(gorm.ColumnType) string { return "int64" },
	"int128":      func(gorm.ColumnType) string { return "string" },
	"int256":      func(gorm.ColumnType) string { return "string" },
	"uint8":       func(gorm.ColumnType) string { return "uint8" },
	"uint16":      func(gorm.ColumnType) string { return "uint16" },
	"uint32":      func(gorm.ColumnType) string { return "uint32" },
	"uint64":      func(gorm.ColumnType) string { return "uint64" },
	"uint128":     func(gorm.ColumnType) string { return "string" },
	"uint256":     func(gorm.ColumnType) string { return "string" },
	"float32":     func(gorm.ColumnType) string { return "float32" },
	"float64":     func(gorm.ColumnType) string { return "float64" },
	"decimal":     func(gorm.ColumnType) string { return "float64" },
	"bool":        func(gorm.ColumnType) string { return "bool" },
	"boolean":     func(gorm.ColumnType) string { return "bool" },
	"string":      func(gorm.ColumnType) string { return "string" },
	"fixedstring": func(gorm.ColumnType) string { return "string" },
	"uuid":        func(gorm.ColumnType) string { return "string" },
	"ipv4":        func(gorm.ColumnType) string { return "string" },
	"ipv6":        func(gorm.ColumnType) string { return "string" },
	"date":        func(gorm.ColumnType) string { return "time.Time" },
	"date32":      func(gorm.ColumnType) string { return "time.Time" },
	"datetime":    datetimeType,
	"datetime64":  datetimeType,
	"enum8":       func(gorm.ColumnType) string { return "string" },
	"enum16":      func(gorm.ColumnType) string { return "string" },
	"array":       func(gorm.ColumnType) string { return "[]string" },
	"map":         func(gorm.ColumnType) string { return "datatypes.JSON" },
	"tuple":       func(gorm.ColumnType) string { return "datatypes.JSON" },
	"json":        func(gorm.ColumnType) string { return "datatypes.JSON" },
}

// damengDataTypeMap 覆盖达梦数据库常见类型到 Go 类型的映射。
var damengDataTypeMap = map[string]func(gorm.ColumnType) string{
	"bit":       func(gorm.ColumnType) string { return "bool" },
	"tinyint":   func(gorm.ColumnType) string { return "int8" },
	"smallint":  func(gorm.ColumnType) string { return "int16" },
	"int":       func(gorm.ColumnType) string { return "int" },
	"integer":   func(gorm.ColumnType) string { return "int" },
	"bigint":    func(gorm.ColumnType) string { return "int64" },
	"number":    func(gorm.ColumnType) string { return "float64" },
	"numeric":   func(gorm.ColumnType) string { return "float64" },
	"decimal":   func(gorm.ColumnType) string { return "float64" },
	"dec":       func(gorm.ColumnType) string { return "float64" },
	"float":     func(gorm.ColumnType) string { return "float64" },
	"double":    func(gorm.ColumnType) string { return "float64" },
	"real":      func(gorm.ColumnType) string { return "float32" },
	"char":      func(gorm.ColumnType) string { return "string" },
	"varchar":   func(gorm.ColumnType) string { return "string" },
	"varchar2":  func(gorm.ColumnType) string { return "string" },
	"nchar":     func(gorm.ColumnType) string { return "string" },
	"nvarchar":  func(gorm.ColumnType) string { return "string" },
	"nvarchar2": func(gorm.ColumnType) string { return "string" },
	"text":      func(gorm.ColumnType) string { return "string" },
	"clob":      func(gorm.ColumnType) string { return "string" },
	"blob":      func(gorm.ColumnType) string { return "[]byte" },
	"binary":    func(gorm.ColumnType) string { return "[]byte" },
	"varbinary": func(gorm.ColumnType) string { return "[]byte" },
	"date":      func(gorm.ColumnType) string { return "time.Time" },
	"time":      func(gorm.ColumnType) string { return "time.Time" },
	"datetime":  datetimeType,
	"timestamp": datetimeType,
	"json":      func(gorm.ColumnType) string { return "datatypes.JSON" },
}

// oracleDataTypeMap 覆盖 Oracle 常见类型到 Go 类型的映射。
var oracleDataTypeMap = map[string]func(gorm.ColumnType) string{
	"number":                         func(gorm.ColumnType) string { return "float64" },
	"numeric":                        func(gorm.ColumnType) string { return "float64" },
	"decimal":                        func(gorm.ColumnType) string { return "float64" },
	"dec":                            func(gorm.ColumnType) string { return "float64" },
	"binary_float":                   func(gorm.ColumnType) string { return "float32" },
	"binary_double":                  func(gorm.ColumnType) string { return "float64" },
	"float":                          func(gorm.ColumnType) string { return "float64" },
	"char":                           func(gorm.ColumnType) string { return "string" },
	"varchar":                        func(gorm.ColumnType) string { return "string" },
	"varchar2":                       func(gorm.ColumnType) string { return "string" },
	"nchar":                          func(gorm.ColumnType) string { return "string" },
	"nvarchar2":                      func(gorm.ColumnType) string { return "string" },
	"long":                           func(gorm.ColumnType) string { return "string" },
	"clob":                           func(gorm.ColumnType) string { return "string" },
	"nclob":                          func(gorm.ColumnType) string { return "string" },
	"raw":                            func(gorm.ColumnType) string { return "[]byte" },
	"long raw":                       func(gorm.ColumnType) string { return "[]byte" },
	"blob":                           func(gorm.ColumnType) string { return "[]byte" },
	"bfile":                          func(gorm.ColumnType) string { return "[]byte" },
	"date":                           func(gorm.ColumnType) string { return "time.Time" },
	"timestamp":                      datetimeType,
	"timestamp with time zone":       datetimeType,
	"timestamp with local time zone": datetimeType,
	"interval year to month":         func(gorm.ColumnType) string { return "string" },
	"interval day to second":         func(gorm.ColumnType) string { return "string" },
	"json":                           func(gorm.ColumnType) string { return "datatypes.JSON" },
	"xmltype":                        func(gorm.ColumnType) string { return "string" },
}
