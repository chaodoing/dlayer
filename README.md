# 验证器和表模型生成器

这是一个 Go 命令行工具，用于根据数据库表结构生成三类代码：

- 基于 `gorm.io/gen` 的模型层和 query 操作层
- 参考 `pkg/generation` 风格生成请求验证结构
- 基于 `github.com/gookit/validate` 生成验证 tag、场景规则和中文错误消息

## 安装依赖

```bash
go mod tidy
```

## 快速开始

复制 YAML 或 TOML 示例配置并按实际数据库修改：

```bash
cp generated.example.yaml generated.yaml
# 或
cp generated.example.toml generated.toml
go run ./cmd/generated
```

默认会读取数据库中的全部表，并生成：

```text
generated/
  model/       # gorm.io/gen 生成的模型
  query/       # gorm.io/gen 生成的 query 操作层
  validator/   # 带 gookit/validate tag、场景和消息的请求验证结构
```

程序会按顺序读取 `generated.yaml`、`generated.yml`、`generated.toml`、`generator.yaml`、`generator.yml`、`generator.toml` 中第一个存在的配置文件。

只生成指定表：

```yaml
tables:
  - users
  - orders
```

## 配置说明

- `driver`：数据库驱动，支持 `mysql`、`tidb`、`postgres`、`gaussdb`、`sqlite`、`sqlserver`、`clickhouse`、`dm`、`oracle`
- `dsn`：数据库连接字符串
- `out_dir`：生成代码的根目录，默认 `generated`
- `model_out`：模型输出目录，默认 `$out_dir/model`
- `query_out`：query 输出目录，默认 `$out_dir/query`
- `validator_out`：验证结构输出目录，默认 `$out_dir/validator`
- `validator_package`：验证包名，默认 `validator`
- `tables`：表名列表；为空时生成全部表
- `table_prefix`：表名前缀，生成模型/验证结构名时会移除
- `ignore_fields`：验证结构中忽略的字段，默认 `created_at`、`updated_at`、`deleted_at`
- `tags`：验证结构字段 tag，默认 `json`、`form`、`xml`、`url`
- `insert_scene`、`update_scene`、`delete_scene`：验证场景名，默认 `insert`、`update`、`delete`
- `type_mappings`：自定义数据库类型映射

完整配置模板见 `generated.example.yaml` 和 `generated.example.toml`。

## 自定义类型映射

当数据库中的类型需要映射为项目自定义类型时，可以在配置文件中设置 `type_mappings`：

```yaml
type_mappings:
  uuid:
    go_type: uuid.UUID
    import_path: github.com/google/uuid
  jsonb:
    go_type: datatypes.JSON
    import_path: gorm.io/datatypes
```

TOML 写法：

```toml
[type_mappings.uuid]
go_type = "uuid.UUID"
import_path = "github.com/google/uuid"

[type_mappings.jsonb]
go_type = "datatypes.JSON"
import_path = "gorm.io/datatypes"
```

该映射会同时影响 `gorm.io/gen` 生成的模型字段类型，以及验证器结构体字段类型和 import。

## 数据库 DSN 示例

```bash
# TiDB 兼容 MySQL 协议
driver: tidb
dsn: "root:password@tcp(127.0.0.1:4000)/demo?charset=utf8mb4&parseTime=True&loc=Local"

# GaussDB 兼容 PostgreSQL 协议
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

## 验证层示例

假设表 `users` 包含 `name`、`email`、`age` 等字段，工具会生成类似结构：

```go
type UserRequest struct {
	ID    int64  `json:"id" form:"id" xml:"id" url:"id" validate:"required|int"`
	Name  string `json:"name" form:"name" xml:"name" url:"name" validate:"required|string|maxLen:64"`
	Email string `json:"email" form:"email" xml:"email" url:"email" validate:"required|string|maxLen:128|email"`
	Age   int    `json:"age" form:"age" xml:"age" url:"age" validate:"required|int"`
}
```

业务代码中可以这样使用：

```go
input := validator.UserRequest{
	Name:  "Alice",
	Email: "alice@example.com",
	Age:   18,
}

if err := validator.Validate(&input, validator.SceneInsert); err != nil {
	return err
}
```

每个请求结构都会生成 `ConfigValidation` 和 `Messages` 方法，用于配置 Insert/Update/Delete 场景和中文错误提示。

## 使用示例

仓库中的 `examples/schema.sql` 提供了一个最小表结构，`examples/usage.go` 展示了生成代码后的调用方式。
