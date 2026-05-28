package validator

import (
	"github.com/gookit/validate"
	"time"
)

// UsersRequest is the request validator for the users table.
type UsersRequest struct {
	Id               int       `json:"id" form:"id" xml:"id" url:"id" validate:"required|int" label:"用户ID"`
	DepartmentId     int       `json:"department_id" form:"department_id" xml:"department_id" url:"department_id" validate:"required|int" label:"所属部门ID"`
	Type             int       `json:"type" form:"type" xml:"type" url:"type" validate:"required|int" label:"用户类型(1:众测,2:专测)"`
	Nickname         string    `json:"nickname" form:"nickname" xml:"nickname" url:"nickname" validate:"required|string|maxLen:50" label:"用户昵称"`
	Username         string    `json:"username" form:"username" xml:"username" url:"username" validate:"required|string|maxLen:50" label:"用户名"`
	Mobile           string    `json:"mobile" form:"mobile" xml:"mobile" url:"mobile" validate:"required|string|maxLen:18" label:"手机号码"`
	Password         string    `json:"password" form:"password" xml:"password" url:"password" validate:"required|string|maxLen:60" label:"登录密码"`
	RegisterAddress  string    `json:"register_address" form:"register_address" xml:"register_address" url:"register_address" validate:"required|string|maxLen:60" label:"注册IP"`
	RegisterTime     time.Time `json:"register_time" form:"register_time" xml:"register_time" url:"register_time" validate:"required" label:"注册时间"`
	LastLoginAddress string    `json:"last_login_address" form:"last_login_address" xml:"last_login_address" url:"last_login_address" validate:"required|string|maxLen:60" label:"登录IP"`
	LastLoginTime    time.Time `json:"last_login_time" form:"last_login_time" xml:"last_login_time" url:"last_login_time" validate:"required" label:"最后登录时间"`
}

// ConfigValidation configures gookit/validate scenes.
func (UsersRequest) ConfigValidation(v *validate.Validation) {
	v.WithScenes(validate.SValues{
		"insert": []string{"DepartmentId", "Type", "Nickname", "Username", "Mobile", "Password", "RegisterAddress", "RegisterTime", "LastLoginAddress", "LastLoginTime"},
		"update": []string{"Id", "DepartmentId", "Type", "Nickname", "Username", "Mobile", "Password", "RegisterAddress", "RegisterTime", "LastLoginAddress", "LastLoginTime"},
		"delete": []string{"Id"},
	})
}

// Messages defines Chinese validation messages.
func (UsersRequest) Messages() map[string]string {
	return validate.MS{
		"int":      "{field}必须是整数",
		"maxLen":   "{field}长度不能大于 %v",
		"required": "{field}不能为空",
		"string":   "{field}必须是字符串",
	}
}
