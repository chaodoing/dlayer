package generator

type validatorMessage struct {
	// Type 是 validate.MS 中使用的规则 key。
	Type string
	// Message 是对应的中文错误提示模板。
	Message string
}

// validatorMessages 是当前生成器内置的中文验证错误消息集合。
var validatorMessages = map[string]validatorMessage{
	"required": {"required", "{field}不能为空"},
	"int":      {"int", "{field}必须是整数"},
	"uint":     {"uint", "{field}必须是非负整数"},
	"bool":     {"bool", "{field}必须是布尔值"},
	"string":   {"string", "{field}必须是字符串"},
	"float":    {"float", "{field}必须是浮点数"},
	"array":    {"array", "{field}必须是数组或切片"},
	"in":       {"in", "{field}不在允许的范围内"},
	"maxLen":   {"maxLen", "{field}长度不能大于 %v"},
	"email":    {"email", "{field}格式不正确"},
}
