package validator

import (
	"github.com/gookit/validate"
	"gorm.io/datatypes"
)

// SysRoleRequest is the request validator for the sys_role table.
type SysRoleRequest struct {
	Id          uint           `json:"id" form:"id" xml:"id" url:"id" validate:"required|uint" label:"主键"`
	Title       string         `json:"title" form:"title" xml:"title" url:"title" validate:"required|string|maxLen:60" label:"角色名称"`
	Name        string         `json:"name" form:"name" xml:"name" url:"name" validate:"required|string|maxLen:60" label:"调用名称"`
	Description string         `json:"description" form:"description" xml:"description" url:"description" validate:"required|string|maxLen:120" label:"描述信息"`
	Rules       datatypes.JSON `json:"rules" form:"rules" xml:"rules" url:"rules" validate:"required" label:"角色允许访问"`
	Sort        uint           `json:"sort" form:"sort" xml:"sort" url:"sort" validate:"required|uint" label:"排序"`
	Disabled    uint8          `json:"disabled" form:"disabled" xml:"disabled" url:"disabled" validate:"required|uint" label:"角色状态"`
}

// ConfigValidation configures gookit/validate scenes.
func (SysRoleRequest) ConfigValidation(v *validate.Validation) {
	v.WithScenes(validate.SValues{
		"insert": []string{"Title", "Name", "Description", "Rules", "Sort", "Disabled"},
		"update": []string{"Id", "Title", "Name", "Description", "Rules", "Sort", "Disabled"},
		"delete": []string{"Id"},
	})
}

// Messages defines Chinese validation messages.
func (SysRoleRequest) Messages() map[string]string {
	return validate.MS{
		"maxLen":   "{field}长度不能大于 %v",
		"required": "{field}不能为空",
		"string":   "{field}必须是字符串",
		"uint":     "{field}必须是非负整数",
	}
}
