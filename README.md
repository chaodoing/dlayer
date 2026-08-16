# dlayer - 验证器与表模型生成器

`dlayer` 是一个强力的 Go 语言代码生成工具与可视化配置系统。用于根据数据库表结构自动生成三类高质、类型安全的 Go 代码：

- 基于 `gorm.io/gen` 的模型层（Model）与类型安全 Query 操作层
- 基于 `github.com/gookit/validate` 的请求验证结构体（Validator）
- 包含验证 Tag、场景规则（Insert/Update/Delete）以及中文错误消息映射

---

## 🌟 核心特性

- 🖥️ **内置 Web 可视化配置面板**：无需手动手写配置文件，提供极佳的图形化界面，支持在线测试数据库连接、搜表、编辑字段类型与一键生成代码。
- 🧩 **单表独立字段类型配置**：支持对特定表中的具体列独立配置 Go 类型（如 `custom.AvatarUrl`）和 Import 路径，模型层与验证层同步生效。
- 🔌 **全主流数据库驱动支持**：支持 MySQL, TiDB, PostgreSQL, GaussDB/openGauss, SQLite, SQL Server, ClickHouse, 达梦 (DM8), Oracle。
- ⚙️ **全局数据类型映射**：支持将数据库全局类型（如 `uuid`, `jsonb`, `geometry`）自动映射为自定义 Go 类型与 Package 导入路径。
- ✂️ **表名前缀裁剪**：支持设置 `table_prefix`（如 `sys_`），自动剥离表名前缀并生成干净的 Go 结构体与文件命名。
- 🏷️ **自定义 Tag & 验证场景**：自由配置 `json`, `form`, `xml`, `url` 等 Tag 名称，支持场景化验证规则及中文错误提示。
- 💾 **多格式配置读写**：支持 YAML (`.yaml`/`.yml`)、TOML (`.toml`) 和 JSON (`.json`) 配置的持久化保存与加载。

---

## 🚀 快速开始

### 方式 1：使用 Web 可视化配置界面（推荐）

在项目根目录下运行 `web` 子命令：

```bash
# 启动 Web 配置界面（默认端口 8080，自动打开浏览器）
go run ./cmd/dlayer web

# 自定义启动端口
go run ./cmd/dlayer web -port 8999
# 或命令行 Flag 形式
go run ./cmd/dlayer -web -port 8999
```

在浏览器中打开 `http://localhost:8080`，即可：
1. **测试连接**：输入数据库 DSN 并测试连通性，自动列出全部数据表。
2. **字段类型配置**：在“单表字段类型配置”中选取表，实时查看列结构并单独指定自定义 Go 类型和 Import 路径。
3. **一键生成**：点击“立即生成代码”，在控制台中查看实时日志与输出文件。

---

### 方式 2：使用 CLI 命令行工具

1. 复制配置文件模板并修改数据库连接参数：

```bash
cp generated.example.yaml generated.yaml
# 或
cp generated.example.toml generated.toml
```

2. 执行生成命令：

```bash
go run ./cmd/dlayer
# 或指定配置文件路径
go run ./cmd/dlayer -config ./configs/dev.yaml
```

默认会在输出根目录下生成：

```text
generated/
  model/       # gorm.io/gen 生成的模型结构
  query/       # gorm.io/gen 生成的类型安全 Query 操作层
  validator/   # 带 gookit/validate tag、场景和中文消息的请求验证结构
```

---

## 📖 完整配置说明

示例配置文件（`generated.yaml` 或 `generated.toml`）：

```yaml
# 数据库驱动: mysql, tidb, postgres, gaussdb, sqlite, sqlserver, clickhouse, dm, oracle
driver: mysql
dsn: "root:password@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"

# 输出目录配置
out_dir: generated
model_out: generated/model
query_out: generated/query
validator_out: generated/validator
validator_package: validator

# 表与前缀设置
table_prefix: "sys_"
tables: # 留空时默认生成全部表
  - sys_admin
  - sys_role

# 忽略校验字段
ignore_fields:
  - created_at
  - updated_at
  - deleted_at

# 验证结构体 Tag
tags:
  - json
  - form
  - xml
  - url

# 场景名称配置
insert_scene: insert
update_scene: update
delete_scene: delete

# 全局数据库类型映射
type_mappings:
  uuid:
    db_type: uuid
    go_type: uuid.UUID
    import_path: github.com/google/uuid
  jsonb:
    db_type: jsonb
    go_type: datatypes.JSON
    import_path: gorm.io/datatypes

# 单表独立字段类型配置
table_configs:
  sys_admin:
    fields:
      avatar:
        go_type: custom.AvatarUrl
        import_path: myproject/pkg/custom
```

