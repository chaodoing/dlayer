package generator

import (
	"fmt"
	"sort"
	"strings"

	dameng "github.com/godoes/gorm-dameng"
	oracle "github.com/godoes/gorm-oracle"
	"gorm.io/driver/clickhouse"
	"gorm.io/driver/gaussdb"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

// openDB 根据配置中的驱动名称创建 GORM 数据库连接。
func openDB(driver, dsn string) (*gorm.DB, error) {
	switch normalizeDriver(driver) {
	case "mysql", "tidb":
		return openMySQLCompatible(dsn)
	case "postgres":
		return gorm.Open(postgres.Open(dsn), &gorm.Config{})
	case "gaussdb":
		return gorm.Open(gaussdb.Open(dsn), &gorm.Config{})
	case "sqlite", "sqlite3":
		return gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	case "sqlserver":
		return gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	case "clickhouse":
		return gorm.Open(clickhouse.Open(dsn), &gorm.Config{})
	case "dm":
		return gorm.Open(dameng.Open(dsn), &gorm.Config{})
	case "oracle":
		return gorm.Open(oracle.Open(dsn), &gorm.Config{})
	default:
		return nil, fmt.Errorf("unsupported driver %q", driver)
	}
}

func openMySQLCompatible(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil && strings.Contains(err.Error(), "missing the slash separating the database name") {
		return nil, fmt.Errorf("invalid mysql/tidb dsn %q: database name must follow a slash, for example root:password@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local: %w", dsn, err)
	}
	return db, err
}

// normalizeDriver 将命令行传入的数据库名称别名归一化。
func normalizeDriver(driver string) string {
	switch strings.ToLower(strings.TrimSpace(driver)) {
	case "postgresql", "pgsql":
		return "postgres"
	case "tidb":
		return "tidb"
	case "gauss", "opengauss", "gaussdb":
		return "gaussdb"
	case "mssql", "sql-server", "sql_server", "sqlserver":
		return "sqlserver"
	case "clickhouse", "click-house", "ch":
		return "clickhouse"
	case "dm", "dameng", "dm8":
		return "dm"
	case "oracle", "ora":
		return "oracle"
	default:
		return strings.ToLower(strings.TrimSpace(driver))
	}
}

// resolveTables 返回本次需要生成的表；未显式指定时读取当前数据库全部表。
func resolveTables(db *gorm.DB, selected []string) ([]string, error) {
	if len(selected) > 0 {
		return selected, nil
	}

	tables, err := db.Migrator().GetTables()
	if err != nil {
		return nil, fmt.Errorf("list tables: %w", err)
	}
	sort.Strings(tables)
	return tables, nil
}

// TestDBConnection 测试数据库连接，成功时返回当前数据库中的全部表名。
func TestDBConnection(driver, dsn string) ([]string, error) {
	db, err := openDB(driver, dsn)
	if err != nil {
		return nil, err
	}
	sqlDB, err := db.DB()
	if err == nil {
		defer sqlDB.Close()
	}

	tables, err := db.Migrator().GetTables()
	if err != nil {
		return nil, fmt.Errorf("connection successful, but failed to list tables: %w", err)
	}
	sort.Strings(tables)
	return tables, nil
}

// ColumnDetail 描述表单列的元数据（包括数据类型、是否可为空、默认 Go 类型等）。
type ColumnDetail struct {
	Name         string `json:"name"`
	DatabaseType string `json:"database_type"`
	GoType       string `json:"go_type"`
	Nullable     bool   `json:"nullable"`
	Primary      bool   `json:"primary"`
	Comment      string `json:"comment"`
}

// GetTableColumns 连接数据库并返回指定表的详细列元数据。
func GetTableColumns(driver, dsn, table string, cfg Config) ([]ColumnDetail, error) {
	db, err := openDB(driver, dsn)
	if err != nil {
		return nil, err
	}
	sqlDB, err := db.DB()
	if err == nil {
		defer sqlDB.Close()
	}

	columnTypes, err := db.Migrator().ColumnTypes(table)
	if err != nil {
		return nil, fmt.Errorf("read column types for table %s: %w", table, err)
	}

	details := make([]ColumnDetail, 0, len(columnTypes))
	for _, col := range columnTypes {
		nullable, _ := col.Nullable()
		primary, _ := col.PrimaryKey()
		comment, _ := col.Comment()
		inferredType := goTypeForColumn(cfg, col)

		details = append(details, ColumnDetail{
			Name:         col.Name(),
			DatabaseType: col.DatabaseTypeName(),
			GoType:       inferredType,
			Nullable:     nullable,
			Primary:      primary,
			Comment:      comment,
		})
	}
	return details, nil
}


