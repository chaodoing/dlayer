package generator

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"gorm.io/gorm"
)

// generateValidators 为每张表生成请求验证结构，并写入公共 Validate 辅助文件。
func generateValidators(db *gorm.DB, cfg Config, tables []string) error {
	if err := os.MkdirAll(cfg.ValidatorOut, 0o755); err != nil {
		return fmt.Errorf("create validator output directory: %w", err)
	}

	if err := writeValidateHelper(cfg); err != nil {
		return err
	}

	for _, table := range tables {
		columns, err := db.Migrator().ColumnTypes(table)
		if err != nil {
			return fmt.Errorf("read columns for table %s: %w", table, err)
		}
		if err := writeValidatorFile(cfg, table, columns); err != nil {
			return err
		}
	}
	return nil
}

// writeValidateHelper 生成验证包的公共场景常量和 Validate 包装函数。
func writeValidateHelper(cfg Config) error {
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "package %s\n\n", cfg.ValidatorPackage)
	buf.WriteString("import \"github.com/gookit/validate\"\n\n")
	fmt.Fprintf(&buf, "const (\n\tSceneInsert = %q\n\tSceneUpdate = %q\n\tSceneDelete = %q\n)\n\n", sceneName(cfg.InsertScene, "insert"), sceneName(cfg.UpdateScene, "update"), sceneName(cfg.DeleteScene, "delete"))
	buf.WriteString("// Validate checks a generated request struct with gookit/validate.\n")
	buf.WriteString("func Validate(value any, scenes ...string) error {\n")
	buf.WriteString("\tv := validate.Struct(value)\n")
	buf.WriteString("\tif v.Validate(scenes...) {\n")
	buf.WriteString("\t\treturn nil\n")
	buf.WriteString("\t}\n")
	buf.WriteString("\treturn v.Errors.ErrOrNil()\n")
	buf.WriteString("}\n")

	path := filepath.Join(cfg.ValidatorOut, "validate.go")
	return writeGoFile(path, buf.Bytes())
}

// writeValidatorFile 将单张表的列元数据转换为一个请求验证结构及框架中间件文件。
func writeValidatorFile(cfg Config, table string, columns []gorm.ColumnType) error {
	structName := validatorStructName(table, cfg.TablePrefix)
	middlewareName := validatorMiddlewareName(table, cfg.TablePrefix)
	infos := make([]fieldInfo, 0, len(columns))

	isIris := strings.EqualFold(cfg.Framework, "iris")
	imports := map[string]struct{}{
		"net/http":                   {},
		"strings":                    {},
		"github.com/gookit/validate": {},
	}
	if isIris {
		imports["github.com/kataras/iris/v12"] = struct{}{}
		imports["pkg/iris/boot"] = struct{}{}
		imports["pkg/iris/o"] = struct{}{}
	} else {
		imports["github.com/gin-gonic/gin"] = struct{}{}
		imports["pkg/gin/boot"] = struct{}{}
		imports["pkg/gin/o"] = struct{}{}
	}

	ignore := ignoreSet(cfg.IgnoreFields)
	tags := cfg.Tags
	if len(tags) == 0 {
		tags = []string{"json", "form", "xml", "url"}
	}

	tableConfig, hasTableConfig := cfg.TableConfigs[table]
	for _, column := range columns {
		if ignore[column.Name()] {
			continue
		}
		info := buildFieldInfo(cfg, column)
		if hasTableConfig {
			if fieldOverride, ok := tableConfig.Fields[column.Name()]; ok && strings.TrimSpace(fieldOverride.GoType) != "" {
				info.Type = strings.TrimSpace(fieldOverride.GoType)
				if strings.TrimSpace(fieldOverride.ImportPath) != "" {
					imports[strings.TrimSpace(fieldOverride.ImportPath)] = struct{}{}
				}
			}
		}
		if importPath := importForType(cfg, info.Type); importPath != "" {
			imports[importPath] = struct{}{}
		}
		infos = append(infos, info)
	}

	var buf bytes.Buffer
	fmt.Fprintf(&buf, "package %s\n\n", cfg.ValidatorPackage)
	if len(imports) > 0 {
		buf.WriteString("import (\n")
		for _, importPath := range sortedKeys(imports) {
			fmt.Fprintf(&buf, "\t%q\n", importPath)
		}
		buf.WriteString(")\n\n")
	}
	fmt.Fprintf(&buf, "// %s %s\n", structName, table)
	fmt.Fprintf(&buf, "type %s struct {\n", structName)
	for _, info := range infos {
		fieldTags := make([]string, 0, len(tags)+2)
		for _, tag := range tags {
			fieldTags = append(fieldTags, fmt.Sprintf("%s:%q", tag, info.ColumnName))
		}
		if info.ValidateTag != "" {
			fieldTags = append(fieldTags, fmt.Sprintf("validate:%q", info.ValidateTag))
		}
		if info.Comment != "" {
			fieldTags = append(fieldTags, fmt.Sprintf("label:%q", info.Comment))
		}
		fmt.Fprintf(&buf, "\t%s %s `%s`\n", info.Name, info.Type, strings.Join(fieldTags, " "))
	}
	buf.WriteString("}\n\n")

	writeMiddleware(&buf, cfg, table, structName, middlewareName)
	writeScenes(&buf, cfg, structName, infos)
	writeMessages(&buf, structName, infos)

	path := filepath.Join(cfg.ValidatorOut, toFileName(table)+"_validator.go")
	return writeGoFile(path, buf.Bytes())
}

