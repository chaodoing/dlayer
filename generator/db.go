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
