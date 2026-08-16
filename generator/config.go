// Package generator 定义代码生成命令的运行配置和参数解析辅助函数。
package generator

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/pelletier/go-toml/v2"
	"gopkg.in/yaml.v3"
)

// Config 描述一次代码生成任务所需的全部输入。
type Config struct {
	// Driver 是数据库驱动名称，支持 mysql、tidb、postgres、gaussdb、sqlite、sqlserver、clickhouse、dm 和 oracle。
	Driver string `json:"driver" yaml:"driver" toml:"driver"`
	// DSN 是数据库连接字符串，不同驱动使用各自的连接格式。
	DSN string `json:"dsn" yaml:"dsn" toml:"dsn"`
	// OutDir 是生成代码的根目录；未显式指定子目录时会作为默认前缀。
	OutDir string `json:"out_dir" yaml:"out_dir" toml:"out_dir"`
	// ModelOut 是 gorm.io/gen 模型文件输出目录。
	ModelOut string `json:"model_out" yaml:"model_out" toml:"model_out"`
	// QueryOut 是 gorm.io/gen 查询层代码输出目录。
	QueryOut string `json:"query_out" yaml:"query_out" toml:"query_out"`
	// ValidatorOut 是请求验证结构代码输出目录。
	ValidatorOut string `json:"validator_out" yaml:"validator_out" toml:"validator_out"`
	// ValidatorPackage 是生成验证器文件使用的 Go package 名称。
	ValidatorPackage string `json:"validator_package" yaml:"validator_package" toml:"validator_package"`
	// Tables 为空时生成当前数据库中的全部表。
	Tables []string `json:"tables" yaml:"tables" toml:"tables"`
	// TablePrefix 用于从表名中剥离公共前缀，影响模型和请求结构体命名。
	TablePrefix string `json:"table_prefix" yaml:"table_prefix" toml:"table_prefix"`
	// IgnoreFields 是生成验证结构时跳过的数据库列名。
	IgnoreFields []string `json:"ignore_fields" yaml:"ignore_fields" toml:"ignore_fields"`
	// Tags 是验证结构字段上生成的结构体 tag 名称，例如 json、form、xml。
	Tags []string `json:"tags" yaml:"tags" toml:"tags"`
	// InsertScene、UpdateScene、DeleteScene 是 gookit/validate 的场景名。
	InsertScene string `json:"insert_scene" yaml:"insert_scene" toml:"insert_scene"`
	UpdateScene string `json:"update_scene" yaml:"update_scene" toml:"update_scene"`
	DeleteScene string `json:"delete_scene" yaml:"delete_scene" toml:"delete_scene"`
	// TypeMappings 是用户自定义的全局数据库类型到 Go 类型映射。
	TypeMappings map[string]TypeMapping `json:"type_mappings" yaml:"type_mappings" toml:"type_mappings"`
	// TableConfigs 是按表名独立配置的字段类型映射。
	TableConfigs map[string]TableConfig `json:"table_configs" yaml:"table_configs" toml:"table_configs"`
}

// TableConfig 描述单张表级别的独立字段配置。
type TableConfig struct {
	Fields map[string]FieldTypeConfig `json:"fields" yaml:"fields" toml:"fields"`
}

// FieldTypeConfig 描述单张表中某个列字段自定义的 Go 类型及导入路径。
type FieldTypeConfig struct {
	GoType     string `json:"go_type" yaml:"go_type" toml:"go_type"`
	ImportPath string `json:"import_path" yaml:"import_path" toml:"import_path"`
}

// TypeMapping 描述单个自定义类型映射。
type TypeMapping struct {
	// DBType 是数据库返回的类型名称，例如 uuid、jsonb、geometry。
	DBType string `json:"db_type" yaml:"db_type" toml:"db_type"`
	// GoType 是生成到模型和验证结构中的 Go 类型，例如 uuid.UUID。
	GoType string `json:"go_type" yaml:"go_type" toml:"go_type"`
	// ImportPath 是 GoType 需要的导入路径；内置类型可为空。
	ImportPath string `json:"import_path" yaml:"import_path" toml:"import_path"`
}

// DefaultConfig 返回配置文件未提供字段时使用的默认值。
func DefaultConfig() Config {
	return Config{
		Driver:           "mysql",
		OutDir:           "generated",
		ValidatorPackage: "validator",
		IgnoreFields:     []string{"created_at", "updated_at", "deleted_at"},
		Tags:             []string{"json", "form", "xml", "url"},
		InsertScene:      "insert",
		UpdateScene:      "update",
		DeleteScene:      "delete",
	}
}

// LoadDefaultConfig 按约定文件名加载配置。
func LoadDefaultConfig() (Config, string, error) {
	for _, path := range []string{"generated.yaml", "generated.yml", "generated.toml", "generator.yaml", "generator.yml", "generator.toml"} {
		if _, err := os.Stat(path); err == nil {
			cfg, err := LoadConfig(path)
			return cfg, path, err
		}
	}
	return Config{}, "", errors.New("config file not found: create generated.yaml, generated.toml, generator.yaml, or generator.toml")
}

