package generator_test

import (
	"path/filepath"
	"testing"

	"dlayer/generator"
)

func TestDefaultConfig(t *testing.T) {
	cfg := generator.DefaultConfig()
	if cfg.Driver != "mysql" {
		t.Errorf("expected default driver mysql, got %s", cfg.Driver)
	}
	if cfg.OutDir != "generated" {
		t.Errorf("expected default out_dir generated, got %s", cfg.OutDir)
	}
}

func TestSaveAndLoadConfigYAML(t *testing.T) {
	tempDir := t.TempDir()
	configPath := filepath.Join(tempDir, "test_config.yaml")

	originalCfg := generator.DefaultConfig()
	originalCfg.Driver = "sqlite"
	originalCfg.DSN = "./test.db"
	originalCfg.Tables = []string{"users", "posts"}

	err := generator.SaveConfig(configPath, originalCfg)
	if err != nil {
		t.Fatalf("SaveConfig failed: %v", err)
	}

	loadedCfg, err := generator.LoadConfig(configPath)
	if err != nil {
		t.Fatalf("LoadConfig failed: %v", err)
	}

	if loadedCfg.Driver != "sqlite" {
		t.Errorf("expected driver sqlite, got %s", loadedCfg.Driver)
	}
	if loadedCfg.DSN != "./test.db" {
		t.Errorf("expected DSN ./test.db, got %s", loadedCfg.DSN)
	}
	if len(loadedCfg.Tables) != 2 || loadedCfg.Tables[0] != "users" {
		t.Errorf("expected tables [users, posts], got %v", loadedCfg.Tables)
	}
}

func TestSupportedDrivers(t *testing.T) {
	drivers := generator.SupportedDrivers()
	if len(drivers) == 0 {
		t.Fatal("expected supported drivers list not to be empty")
	}
	foundMySQL := false
	for _, d := range drivers {
		if d.Name == "mysql" {
			foundMySQL = true
			break
		}
	}
	if !foundMySQL {
		t.Error("expected mysql driver in supported drivers list")
	}
}

func TestTableConfigsSaveAndLoad(t *testing.T) {
	tempDir := t.TempDir()
	configPath := filepath.Join(tempDir, "table_config.yaml")

	cfg := generator.DefaultConfig()
	cfg.Driver = "sqlite"
	cfg.DSN = "./demo.db"
	cfg.TableConfigs = map[string]generator.TableConfig{
		"sys_admin": {
			Fields: map[string]generator.FieldTypeConfig{
				"avatar": {
					GoType:     "custom.AvatarUrl",
					ImportPath: "myproject/pkg/custom",
				},
			},
		},
	}

	if err := generator.SaveConfig(configPath, cfg); err != nil {
		t.Fatalf("SaveConfig failed: %v", err)
	}

	loaded, err := generator.LoadConfig(configPath)
	if err != nil {
		t.Fatalf("LoadConfig failed: %v", err)
	}

	tCfg, ok := loaded.TableConfigs["sys_admin"]
	if !ok {
		t.Fatalf("expected sys_admin in TableConfigs")
	}
	fCfg, ok := tCfg.Fields["avatar"]
	if !ok {
		t.Fatalf("expected avatar in sys_admin fields")
	}
	if fCfg.GoType != "custom.AvatarUrl" {
		t.Errorf("expected custom.AvatarUrl, got %s", fCfg.GoType)
	}
	if fCfg.ImportPath != "myproject/pkg/custom" {
		t.Errorf("expected myproject/pkg/custom, got %s", fCfg.ImportPath)
	}
}

func TestCommentedTOMLSaveAndLoad(t *testing.T) {
	tempDir := t.TempDir()
	configPath := filepath.Join(tempDir, "table_config.toml")

	cfg := generator.DefaultConfig()
	cfg.Driver = "sqlite"
	cfg.DSN = "./demo.db"
	cfg.TableConfigs = map[string]generator.TableConfig{
		"sys_admin": {
			Fields: map[string]generator.FieldTypeConfig{
				"avatar": {
					GoType:     "custom.AvatarUrl",
					ImportPath: "myproject/pkg/custom",
				},
			},
		},
	}

	if err := generator.SaveConfig(configPath, cfg); err != nil {
		t.Fatalf("SaveConfig failed: %v", err)
	}

	loaded, err := generator.LoadConfig(configPath)
	if err != nil {
		t.Fatalf("LoadConfig failed: %v", err)
	}

	if loaded.Driver != "sqlite" {
		t.Errorf("expected driver sqlite, got %s", loaded.Driver)
	}
	tCfg, ok := loaded.TableConfigs["sys_admin"]
	if !ok {
		t.Fatalf("expected sys_admin in TableConfigs")
	}
	fCfg, ok := tCfg.Fields["avatar"]
	if !ok {
		t.Fatalf("expected avatar in sys_admin fields")
	}
	if fCfg.GoType != "custom.AvatarUrl" {
		t.Errorf("expected custom.AvatarUrl, got %s", fCfg.GoType)
	}
}

func TestOutDirResolution(t *testing.T) {
	cfg := generator.Config{
		OutDir:       "data-layer",
		ModelOut:     "models",
		QueryOut:     "query",
		ValidatorOut: "request",
	}
	cfg.ApplyDefaults()

	expectedModel := filepath.Join("data-layer", "models")
	expectedQuery := filepath.Join("data-layer", "query")
	expectedValidator := filepath.Join("data-layer", "request")

	if cfg.ModelOut != expectedModel {
		t.Errorf("expected ModelOut %s, got %s", expectedModel, cfg.ModelOut)
	}
	if cfg.QueryOut != expectedQuery {
		t.Errorf("expected QueryOut %s, got %s", expectedQuery, cfg.QueryOut)
	}
	if cfg.ValidatorOut != expectedValidator {
		t.Errorf("expected ValidatorOut %s, got %s", expectedValidator, cfg.ValidatorOut)
	}
}



