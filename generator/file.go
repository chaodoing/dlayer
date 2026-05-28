package generator

import (
	"fmt"
	"go/format"
	"os"
	"path/filepath"
)

// writeFile 创建父目录并写入文件内容。
func writeFile(path string, content []byte) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	if err := os.WriteFile(path, content, 0o644); err != nil {
		return fmt.Errorf("write %s: %w", path, err)
	}
	return nil
}

// writeGoFile 在写入前先用 go/format 格式化生成的 Go 源码。
func writeGoFile(path string, content []byte) error {
	formatted, err := format.Source(content)
	if err != nil {
		return fmt.Errorf("format %s: %w\n%s", path, err, string(content))
	}
	return writeFile(path, formatted)
}