// LoadConfig 从 YAML 或 TOML 配置文件加载生成配置。
func LoadConfig(path string) (Config, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return Config{}, fmt.Errorf("read config %s: %w", path, err)
	}

	cfg := DefaultConfig()
	switch strings.ToLower(filepath.Ext(path)) {
	case ".yaml", ".yml":
		if err := yaml.Unmarshal(content, &cfg); err != nil {
			return Config{}, fmt.Errorf("parse yaml config %s: %w", path, err)
		}
	case ".toml":
		if err := toml.Unmarshal(content, &cfg); err != nil {
			return Config{}, fmt.Errorf("parse toml config %s: %w", path, err)
		}
	default:
		return Config{}, fmt.Errorf("unsupported config file extension %q", filepath.Ext(path))
	}
	cfg.ApplyDefaults()
	cfg.resolveOutputPaths(filepath.Dir(path))
	normalizeTypeMappings(cfg.TypeMappings)
	if err := cfg.Validate(); err != nil {
		return Config{}, fmt.Errorf("validate config %s: %w", path, err)
	}
	return cfg, nil
}

// ApplyDefaults 补齐依赖其他配置推导出的默认值。
func (c *Config) ApplyDefaults() {
	defaults := DefaultConfig()
	if c.Driver == "" {
		c.Driver = defaults.Driver
	}
	if c.OutDir == "" {
		c.OutDir = defaults.OutDir
	}
	if c.ModelOut == "" {
		c.ModelOut = c.OutDir + "/model"
	}
	if c.QueryOut == "" {
		c.QueryOut = c.OutDir + "/query"
	}
	if c.ValidatorOut == "" {
		c.ValidatorOut = c.OutDir + "/validator"
	}
	if c.ValidatorPackage == "" {
		c.ValidatorPackage = defaults.ValidatorPackage
	}
	if len(c.IgnoreFields) == 0 {
		c.IgnoreFields = defaults.IgnoreFields
	}
	if len(c.Tags) == 0 {
		c.Tags = defaults.Tags
	}
	if c.InsertScene == "" {
		c.InsertScene = defaults.InsertScene
	}
	if c.UpdateScene == "" {
		c.UpdateScene = defaults.UpdateScene
	}
	if c.DeleteScene == "" {
		c.DeleteScene = defaults.DeleteScene
	}
}

// resolveOutputPaths 将相对输出目录解析为基于配置文件所在目录的绝对路径。
func (c *Config) resolveOutputPaths(baseDir string) {
	baseDir = filepath.Clean(baseDir)
	c.OutDir = resolvePath(baseDir, c.OutDir)
	c.ModelOut = resolvePath(baseDir, c.ModelOut)
	c.QueryOut = resolvePath(baseDir, c.QueryOut)
	c.ValidatorOut = resolvePath(baseDir, c.ValidatorOut)
}

func resolvePath(baseDir, path string) string {
	path = strings.TrimSpace(path)
	if path == "" || filepath.IsAbs(path) {
		return path
	}
	return filepath.Join(baseDir, path)
}

// Validate 检查必需配置和数据库驱动是否有效。
func (c Config) Validate() error {
	if c.Driver == "" {
		return errors.New("driver is required")
	}
	if c.DSN == "" {
		return errors.New("dsn is required")
	}
	if c.ModelOut == "" {
		return errors.New("model output path is required")
	}
	if c.QueryOut == "" {
		return errors.New("query output path is required")
	}
	if c.ValidatorOut == "" {
		return errors.New("validator output path is required")
	}
	if c.ValidatorPackage == "" {
		return errors.New("validator package is required")
	}

	switch strings.ToLower(c.Driver) {
	case "mysql", "tidb", "postgres", "postgresql", "pgsql", "gauss", "opengauss", "gaussdb", "sqlite", "sqlite3", "sqlserver", "sql-server", "sql_server", "mssql", "clickhouse", "click-house", "ch", "dm", "dameng", "dm8", "oracle", "ora":
		return nil
	default:
		return fmt.Errorf("unsupported driver %q", c.Driver)
	}
}

// SplitTables 将命令行中的逗号分隔表名解析为去重后的列表。
func SplitTables(raw string) []string {
	return splitCSV(raw)
}

// SplitList 将通用逗号分隔参数解析为去重后的列表。
func SplitList(raw string) []string {
	return splitCSV(raw)
}

// ParseTypeMappings 解析自定义类型映射。
// 格式：db_type=GoType 或 db_type=GoType@import/path，多个映射使用逗号分隔。
func ParseTypeMappings(raw string) (map[string]TypeMapping, error) {
	raw = strings.TrimSpace(raw)
	if raw == "" {
		return nil, nil
	}

	result := map[string]TypeMapping{}
	for _, item := range strings.Split(raw, ",") {
		item = strings.TrimSpace(item)
		if item == "" {
			continue
		}

		key, value, ok := strings.Cut(item, "=")
		if !ok {
			return nil, fmt.Errorf("invalid type mapping %q: expected db_type=GoType or db_type=GoType@import/path", item)
		}
		key = normalizeTypeKey(key)
		value = strings.TrimSpace(value)
		if key == "" || value == "" {
			return nil, fmt.Errorf("invalid type mapping %q: db_type and GoType are required", item)
		}

		goType, importPath, _ := strings.Cut(value, "@")
		goType = strings.TrimSpace(goType)
		importPath = strings.TrimSpace(importPath)
		if goType == "" {
			return nil, fmt.Errorf("invalid type mapping %q: GoType is required", item)
		}

		result[key] = TypeMapping{
			DBType:     key,
			GoType:     goType,
			ImportPath: importPath,
		}
	}
	return result, nil
}

