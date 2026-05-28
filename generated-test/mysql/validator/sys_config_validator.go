package validator

import (
	"github.com/gookit/validate"
)

// SysConfigRequest is the request validator for the sys_config table.
type SysConfigRequest struct {
	Id       uint32 `json:"id" form:"id" xml:"id" url:"id" validate:"required|uint" label:"主键"`
	Code     string `json:"code" form:"code" xml:"code" url:"code" validate:"required|string|maxLen:60" label:"调用编码"`
	Title    string `json:"title" form:"title" xml:"title" url:"title" validate:"required|string|maxLen:30" label:"配置名称"`
	Sort     uint32 `json:"sort" form:"sort" xml:"sort" url:"sort" validate:"required|uint" label:"显示顺序"`
	Disabled uint8  `json:"disabled" form:"disabled" xml:"disabled" url:"disabled" validate:"uint" label:"禁用状态"`
}

// ConfigValidation configures gookit/validate scenes.
func (SysConfigRequest) ConfigValidation(v *validate.Validation) {
	v.WithScenes(validate.SValues{
		"insert": []string{"Code", "Title", "Sort", "Disabled"},
		"update": []string{"Id", "Code", "Title", "Sort", "Disabled"},
		"delete": []string{"Id"},
	})
}

// Messages defines Chinese validation messages.
func (SysConfigRequest) Messages() map[string]string {
	return validate.MS{
		"maxLen":   "{field}长度不能大于 %v",
		"required": "{field}不能为空",
		"string":   "{field}必须是字符串",
		"uint":     "{field}必须是非负整数",
	}
}
