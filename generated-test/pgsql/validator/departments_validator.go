package validator

import (
	"github.com/gookit/validate"
)

// DepartmentsRequest is the request validator for the departments table.
type DepartmentsRequest struct {
	Id            int    `json:"id" form:"id" xml:"id" url:"id" validate:"required|int" label:"部门ID"`
	ParentId      *int   `json:"parent_id" form:"parent_id" xml:"parent_id" url:"parent_id" validate:"int" label:"上级部门ID"`
	LeaderAdminId int    `json:"leader_admin_id" form:"leader_admin_id" xml:"leader_admin_id" url:"leader_admin_id" validate:"required|int" label:"部门后台负责人"`
	RoleId        int    `json:"role_id" form:"role_id" xml:"role_id" url:"role_id" validate:"required|int" label:"角色id"`
	Name          string `json:"name" form:"name" xml:"name" url:"name" validate:"required|string|maxLen:50" label:"部门名称"`
	Path          string `json:"path" form:"path" xml:"path" url:"path" validate:"required|string|maxLen:50" label:"部门层级路径"`
	Level         int    `json:"level" form:"level" xml:"level" url:"level" validate:"required|int" label:"部门层级"`
	Sort          int    `json:"sort" form:"sort" xml:"sort" url:"sort" validate:"required|int" label:"部门排序"`
	Disabled      bool   `json:"disabled" form:"disabled" xml:"disabled" url:"disabled" validate:"bool" label:"禁用状态"`
}

// ConfigValidation configures gookit/validate scenes.
func (DepartmentsRequest) ConfigValidation(v *validate.Validation) {
	v.WithScenes(validate.SValues{
		"insert": []string{"ParentId", "LeaderAdminId", "RoleId", "Name", "Path", "Level", "Sort", "Disabled"},
		"update": []string{"Id", "ParentId", "LeaderAdminId", "RoleId", "Name", "Path", "Level", "Sort", "Disabled"},
		"delete": []string{"Id"},
	})
}

// Messages defines Chinese validation messages.
func (DepartmentsRequest) Messages() map[string]string {
	return validate.MS{
		"bool":     "{field}必须是布尔值",
		"int":      "{field}必须是整数",
		"maxLen":   "{field}长度不能大于 %v",
		"required": "{field}不能为空",
		"string":   "{field}必须是字符串",
	}
}
