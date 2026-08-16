package web

import (
	"embed"
	"encoding/json"
	"fmt"
	"io/fs"
	"net"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"dlayer/generator"
)

//go:embed static/*
var staticFS embed.FS

// Server 封装 Web 配置界面 HTTP 服务器。
type Server struct {
	port       int
	configPath string
}

// NewServer 创建 Web 服务器实例。
func NewServer(port int, configPath string) *Server {
	if port <= 0 {
		port = 8080
	}
	return &Server{
		port:       port,
		configPath: configPath,
	}
}

// Start 启动 HTTP 服务器并在控制台输出访问地址。
func (s *Server) Start(openBrowser bool) error {
	subFS, err := fs.Sub(staticFS, "static")
	if err != nil {
		return fmt.Errorf("embedded static assets error: %w", err)
	}

	mux := http.NewServeMux()
	fileServer := http.FileServer(http.FS(subFS))
	mux.Handle("/", fileServer)

	mux.HandleFunc("/api/drivers", s.handleDrivers)
	mux.HandleFunc("/api/config", s.handleConfig)
	mux.HandleFunc("/api/config/save", s.handleSaveConfig)
	mux.HandleFunc("/api/db/test", s.handleTestDB)
	mux.HandleFunc("/api/db/columns", s.handleColumns)
	mux.HandleFunc("/api/generate", s.handleGenerate)

	addr := fmt.Sprintf(":%d", s.port)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return fmt.Errorf("failed to listen on %s: %w", addr, err)
	}

	actualPort := listener.Addr().(*net.TCPAddr).Port
	url := fmt.Sprintf("http://localhost:%d", actualPort)

	fmt.Println("==================================================")
	fmt.Println("  🚀 dlayer Web 配置界面已启动！")
	fmt.Printf("  👉 访问地址: %s\n", url)
	fmt.Println("==================================================")

	if openBrowser {
		go openURL(url)
	}

	server := &http.Server{
		Handler:      mux,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 60 * time.Second,
	}
	return server.Serve(listener)
}

func (s *Server) handleDrivers(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	drivers := generator.SupportedDrivers()
	writeJSON(w, http.StatusOK, map[string]interface{}{
		"drivers": drivers,
	})
}

func (s *Server) handleConfig(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var cfg generator.Config
	var loadedPath string
	var err error

	if s.configPath != "" {
		cfg, err = generator.LoadConfig(s.configPath)
		loadedPath = s.configPath
	} else {
		cfg, loadedPath, err = generator.LoadDefaultConfig()
	}

	if err != nil {
		cfg = generator.DefaultConfig()
		cfg.ApplyDefaults()
		if loadedPath == "" {
			loadedPath = "generated.yaml"
		}
	}

	writeJSON(w, http.StatusOK, map[string]interface{}{
		"config": cfg,
		"path":   loadedPath,
	})
}

type saveConfigRequest struct {
	Path   string           `json:"path"`
	Config generator.Config `json:"config"`
}

func (s *Server) handleSaveConfig(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req saveConfigRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   fmt.Sprintf("Invalid JSON request: %v", err),
		})
		return
	}

	savePath := req.Path
	if strings.TrimSpace(savePath) == "" {
		savePath = "generated.yaml"
	}

	req.Config.ApplyDefaults()
	if err := req.Config.Validate(); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   fmt.Sprintf("Config validation failed: %v", err),
		})
		return
	}

	if err := generator.SaveConfig(savePath, req.Config); err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]interface{}{
			"success": false,
			"error":   fmt.Sprintf("Save config failed: %v", err),
		})
		return
	}

	writeJSON(w, http.StatusOK, map[string]interface{}{
		"success": true,
		"message": fmt.Sprintf("Successfully saved config to %s", savePath),
		"path":    savePath,
	})
}

type testDBRequest struct {
	Driver string `json:"driver"`
	DSN    string `json:"dsn"`
}

func (s *Server) handleTestDB(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req testDBRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   fmt.Sprintf("Invalid JSON request: %v", err),
		})
		return
	}

	if strings.TrimSpace(req.Driver) == "" {
		writeJSON(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   "Database driver is required",
		})
		return
	}
	if strings.TrimSpace(req.DSN) == "" {
		writeJSON(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   "Database DSN is required",
		})
		return
	}

	tables, err := generator.TestDBConnection(req.Driver, req.DSN)
	if err != nil {
		writeJSON(w, http.StatusOK, map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	writeJSON(w, http.StatusOK, map[string]interface{}{
		"success":     true,
		"tables":      tables,
		"table_count": len(tables),
	})
}

type tableColumnsRequest struct {
	Driver string           `json:"driver"`
	DSN    string           `json:"dsn"`
	Table  string           `json:"table"`
	Config generator.Config `json:"config"`
}

func (s *Server) handleColumns(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req tableColumnsRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   fmt.Sprintf("Invalid JSON request: %v", err),
		})
		return
	}

	if strings.TrimSpace(req.Driver) == "" || strings.TrimSpace(req.DSN) == "" || strings.TrimSpace(req.Table) == "" {
		writeJSON(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   "Driver, DSN, and Table are required",
		})
		return
	}

	columns, err := generator.GetTableColumns(req.Driver, req.DSN, req.Table, req.Config)
	if err != nil {
		writeJSON(w, http.StatusOK, map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	writeJSON(w, http.StatusOK, map[string]interface{}{
		"success": true,
		"table":   req.Table,
		"columns": columns,
	})
}

type generateRequest struct {
	Config   generator.Config `json:"config"`
	SavePath string           `json:"save_path"`
	AutoSave bool             `json:"auto_save"`
}

func (s *Server) handleGenerate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req generateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   fmt.Sprintf("Invalid JSON request: %v", err),
		})
		return
	}

	cfg := req.Config
	cfg.ApplyDefaults()

	if err := cfg.Validate(); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   fmt.Sprintf("Invalid config: %v", err),
		})
		return
	}

	if req.AutoSave && strings.TrimSpace(req.SavePath) != "" {
		_ = generator.SaveConfig(req.SavePath, cfg)
	}

	err := generator.Run(cfg)
	if err != nil {
		writeJSON(w, http.StatusOK, map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	// List generated files in OutDir
	var generatedFiles []string
	if cfg.OutDir != "" {
		_ = filepath.Walk(cfg.OutDir, func(path string, info os.FileInfo, err error) error {
			if err == nil && !info.IsDir() {
				generatedFiles = append(generatedFiles, path)
			}
			return nil
		})
	}

	writeJSON(w, http.StatusOK, map[string]interface{}{
		"success": true,
		"message": "Code generation completed successfully!",
		"files":   generatedFiles,
		"config":  cfg,
	})
}

func writeJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(data)
}

func openURL(url string) {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "darwin":
		cmd = "open"
		args = []string{url}
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start", url}
	default:
		cmd = "xdg-open"
		args = []string{url}
	}
	_ = exec.Command(cmd, args...).Start()
}
