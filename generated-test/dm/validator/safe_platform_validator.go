package validator

import (
	"github.com/gookit/validate"
)

// SafePlatformRequest is the request validator for the safe_platform table.
type SafePlatformRequest struct {
	Id       int    `json:"id" form:"id" xml:"id" url:"id" validate:"required|int" label:"主键"`
	Title    string `json:"title" form:"title" xml:"title" url:"title" validate:"required|string|maxLen:120" label:"平台名称"`
	Disabled int8   `json:"disabled" form:"disabled" xml:"disabled" url:"disabled" validate:"int" label:"禁用状态"`
	Sort     int16  `json:"sort" form:"sort" xml:"sort" url:"sort" validate:"required|int" label:"平台排序"`
}

// ConfigValidation configures gookit/validate scenes.
func (SafePlatformRequest) ConfigValidation(v *validate.Validation) {
	v.WithScenes(validate.SValues{
		"insert": []string{"Title", "Disabled", "Sort"},
		"update": []string{"Id", "Title", "Disabled", "Sort"},
		"delete": []string{"Id"},
	})
}

// Messages defines Chinese validation messages.
func (SafePlatformRequest) Messages() map[string]string {
	return validate.MS{
		"int":      "{field}必须是整数",
		"maxLen":   "{field}长度不能大于 %v",
		"required": "{field}不能为空",
		"string":   "{field}必须是字符串",
	}
}
