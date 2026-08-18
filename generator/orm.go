package generator

import (
	"fmt"
	"io"
	"log"
	"path/filepath"
	"strings"

	"gorm.io/gen"
	"gorm.io/gorm"
)

// generateORM 使用 gorm.io/gen 生成模型与类型安全查询层。
func generateORM(db *gorm.DB, cfg Config, tables []string) (err error) {
	queryOut, err := filepath.Abs(cfg.QueryOut)
	if err != nil {
		return fmt.Errorf("resolve query output path: %w", err)
	}

	g := gen.NewGenerator(gen.Config{
		OutPath:           queryOut,
		ModelPkgPath:      genModelPkgPath(queryOut, cfg.ModelOut),
		Mode:              gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
		FieldCoverable:    false,
		FieldNullable:     false,
		FieldSignable:     true,
		FieldWithTypeTag:  true,
		FieldWithIndexTag: true,
		WithUnitTest:      false,
	})
	g.WithOpts(fieldMappingModelOpt(cfg), nullablePointerModelOpt())
	g.WithImportPkgPath(fieldMappingImportPaths(cfg)...)

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

			// Compute struct name stripping TablePrefix if configured
			rawName := table
			if cfg.TablePrefix != "" && strings.HasPrefix(rawName, cfg.TablePrefix) {
				rawName = strings.TrimPrefix(rawName, cfg.TablePrefix)
			}
			modelName := toExportedName(rawName)

			// Build ModelOpts including per-table field type overrides
			var opts []gen.ModelOpt
			if tCfg, ok := cfg.TableConfigs[table]; ok {
				for colName, fieldCfg := range tCfg.Fields {
					if fieldCfg.GoType != "" {
						opts = append(opts, gen.FieldType(colName, fieldCfg.GoType))
					}
				}
			}

			if model := g.GenerateModelAs(table, modelName, opts...); model != nil {
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

// genModelPkgPath 返回 gorm.io/gen 所需的 ModelPkgPath。
// gorm gen 在 ModelPkgPath 不含路径分隔符时，会把 model 写到 query 目录的同级子目录；
// 若直接传入 model_out 文件路径（如 generated/model），会被当成独立目录解析，
// 在非模块根目录运行时会触发 parse model pkg path 警告。
func genModelPkgPath(queryOut, modelOut string) string {
	queryOut = filepath.Clean(queryOut)
	modelOut = filepath.Clean(modelOut)

	expectedModelOut := filepath.Join(filepath.Dir(queryOut), filepath.Base(modelOut))
	if modelOut == expectedModelOut {
		return filepath.Base(modelOut)
	}

	absModelOut, err := filepath.Abs(modelOut)
	if err != nil {
		return filepath.Base(modelOut)
	}
	return absModelOut
}