// writeMiddleware 为请求校验结构生成 Iris 或 Gin 对应框架的请求中间件。
func writeMiddleware(buf *bytes.Buffer, cfg Config, table, structName, middlewareName string) {
	isIris := strings.EqualFold(cfg.Framework, "iris")
	if isIris {
		fmt.Fprintf(buf, "// %s %s请求中间件（Iris）\n", middlewareName, table)
		fmt.Fprintf(buf, "func %s(ctx iris.Context, container *boot.Container) {\n", middlewareName)
		fmt.Fprintf(buf, "\tvar data %s\n", structName)
		buf.WriteString("\tif err := ctx.ReadJSON(&data); err != nil {\n")
		buf.WriteString("\t\to.O(ctx, 204, err.Error())\n")
		buf.WriteString("\t\treturn\n")
		buf.WriteString("\t}\n")
		buf.WriteString("\tif strings.EqualFold(ctx.Method(), http.MethodPost) {\n")
		buf.WriteString("\t\tif err := Validate(&data, SceneInsert); err != nil {\n")
		buf.WriteString("\t\t\to.O(ctx, 206, err.Error())\n")
		buf.WriteString("\t\t\treturn\n")
		buf.WriteString("\t\t}\n")
		buf.WriteString("\t\tctx.Next()\n")
		buf.WriteString("\t} else if strings.EqualFold(ctx.Method(), http.MethodPut) {\n")
		buf.WriteString("\t\tif err := Validate(&data, SceneUpdate); err != nil {\n")
		buf.WriteString("\t\t\to.O(ctx, 206, err.Error())\n")
		buf.WriteString("\t\t\treturn\n")
		buf.WriteString("\t\t}\n")
		buf.WriteString("\t\tctx.Next()\n")
		buf.WriteString("\t}\n")
		buf.WriteString("}\n\n")
	} else {
		fmt.Fprintf(buf, "// %s %s请求中间件（Gin）\n", middlewareName, table)
		fmt.Fprintf(buf, "func %s(ctx *gin.Context, container *boot.Containers) {\n", middlewareName)
		fmt.Fprintf(buf, "\tvar data %s\n", structName)
		buf.WriteString("\tif err := ctx.ShouldBindJSON(&data); err != nil {\n")
		buf.WriteString("\t\to.O(ctx, 204, err.Error())\n")
		buf.WriteString("\t\treturn\n")
		buf.WriteString("\t}\n")
		buf.WriteString("\tif strings.EqualFold(ctx.Request.Method, http.MethodPost) {\n")
		buf.WriteString("\t\tif err := Validate(&data, SceneInsert); err != nil {\n")
		buf.WriteString("\t\t\to.O(ctx, 206, err.Error())\n")
		buf.WriteString("\t\t\treturn\n")
		buf.WriteString("\t\t}\n")
		buf.WriteString("\t\tctx.Next()\n")
		buf.WriteString("\t} else if strings.EqualFold(ctx.Request.Method, http.MethodPut) {\n")
		buf.WriteString("\t\tif err := Validate(&data, SceneUpdate); err != nil {\n")
		buf.WriteString("\t\t\to.O(ctx, 206, err.Error())\n")
		buf.WriteString("\t\t\treturn\n")
		buf.WriteString("\t\t}\n")
		buf.WriteString("\t\tctx.Next()\n")
		buf.WriteString("\t}\n")
		buf.WriteString("}\n\n")
	}
}

