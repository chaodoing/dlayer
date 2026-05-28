package validator

import (
	"github.com/gookit/validate"
)

// SafePlatformRequest is the request validator for the safe_platform table.
type SafePlatformRequest struct {
	Id       uint16 `json:"id" form:"id" xml:"id" url:"id" validate:"required|uint" label:"主键"`
	Title    string `json:"title" form:"title" xml:"title" url:"title" validate:"required|string|maxLen:30" label:"平台名称"`
	Disabled uint8  `json:"disabled" form:"disabled" xml:"disabled" url:"disabled" validate:"uint" label:"禁用状态"`
	Sort     uint16 `json:"sort" form:"sort" xml:"sort" url:"sort" validate:"required|uint" label:"平台排序"`
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
		"maxLen":   "{field}长度不能大于 %v",
		"required": "{field}不能为空",
		"string":   "{field}必须是字符串",
		"uint":     "{field}必须是非负整数",
	}
}
