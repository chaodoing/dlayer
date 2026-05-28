package validator

import (
	"github.com/gookit/validate"
)

// TaskOperatorsRequest is the request validator for the task_operators table.
type TaskOperatorsRequest struct {
	TaskId   int    `json:"task_id" form:"task_id" xml:"task_id" url:"task_id" validate:"required|int" label:"主键"`
	Operator string `json:"operator" form:"operator" xml:"operator" url:"operator" validate:"required|string|maxLen:50" label:"运营商"`
}

// ConfigValidation configures gookit/validate scenes.
func (TaskOperatorsRequest) ConfigValidation(v *validate.Validation) {
	v.WithScenes(validate.SValues{
		"insert": []string{"TaskId", "Operator"},
		"update": []string{"TaskId", "Operator"},
		"delete": []string{},
	})
}

// Messages defines Chinese validation messages.
func (TaskOperatorsRequest) Messages() map[string]string {
	return validate.MS{
		"int":      "{field}必须是整数",
		"maxLen":   "{field}长度不能大于 %v",
		"required": "{field}不能为空",
		"string":   "{field}必须是字符串",
	}
}
