package validator

import (
	"github.com/gookit/validate"
)

// SysRegionRequest is the request validator for the sys_region table.
type SysRegionRequest struct {
	Code       uint32 `json:"code" form:"code" xml:"code" url:"code" validate:"required|uint" label:"城市编码"`
	Name       string `json:"name" form:"name" xml:"name" url:"name" validate:"required|string|maxLen:255" label:"城市名称"`
	ParentCode uint32 `json:"parent_code" form:"parent_code" xml:"parent_code" url:"parent_code" validate:"required|uint" label:"上级城市"`
	Level      uint8  `json:"level" form:"level" xml:"level" url:"level" validate:"uint" label:"城市等级"`
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
		"maxLen":   "{field}长度不能大于 %v",
		"required": "{field}不能为空",
		"string":   "{field}必须是字符串",
		"uint":     "{field}必须是非负整数",
	}
}
