package generator

import (
	"regexp"
	"sort"
	"strings"
)

var nonIdentifierChars = regexp.MustCompile(`[^a-zA-Z0-9]+`)

// toExportedName 将数据库命名转换为可导出的 Go 标识符。
func toExportedName(name string) string {
	parts := nonIdentifierChars.Split(name, -1)
	var b strings.Builder
	for _, part := range parts {
		if part == "" {
			continue
		}
		lower := strings.ToLower(part)
		b.WriteString(strings.ToUpper(lower[:1]))
		if len(lower) > 1 {
			b.WriteString(lower[1:])
		}
	}
	if b.Len() == 0 {
		return "Field"
	}
	value := b.String()
	if value[0] >= '0' && value[0] <= '9' {
		return "Field" + value
	}
	return value
}

// singularize 对常见英文复数表名做轻量单数化处理。
func singularize(table string) string {
	lower := strings.ToLower(table)
	switch {
	case strings.HasSuffix(lower, "ies") && len(table) > 3:
		return table[:len(table)-3] + "y"
	case strings.HasSuffix(lower, "ses") && len(table) > 3:
		return table[:len(table)-2]
	case strings.HasSuffix(lower, "s") && !strings.HasSuffix(lower, "ss") && !strings.HasSuffix(lower, "us") && len(table) > 3:
		return table[:len(table)-1]
	default:
		return table
	}
}

// validatorStructName 根据表名和前缀生成请求验证结构名称。
// 不做单数化，避免 sys_dictionary 和 sys_dictionaries 这类表名生成同名结构。
func validatorStructName(table, prefix string) string {
	name := table
	if prefix != "" && strings.HasPrefix(name, prefix) {
		name = strings.TrimPrefix(name, prefix)
	}
	return toExportedName(name) + "Request"
}

// toFileName 将表名转换为安全的 snake_case 文件名前缀。
func toFileName(name string) string {
	parts := nonIdentifierChars.Split(strings.ToLower(name), -1)
	filtered := parts[:0]
	for _, part := range parts {
		if part != "" {
			filtered = append(filtered, part)
		}
	}
	if len(filtered) == 0 {
		return "table"
	}
	return strings.Join(filtered, "_")
}

// sortedKeys 返回排序后的 map key，保证生成代码稳定。
func sortedKeys[V any](values map[string]V) []string {
	keys := make([]string, 0, len(values))
	for key := range values {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}

// dedupeStrings 在保留原始顺序的同时移除空值和重复值。
func dedupeStrings(values []string) []string {
	out := make([]string, 0, len(values))
	seen := make(map[string]struct{}, len(values))
	for _, value := range values {
		if strings.TrimSpace(value) == "" {
			continue
		}
		if _, ok := seen[value]; ok {
			continue
		}
		seen[value] = struct{}{}
		out = append(out, value)
	}
	return out
}
