package main

import (
	"fmt"
	"os"

	"generated/generator"
)

func main() {
	cfg, path, err := generator.LoadDefaultConfig()
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
