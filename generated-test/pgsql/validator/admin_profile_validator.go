package validator

import (
	"github.com/gookit/validate"
)

// AdminProfileRequest is the request validator for the admin_profile table.
type AdminProfileRequest struct {
	AdminId          int    `json:"admin_id" form:"admin_id" xml:"admin_id" url:"admin_id" validate:"required|int" label:"管理员 id"`
	Email            string `json:"email" form:"email" xml:"email" url:"email" validate:"required|string|maxLen:128|email" label:"邮箱"`
	Nickname         string `json:"nickname" form:"nickname" xml:"nickname" url:"nickname" validate:"required|string|maxLen:32" label:"昵称"`
	Country          string `json:"country" form:"country" xml:"country" url:"country" validate:"required|string|maxLen:16" label:"国家/地区代码"`
	Region           string `json:"region" form:"region" xml:"region" url:"region" validate:"required|string|maxLen:32" label:"区域代码"`
	Address          string `json:"address" form:"address" xml:"address" url:"address" validate:"required|string|maxLen:255" label:"具体地址"`
	Bio              string `json:"bio" form:"bio" xml:"bio" url:"bio" validate:"required|string|maxLen:500" label:"个人简介"`
	RealNameVerified int    `json:"real_name_verified" form:"real_name_verified" xml:"real_name_verified" url:"real_name_verified" validate:"int" label:"实名认证 0未认证 1已认证"`
}

// ConfigValidation configures gookit/validate scenes.
func (AdminProfileRequest) ConfigValidation(v *validate.Validation) {
	v.WithScenes(validate.SValues{
		"insert": []string{"AdminId", "Email", "Nickname", "Country", "Region", "Address", "Bio", "RealNameVerified"},
		"update": []string{"AdminId", "Email", "Nickname", "Country", "Region", "Address", "Bio", "RealNameVerified"},
		"delete": []string{},
	})
}

// Messages defines Chinese validation messages.
func (AdminProfileRequest) Messages() map[string]string {
	return validate.MS{
		"email":    "{field}格式不正确",
		"int":      "{field}必须是整数",
		"maxLen":   "{field}长度不能大于 %v",
		"required": "{field}不能为空",
		"string":   "{field}必须是字符串",
	}
}
