package validator

import (
	"github.com/gookit/validate"
)

// AdminRequest is the request validator for the admin table.
type AdminRequest struct {
	Id           int    `json:"id" form:"id" xml:"id" url:"id" validate:"required|int" label:"主键"`
	RoleId       int    `json:"role_id" form:"role_id" xml:"role_id" url:"role_id" validate:"required|int" label:"角色 id"`
	DepartmentId int    `json:"department_id" form:"department_id" xml:"department_id" url:"department_id" validate:"required|int" label:"部门ID"`
	Name         string `json:"name" form:"name" xml:"name" url:"name" validate:"required|string|maxLen:12" label:"用户姓名"`
	Avatar       string `json:"avatar" form:"avatar" xml:"avatar" url:"avatar" validate:"required|string|maxLen:255" label:"用户头像"`
	Account      string `json:"account" form:"account" xml:"account" url:"account" validate:"required|string|maxLen:60" label:"登录账号"`
	Password     string `json:"password" form:"password" xml:"password" url:"password" validate:"required|string|maxLen:64" label:"登录密码"`
	Method       string `json:"method" form:"method" xml:"method" url:"method" validate:"required|string|maxLen:10" label:"密码加密方法"`
	Salt         string `json:"salt" form:"salt" xml:"salt" url:"salt" validate:"required|string|maxLen:32" label:"认证器密钥"`
	Disabled     int    `json:"disabled" form:"disabled" xml:"disabled" url:"disabled" validate:"int" label:"用户锁定状态"`
}

// ConfigValidation configures gookit/validate scenes.
func (AdminRequest) ConfigValidation(v *validate.Validation) {
	v.WithScenes(validate.SValues{
		"insert": []string{"RoleId", "DepartmentId", "Name", "Avatar", "Account", "Password", "Method", "Salt", "Disabled"},
		"update": []string{"Id", "RoleId", "DepartmentId", "Name", "Avatar", "Account", "Password", "Method", "Salt", "Disabled"},
		"delete": []string{"Id"},
	})
}

// Messages defines Chinese validation messages.
func (AdminRequest) Messages() map[string]string {
	return validate.MS{
		"int":      "{field}必须是整数",
		"maxLen":   "{field}长度不能大于 %v",
		"required": "{field}不能为空",
		"string":   "{field}必须是字符串",
	}
}