func normalizeTypeMappings(mappings map[string]TypeMapping) {
	for key, mapping := range mappings {
		normalized := normalizeTypeKey(key)
		if mapping.DBType == "" {
			mapping.DBType = normalized
		} else {
			mapping.DBType = normalizeTypeKey(mapping.DBType)
		}
		if normalized != key {
			delete(mappings, key)
		}
		mappings[normalized] = mapping
	}
}

// splitCSV 负责统一处理空值、空白字符和重复值。
func splitCSV(raw string) []string {
	if strings.TrimSpace(raw) == "" {
		return nil
	}

	parts := strings.Split(raw, ",")
	tables := make([]string, 0, len(parts))
	seen := make(map[string]struct{}, len(parts))
	for _, part := range parts {
		name := strings.TrimSpace(part)
		if name == "" {
			continue
		}
		if _, ok := seen[name]; ok {
			continue
		}
		seen[name] = struct{}{}
		tables = append(tables, name)
	}
	return tables
}

func normalizeTypeKey(value string) string {
	value = strings.ToLower(strings.TrimSpace(value))
	if open := strings.IndexByte(value, '('); open > 0 {
		value = strings.TrimSpace(value[:open])
	}
	return value
}

// SaveConfig 将配置保存至指定路径 (支持 .yaml, .yml, .toml, .json)。
func SaveConfig(path string, cfg Config) error {
	ext := strings.ToLower(filepath.Ext(path))
	var data []byte
	var err error

	switch ext {
	case ".yaml", ".yml":
		data, err = yaml.Marshal(cfg)
	case ".toml":
		data, err = toml.Marshal(cfg)
	case ".json":
		data, err = json.MarshalIndent(cfg, "", "  ")
	default:
		return fmt.Errorf("unsupported config extension %q (supported: .yaml, .yml, .toml, .json)", ext)
	}
	if err != nil {
		return fmt.Errorf("marshal config: %w", err)
	}

	if dir := filepath.Dir(path); dir != "" && dir != "." {
		if err := os.MkdirAll(dir, 0o755); err != nil {
			return fmt.Errorf("create config dir %s: %w", dir, err)
		}
	}

	if err := os.WriteFile(path, data, 0o644); err != nil {
		return fmt.Errorf("write config %s: %w", path, err)
	}
	return nil
}

// DriverMeta 描述受支持数据库驱动的信息和 DSN 示例。
type DriverMeta struct {
	Name        string `json:"name"`
	Label       string `json:"label"`
	ExampleDSN  string `json:"example_dsn"`
	Description string `json:"description"`
}

// SupportedDrivers 返回所有支持的数据库驱动及其 DSN 示例说明。
func SupportedDrivers() []DriverMeta {
	return []DriverMeta{
		{
			Name:        "mysql",
			Label:       "MySQL / TiDB",
			ExampleDSN:  "root:password@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local",
			Description: "MySQL 和兼容 MySQL 协议的数据库（如 TiDB）",
		},
		{
			Name:        "postgres",
			Label:       "PostgreSQL",
			ExampleDSN:  "host=127.0.0.1 port=5432 user=postgres password=secret dbname=demo sslmode=disable",
			Description: "PostgreSQL 关系型数据库",
		},
		{
			Name:        "sqlite",
			Label:       "SQLite",
			ExampleDSN:  "./demo.db",
			Description: "SQLite 本地嵌入式单文件数据库",
		},
		{
			Name:        "gaussdb",
			Label:       "GaussDB / openGauss",
			ExampleDSN:  "host=127.0.0.1 port=8000 user=gauss password=secret dbname=demo sslmode=disable",
			Description: "华为 GaussDB / openGauss 数据库",
		},
		{
			Name:        "sqlserver",
			Label:       "SQL Server",
			ExampleDSN:  "sqlserver://sa:password@127.0.0.1:1433?database=demo",
			Description: "Microsoft SQL Server 数据库",
		},
		{
			Name:        "clickhouse",
			Label:       "ClickHouse",
			ExampleDSN:  "clickhouse://default:password@127.0.0.1:9000/default?dial_timeout=10s",
			Description: "ClickHouse 列式分析型数据库",
		},
		{
			Name:        "dm",
			Label:       "达梦 (DM8)",
			ExampleDSN:  "dm://SYSDBA:SYSDBA@127.0.0.1:5236?schema=SYSDBA",
			Description: "国产达梦数据库 DM8",
		},
		{
			Name:        "oracle",
			Label:       "Oracle",
			ExampleDSN:  "oracle://user:password@127.0.0.1:1521/service",
			Description: "Oracle 关系型数据库",
		},
	}
}

