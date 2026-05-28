package validator

import (
	"github.com/gookit/validate"
)

// DepartmentOperatorsRequest is the request validator for the department_operators table.
type DepartmentOperatorsRequest struct {
	DepartmentId int    `json:"department_id" form:"department_id" xml:"department_id" url:"department_id" validate:"required|int" label:"部门ID"`
	Operator     string `json:"operator" form:"operator" xml:"operator" url:"operator" validate:"required|string|maxLen:50" label:"运营商"`
}

// ConfigValidation configures gookit/validate scenes.
func (DepartmentOperatorsRequest) ConfigValidation(v *validate.Validation) {
	v.WithScenes(validate.SValues{
		"insert": []string{"DepartmentId", "Operator"},
		"update": []string{"DepartmentId", "Operator"},
		"delete": []string{},
	})
}

// Messages defines Chinese validation messages.
func (DepartmentOperatorsRequest) Messages() map[string]string {
	return validate.MS{
		"int":      "{field}必须是整数",
		"maxLen":   "{field}长度不能大于 %v",
		"required": "{field}不能为空",
		"string":   "{field}必须是字符串",
	}
}
