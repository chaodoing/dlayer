package validator

import (
	"github.com/gookit/validate"
)

// SettingsRequest is the request validator for the settings table.
type SettingsRequest struct {
	Id       int    `json:"id" form:"id" xml:"id" url:"id" validate:"required|int" label:"主键"`
	Code     string `json:"code" form:"code" xml:"code" url:"code" validate:"required|string|maxLen:60" label:"调用编码"`
	Title    string `json:"title" form:"title" xml:"title" url:"title" validate:"required|string|maxLen:30" label:"配置名称"`
	Sort     int    `json:"sort" form:"sort" xml:"sort" url:"sort" validate:"required|int" label:"显示顺序"`
	Disabled int    `json:"disabled" form:"disabled" xml:"disabled" url:"disabled" validate:"int" label:"禁用状态"`
}

// ConfigValidation configures gookit/validate scenes.
func (SettingsRequest) ConfigValidation(v *validate.Validation) {
	v.WithScenes(validate.SValues{
		"insert": []string{"Code", "Title", "Sort", "Disabled"},
		"update": []string{"Id", "Code", "Title", "Sort", "Disabled"},
		"delete": []string{"Id"},
	})
}

// Messages defines Chinese validation messages.
func (SettingsRequest) Messages() map[string]string {
	return validate.MS{
		"int":      "{field}必须是整数",
		"maxLen":   "{field}长度不能大于 %v",
		"required": "{field}不能为空",
		"string":   "{field}必须是字符串",
	}
}
