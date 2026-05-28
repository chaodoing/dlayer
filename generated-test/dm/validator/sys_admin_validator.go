package validator

import (
	"github.com/gookit/validate"
)

// SysAdminRequest is the request validator for the sys_admin table.
type SysAdminRequest struct {
	Id       int64  `json:"id" form:"id" xml:"id" url:"id" validate:"required|int" label:"主键"`
	RoleId   int64  `json:"role_id" form:"role_id" xml:"role_id" url:"role_id" validate:"required|int" label:"角色 id"`
	Name     string `json:"name" form:"name" xml:"name" url:"name" validate:"required|string|maxLen:48" label:"用户姓名"`
	Portrait string `json:"portrait" form:"portrait" xml:"portrait" url:"portrait" validate:"required|string|maxLen:1020" label:"用户头像"`
	Account  string `json:"account" form:"account" xml:"account" url:"account" validate:"required|string|maxLen:240" label:"登录账号"`
	Password string `json:"password" form:"password" xml:"password" url:"password" validate:"required|string|maxLen:256" label:"登录密码"`
	Method   string `json:"method" form:"method" xml:"method" url:"method" validate:"required|string|maxLen:24" label:"密码加密方法"`
	Salt     string `json:"salt" form:"salt" xml:"salt" url:"salt" validate:"required|string|maxLen:128" label:"认证器密钥"`
	Disabled int    `json:"disabled" form:"disabled" xml:"disabled" url:"disabled" validate:"int" label:"用户锁定状态"`
}

// ConfigValidation configures gookit/validate scenes.
func (SysAdminRequest) ConfigValidation(v *validate.Validation) {
	v.WithScenes(validate.SValues{
		"insert": []string{"RoleId", "Name", "Portrait", "Account", "Password", "Method", "Salt", "Disabled"},
		"update": []string{"Id", "RoleId", "Name", "Portrait", "Account", "Password", "Method", "Salt", "Disabled"},
		"delete": []string{"Id"},
	})
}

// Messages defines Chinese validation messages.
func (SysAdminRequest) Messages() map[string]string {
	return validate.MS{
		"int":      "{field}必须是整数",
		"maxLen":   "{field}长度不能大于 %v",
		"required": "{field}不能为空",
		"string":   "{field}必须是字符串",
	}
}
