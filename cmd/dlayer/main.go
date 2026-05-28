package main

import (
	"flag"
	"fmt"
	"os"

	"generated/generator"
)

func main() {
	var configPath string
	flag.StringVar(&configPath, "config", "", "config file path (.yaml, .yml, .toml)")
	flag.StringVar(&configPath, "c", "", "config file path (.yaml, .yml, .toml)")
	flag.Parse()

	cfg, path, err := loadConfig(configPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "load config failed: %v\n", err)
		os.Exit(2)
	}

	if err := generator.Run(cfg); err != nil {
		fmt.Fprintf(os.Stderr, "generate failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("generated model/query/validator code under %s by %s\n", cfg.OutDir, path)
}

func loadConfig(path string) (generator.Config, string, error) {
	if path == "" {
		return generator.LoadDefaultConfig()
	}
	cfg, err := generator.LoadConfig(path)
	return cfg, path, err
}
