package generator

import (
	"fmt"
	"io"
	"log"

	"gorm.io/gen"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// generateORM 使用 gorm.io/gen 生成模型与类型安全查询层。
func generateORM(db *gorm.DB, cfg Config, tables []string) (err error) {
	db.NamingStrategy = schema.NamingStrategy{
		TablePrefix:   cfg.TablePrefix,
		SingularTable: true,
		NoLowerCase:   true,
	}

	g := gen.NewGenerator(gen.Config{
		OutPath:           cfg.QueryOut,
		ModelPkgPath:      cfg.ModelOut,
		Mode:              gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
		FieldCoverable:    true,
		FieldSignable:     true,
		FieldWithTypeTag:  true,
		FieldWithIndexTag: true,
		WithUnitTest:      false,
	})

	switch normalizeDriver(cfg.Driver) {
	case "mysql", "tidb":
		g.WithDataTypeMap(mergeDataTypeMap(mysqlDataTypeMap, cfg.TypeMappings))
	case "postgres", "gaussdb":
		g.WithDataTypeMap(mergeDataTypeMap(postgresDataTypeMap, cfg.TypeMappings))
	case "sqlserver":
		g.WithDataTypeMap(mergeDataTypeMap(sqlserverDataTypeMap, cfg.TypeMappings))
	case "clickhouse":
		g.WithDataTypeMap(mergeDataTypeMap(clickhouseDataTypeMap, cfg.TypeMappings))
	case "dm":
		g.WithDataTypeMap(mergeDataTypeMap(damengDataTypeMap, cfg.TypeMappings))
	case "oracle":
		g.WithDataTypeMap(mergeDataTypeMap(oracleDataTypeMap, cfg.TypeMappings))
	case "sqlite", "sqlite3":
		g.WithDataTypeMap(mergeDataTypeMap(sqliteDataTypeMap, cfg.TypeMappings))
	}
	g.UseDB(db)

	models := make([]interface{}, 0, len(tables))
	for _, table := range tables {
		func() {
			defer func() {
				if r := recover(); r != nil {
					fmt.Printf("gen: skip table %s: %v\n", table, r)
				}
			}()
			if model := g.GenerateModel(table); model != nil {
				models = append(models, model)
			}
		}()
	}
	if len(models) == 0 {
		return fmt.Errorf("no models generated")
	}
	g.ApplyBasic(models...)

	origOutput := log.Writer()
	log.SetOutput(io.Discard)
	defer log.SetOutput(origOutput)

	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("gen: execute failed: %v", r)
		}
	}()
	g.Execute()
	return nil
}
