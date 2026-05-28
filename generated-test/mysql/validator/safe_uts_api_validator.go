package validator

import (
	"github.com/gookit/validate"
)

// SafeUtsApiRequest is the request validator for the safe_uts_api table.
type SafeUtsApiRequest struct {
	UtsId    uint   `json:"uts_id" form:"uts_id" xml:"uts_id" url:"uts_id" validate:"required|uint" label:"探针设备ID"`
	Username string `json:"username" form:"username" xml:"username" url:"username" validate:"required|string|maxLen:50" label:"用户名称"`
	Password string `json:"password" form:"password" xml:"password" url:"password" validate:"required|string|maxLen:50" label:"登录密码"`
	Secret   string `json:"secret" form:"secret" xml:"secret" url:"secret" validate:"required|string|maxLen:50" label:"认证密钥"`
}

// ConfigValidation configures gookit/validate scenes.
func (SafeUtsApiRequest) ConfigValidation(v *validate.Validation) {
	v.WithScenes(validate.SValues{
		"insert": []string{"UtsId", "Username", "Password", "Secret"},
		"update": []string{"UtsId", "Username", "Password", "Secret"},
		"delete": []string{},
	})
}

// Messages defines Chinese validation messages.
func (SafeUtsApiRequest) Messages() map[string]string {
	return validate.MS{
		"maxLen":   "{field}长度不能大于 %v",
		"required": "{field}不能为空",
		"string":   "{field}必须是字符串",
		"uint":     "{field}必须是非负整数",
	}
}
