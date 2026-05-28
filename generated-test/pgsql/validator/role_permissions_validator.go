package validator

import (
	"github.com/gookit/validate"
)

// RolePermissionsRequest is the request validator for the role_permissions table.
type RolePermissionsRequest struct {
	RoleId int `json:"role_id" form:"role_id" xml:"role_id" url:"role_id" validate:"required|int" label:"角色ID"`
	RuleId int `json:"rule_id" form:"rule_id" xml:"rule_id" url:"rule_id" validate:"required|int" label:"规则ID"`
}

// ConfigValidation configures gookit/validate scenes.
func (RolePermissionsRequest) ConfigValidation(v *validate.Validation) {
	v.WithScenes(validate.SValues{
		"insert": []string{"RoleId", "RuleId"},
		"update": []string{"RoleId", "RuleId"},
		"delete": []string{},
	})
}

// Messages defines Chinese validation messages.
func (RolePermissionsRequest) Messages() map[string]string {
	return validate.MS{
		"int":      "{field}必须是整数",
		"required": "{field}不能为空",
	}
}