type fieldInfo struct {
	// Name 是生成到 Go 结构体中的字段名。
	Name string
	// ColumnName 是数据库原始列名，用于生成 json/form/xml/url tag。
	ColumnName string
	// Type 是字段的 Go 类型字符串。
	Type string
	// ValidateTag 是拼接后的 gookit/validate 规则。
	ValidateTag string
	// Comment 是数据库列注释，会写入 label tag。
	Comment string
	// Primary 标记该字段是否为主键，用于场景生成。
	Primary bool
}

// buildFieldInfo 汇总列名、Go 类型、验证规则和注释等模板所需字段信息。
func buildFieldInfo(cfg Config, column gorm.ColumnType) fieldInfo {
	nullable, nullableOK := column.Nullable()
	primary, _ := column.PrimaryKey()
	autoIncrement, _ := column.AutoIncrement()
	length, hasLength := column.Length()
	comment, _ := column.Comment()
	defaultValue, hasDefault := column.DefaultValue()
	hasDefault = hasDefault && strings.TrimSpace(defaultValue) != ""

	goType := goTypeForColumn(cfg, column)
	if nullableOK && nullable && canUsePointer(goType) {
		goType = "*" + goType
	}

	rules := make([]string, 0, 3)
	if (!nullableOK || !nullable) && !hasDefault && (!autoIncrement || isIDColumn(column.Name(), primary)) {
		rules = append(rules, "required")
	}
	rules = append(rules, typeValidateRules(column)...)
	if strings.TrimPrefix(goType, "*") == "string" && hasLength && length > 0 && length < 65535 && isSizedString(column) {
		rules = append(rules, "maxLen:"+strconv.FormatInt(length, 10))
	}
	if isEmailColumn(column.Name()) {
		rules = append(rules, "email")
	}

	return fieldInfo{
		Name:        toExportedName(column.Name()),
		ColumnName:  column.Name(),
		Type:        goType,
		ValidateTag: strings.Join(dedupeStrings(rules), "|"),
		Comment:     comment,
		Primary:     primary,
	}
}

// writeScenes 生成 gookit/validate 的 Insert、Update、Delete 场景配置。
func writeScenes(buf *bytes.Buffer, cfg Config, structName string, fields []fieldInfo) {
	fmt.Fprintf(buf, "// ConfigValidation 配置验证规则场景\n")
	fmt.Fprintf(buf, "func (%s) ConfigValidation(v *validate.Validation) {\n", structName)
	buf.WriteString("\tv.WithScenes(validate.SValues{\n")
	writeScene(buf, sceneVar(cfg.InsertScene, "insert", "SceneInsert"), insertFields(fields))
	writeScene(buf, sceneVar(cfg.UpdateScene, "update", "SceneUpdate"), fieldNames(fields))
	writeScene(buf, sceneVar(cfg.DeleteScene, "delete", "SceneDelete"), deleteFields(fields))
	buf.WriteString("\t})\n")
	buf.WriteString("}\n\n")
}

func sceneVar(value, fallback, constVar string) string {
	val := sceneName(value, fallback)
	if val == fallback {
		return constVar
	}
	return fmt.Sprintf("%q", val)
}

