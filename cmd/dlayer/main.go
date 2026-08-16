package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"dlayer/generator"
	"dlayer/web"
)

func main() {
	// Parse subcommand if provided
	if len(os.Args) > 1 && (os.Args[1] == "web" || os.Args[1] == "server") {
		runWebServer(os.Args[2:])
		return
	}

	var configPath string
	var webMode bool
	var webPort int
	var noOpen bool

	flag.StringVar(&configPath, "config", "", "config file path (.yaml, .yml, .toml, .json)")
	flag.StringVar(&configPath, "c", "", "config file path (.yaml, .yml, .toml, .json)")
	flag.BoolVar(&webMode, "web", false, "start web configuration server")
	flag.BoolVar(&webMode, "w", false, "start web configuration server")
	flag.IntVar(&webPort, "port", 8080, "web server port")
	flag.IntVar(&webPort, "p", 8080, "web server port")
	flag.BoolVar(&noOpen, "no-open", false, "do not automatically open web browser")
	flag.Parse()

	if webMode {
		server := web.NewServer(webPort, configPath)
		if err := server.Start(!noOpen); err != nil {
			fmt.Fprintf(os.Stderr, "web server error: %v\n", err)
			os.Exit(1)
		}
		return
	}

	// Standard CLI generation mode
	cfg, path, err := loadConfig(configPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "load config failed: %v\n", err)
		os.Exit(2)
	}

	if err := generator.Run(cfg); err != nil {
		fmt.Fprintf(os.Stderr, "generate failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("✓ successfully generated model/query/validator code under %s by %s\n", cfg.OutDir, path)
}

func runWebServer(args []string) {
	fs := flag.NewFlagSet("web", flag.ExitOnError)
	var configPath string
	var port int
	var noOpen bool

	fs.StringVar(&configPath, "config", "", "config file path")
	fs.StringVar(&configPath, "c", "", "config file path")
	fs.IntVar(&port, "port", 8080, "web server port")
	fs.IntVar(&port, "p", 8080, "web server port")
	fs.BoolVar(&noOpen, "no-open", false, "do not automatically open web browser")

	_ = fs.Parse(args)

	// Check if port environment variable is set
	if envPort := os.Getenv("PORT"); envPort != "" {
		if p, err := strconv.Atoi(envPort); err == nil && p > 0 {
			port = p
		}
	}

	server := web.NewServer(port, strings.TrimSpace(configPath))
	if err := server.Start(!noOpen); err != nil {
		fmt.Fprintf(os.Stderr, "web server error: %v\n", err)
		os.Exit(1)
	}
}

func loadConfig(path string) (generator.Config, string, error) {
	if path == "" {
		return generator.LoadDefaultConfig()
	}
	cfg, err := generator.LoadConfig(path)
	return cfg, path, err
}