---

## 🗄️ 数据库 DSN 示例参考

```bash
# MySQL / TiDB
driver: mysql
dsn: "root:password@tcp(127.0.0.1:3306)/demo?charset=utf8mb4&parseTime=True&loc=Local"

# PostgreSQL
driver: postgres
dsn: "host=127.0.0.1 port=5432 user=postgres password=secret dbname=demo sslmode=disable"

# SQLite
driver: sqlite
dsn: "./demo.db"

# 华为 GaussDB / openGauss
driver: gaussdb
dsn: "host=127.0.0.1 port=8000 user=gauss password=secret dbname=demo sslmode=disable"

# SQL Server
driver: sqlserver
dsn: "sqlserver://sa:password@127.0.0.1:1433?database=demo"

# ClickHouse
driver: clickhouse
dsn: "clickhouse://default:password@127.0.0.1:9000/default?dial_timeout=10s"

# 达梦 DM8
driver: dm
dsn: "dm://SYSDBA:SYSDBA@127.0.0.1:5236?schema=SYSDBA"

# Oracle
driver: oracle
dsn: "oracle://user:password@127.0.0.1:1521/service"
```

---

## 🛠️ 生成代码使用示例

假设表 `sys_admin` 包含 `id`, `name`, `email`, `avatar` 等字段，配置了 `table_prefix: "sys_"` 以及 `avatar` 字段类型覆盖：

### 生成的模型层 (`generated/model/sys_admin.gen.go`)

```go
package model

import (
	"myproject/pkg/custom"
)

type Admin struct {
	ID       int64            `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Name     string           `gorm:"column:name;type:varchar(64);not null" json:"name"`
	Email    string           `gorm:"column:email;type:varchar(128);not null" json:"email"`
	Avatar   custom.AvatarUrl `gorm:"column:avatar;type:varchar(255)" json:"avatar"`
}
```

### 生成的验证器层 (`generated/validator/sys_admin_validator.go`)

```go
package validator

import (
	"github.com/gookit/validate"
	"myproject/pkg/custom"
)

type AdminRequest struct {
	ID     int64            `json:"id" form:"id" xml:"id" url:"id" validate:"int"`
	Name   string           `json:"name" form:"name" xml:"name" url:"name" validate:"required|string|maxLen:64"`
	Email  string           `json:"email" form:"email" xml:"email" url:"email" validate:"required|string|maxLen:128|email"`
	Avatar custom.AvatarUrl `json:"avatar" form:"avatar" xml:"avatar" url:"avatar" validate:"string"`
}

func (AdminRequest) ConfigValidation(v *validate.Validation) {
	v.WithScenes(validate.SValues{
		"insert": []string{"Name", "Email", "Avatar"},
		"update": []string{"ID", "Name", "Email", "Avatar"},
		"delete": []string{"ID"},
	})
}
```

### 业务代码中调用验证

```go
package main

import (
	"fmt"
	"yourproject/generated/validator"
)

func handleCreateAdmin() error {
	input := validator.AdminRequest{
		Name:  "Alice",
		Email: "alice@example.com",
	}

	// 校验新增场景
	if err := validator.Validate(&input, validator.SceneInsert); err != nil {
		return fmt.Errorf("参数验证错误: %w", err)
	}
	return nil
}
```

---

## 📂 项目结构

```text
.
├── cmd/
│   └── dlayer/         # CLI & Web Server 命令行入口
├── generator/          # 代码生成器核心引擎（ORM、Validator、Config、DB 元数据读取）
├── web/                # Web 可视化配置界面（HTTP 路由与嵌入式 SPA 静态文件）
├── examples/           # SQL 示例与使用示范
├── generated.example.yaml # YAML 配置模板示例
└── generated.example.toml # TOML 配置模板示例
```