// writeScene 写入单个场景到 validate.SValues。
func writeScene(buf *bytes.Buffer, sceneExpr string, fields []string) {
	fmt.Fprintf(buf, "\t%s: []string{", sceneExpr)
	for i, field := range fields {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(buf, "%q", field)
	}
	buf.WriteString("},\n")
}

// writeMessages 根据字段使用到的规则生成中文错误提示。
func writeMessages(buf *bytes.Buffer, structName string, fields []fieldInfo) {
	messages := collectMessages(fields)
	if len(messages) == 0 {
		return
	}
	fmt.Fprintf(buf, "// Messages defines Chinese validation messages.\n")
	fmt.Fprintf(buf, "func (%s) Messages() map[string]string {\n", structName)
	buf.WriteString("\treturn validate.MS{\n")
	for _, key := range sortedKeys(messages) {
		fmt.Fprintf(buf, "\t\t%q: %q,\n", key, messages[key])
	}
	buf.WriteString("\t}\n")
	buf.WriteString("}\n")
}

// insertFields 返回新增场景需要校验的字段，默认排除 id 主键。
func insertFields(fields []fieldInfo) []string {
	names := make([]string, 0, len(fields))
	for _, field := range fields {
		if field.Primary && strings.EqualFold(field.ColumnName, "id") {
			continue
		}
		names = append(names, field.Name)
	}
	return names
}

// fieldNames 返回所有生成到请求结构中的字段名。
func fieldNames(fields []fieldInfo) []string {
	names := make([]string, 0, len(fields))
	for _, field := range fields {
		names = append(names, field.Name)
	}
	return names
}

// deleteFields 返回删除场景需要校验的字段，优先使用 id 主键。
func deleteFields(fields []fieldInfo) []string {
	for _, field := range fields {
		if field.Primary && strings.EqualFold(field.ColumnName, "id") {
			return []string{field.Name}
		}
	}
	for _, field := range fields {
		if strings.EqualFold(field.ColumnName, "id") {
			return []string{field.Name}
		}
	}
	return nil
}

// collectMessages 从字段验证规则中收集需要输出的错误消息。
func collectMessages(fields []fieldInfo) map[string]string {
	messages := map[string]string{}
	for _, field := range fields {
		for _, rule := range strings.Split(field.ValidateTag, "|") {
			name := strings.TrimSpace(strings.SplitN(rule, ":", 2)[0])
			if name == "" {
				continue
			}
			if msg, ok := validatorMessages[name]; ok {
				messages[msg.Type] = msg.Message
			}
		}
	}
	return messages
}

// sceneName 返回用户配置的场景名，空值时使用默认场景名。
func sceneName(value, fallback string) string {
	if strings.TrimSpace(value) == "" {
		return fallback
	}
	return strings.TrimSpace(value)
}

// ignoreSet 将忽略字段列表转换为便于查找的集合。
func ignoreSet(fields []string) map[string]bool {
	ignore := map[string]bool{}
	for _, field := range fields {
		if strings.TrimSpace(field) != "" {
			ignore[strings.TrimSpace(field)] = true
		}
	}
	return ignore
}

// isIDColumn 判断列是否为约定的 id 主键。
func isIDColumn(name string, primary bool) bool {
	return primary && strings.EqualFold(name, "id")
}

// isEmailColumn 通过列名约定为 email 字段追加 email 格式校验。
func isEmailColumn(name string) bool {
	lower := strings.ToLower(name)
	return lower == "email" || strings.HasSuffix(lower, "_email")
}

// enumToIn 将 MySQL enum('a','b') 列定义转换为 gookit/validate 的 in:a,b 参数。
func enumToIn(colType string) string {
	colType = strings.TrimSpace(colType)
	if !strings.HasPrefix(strings.ToLower(colType), "enum(") {
		return ""
	}
	body := strings.TrimPrefix(colType, "enum(")
	body = strings.TrimSuffix(body, ")")
	raw := strings.Split(body, ",")
	values := make([]string, 0, len(raw))
	for _, value := range raw {
		value = strings.Trim(value, `' "`)
		if value != "" {
			values = append(values, value)
		}
	}
	return strings.Join(values, ",")
}
