// Package generator 负责连接数据库，并根据表结构生成 GORM 模型、查询层和请求验证器代码。
package generator

import "fmt"

// Run 执行一次完整生成流程：打开数据库、解析表列表、生成 ORM 层和验证层。
func Run(cfg Config) error {
	db, err := openDB(cfg.Driver, cfg.DSN)
	if err != nil {
		return err
	}

	tables, err := resolveTables(db, cfg.Tables)
	if err != nil {
		return err
	}
	if len(tables) == 0 {
		return fmt.Errorf("no database tables found")
	}

	if err := generateORM(db, cfg, tables); err != nil {
		return err
	}
	if err := generateValidators(db, cfg, tables); err != nil {
		return err
	}
	return nil
}
