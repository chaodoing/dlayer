package validator

import (
	"github.com/gookit/validate"
)

// SysRegionRequest is the request validator for the sys_region table.
type SysRegionRequest struct {
	Code       int    `json:"code" form:"code" xml:"code" url:"code" validate:"required|int" label:"城市编码"`
	Name       string `json:"name" form:"name" xml:"name" url:"name" validate:"required|string|maxLen:1020" label:"城市名称"`
	ParentCode int    `json:"parent_code" form:"parent_code" xml:"parent_code" url:"parent_code" validate:"required|int" label:"上级城市"`
	Level      int    `json:"level" form:"level" xml:"level" url:"level" validate:"int" label:"城市等级"`
}

// ConfigValidation configures gookit/validate scenes.
func (SysRegionRequest) ConfigValidation(v *validate.Validation) {
	v.WithScenes(validate.SValues{
		"insert": []string{"Code", "Name", "ParentCode", "Level"},
		"update": []string{"Code", "Name", "ParentCode", "Level"},
		"delete": []string{},
	})
}

// Messages defines Chinese validation messages.
func (SysRegionRequest) Messages() map[string]string {
	return validate.MS{
		"int":      "{field}必须是整数",
		"maxLen":   "{field}长度不能大于 %v",
		"required": "{field}不能为空",
		"string":   "{field}必须是字符串",
	}
}
