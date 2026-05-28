package validator

import (
	"github.com/gookit/validate"
)

// SysConfigValueRequest is the request validator for the sys_config_value table.
type SysConfigValueRequest struct {
	Id       int64   `json:"id" form:"id" xml:"id" url:"id" validate:"required|int" label:"主键"`
	ConfigId int     `json:"config_id" form:"config_id" xml:"config_id" url:"config_id" validate:"required|int" label:"配置ID"`
	Title    string  `json:"title" form:"title" xml:"title" url:"title" validate:"required|string|maxLen:200" label:"配置标题"`
	Code     *string `json:"code" form:"code" xml:"code" url:"code" validate:"string|maxLen:240" label:"配置编码"`
	Value    string  `json:"value" form:"value" xml:"value" url:"value" validate:"required|string" label:"配置值"`
	Options  string  `json:"options" form:"options" xml:"options" url:"options" validate:"required|string|maxLen:32767" label:"配置项"`
	Type     string  `json:"type" form:"type" xml:"type" url:"type" validate:"required|string|maxLen:128" label:"配置类型"`
	Sort     int64   `json:"sort" form:"sort" xml:"sort" url:"sort" validate:"required|int" label:"排序"`
	Note     string  `json:"note" form:"note" xml:"note" url:"note" validate:"required|string|maxLen:480" label:"配置说明"`
	Disabled int     `json:"disabled" form:"disabled" xml:"disabled" url:"disabled" validate:"int" label:"锁定状态"`
}

// ConfigValidation configures gookit/validate scenes.
func (SysConfigValueRequest) ConfigValidation(v *validate.Validation) {
	v.WithScenes(validate.SValues{
		"insert": []string{"ConfigId", "Title", "Code", "Value", "Options", "Type", "Sort", "Note", "Disabled"},
		"update": []string{"Id", "ConfigId", "Title", "Code", "Value", "Options", "Type", "Sort", "Note", "Disabled"},
		"delete": []string{"Id"},
	})
}

// Messages defines Chinese validation messages.
func (SysConfigValueRequest) Messages() map[string]string {
	return validate.MS{
		"int":      "{field}必须是整数",
		"maxLen":   "{field}长度不能大于 %v",
		"required": "{field}不能为空",
		"string":   "{field}必须是字符串",
	}
}
