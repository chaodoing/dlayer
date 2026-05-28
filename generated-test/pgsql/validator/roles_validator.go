package validator

import (
	"github.com/gookit/validate"
)

// RolesRequest is the request validator for the roles table.
type RolesRequest struct {
	Id          int    `json:"id" form:"id" xml:"id" url:"id" validate:"required|int" label:"主键"`
	Title       string `json:"title" form:"title" xml:"title" url:"title" validate:"required|string|maxLen:60" label:"角色名称"`
	Name        string `json:"name" form:"name" xml:"name" url:"name" validate:"required|string|maxLen:60" label:"调用名称"`
	Description string `json:"description" form:"description" xml:"description" url:"description" validate:"required|string|maxLen:120" label:"描述信息"`
	Sort        int    `json:"sort" form:"sort" xml:"sort" url:"sort" validate:"required|int" label:"排序"`
	Disabled    int    `json:"disabled" form:"disabled" xml:"disabled" url:"disabled" validate:"required|int" label:"角色状态"`
}

// ConfigValidation configures gookit/validate scenes.
func (RolesRequest) ConfigValidation(v *validate.Validation) {
	v.WithScenes(validate.SValues{
		"insert": []string{"Title", "Name", "Description", "Sort", "Disabled"},
		"update": []string{"Id", "Title", "Name", "Description", "Sort", "Disabled"},
		"delete": []string{"Id"},
	})
}

// Messages defines Chinese validation messages.
func (RolesRequest) Messages() map[string]string {
	return validate.MS{
		"int":      "{field}必须是整数",
		"maxLen":   "{field}长度不能大于 %v",
		"required": "{field}不能为空",
		"string":   "{field}必须是字符串",
	}
}
