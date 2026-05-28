package validator

import (
	"github.com/gookit/validate"
)

// TaskNetworkTypesRequest is the request validator for the task_network_types table.
type TaskNetworkTypesRequest struct {
	TaskId      int    `json:"task_id" form:"task_id" xml:"task_id" url:"task_id" validate:"required|int" label:"任务id"`
	NetworkType string `json:"network_type" form:"network_type" xml:"network_type" url:"network_type" validate:"required|string|maxLen:50" label:"需要测试的网络类型"`
}

// ConfigValidation configures gookit/validate scenes.
func (TaskNetworkTypesRequest) ConfigValidation(v *validate.Validation) {
	v.WithScenes(validate.SValues{
		"insert": []string{"TaskId", "NetworkType"},
		"update": []string{"TaskId", "NetworkType"},
		"delete": []string{},
	})
}

// Messages defines Chinese validation messages.
func (TaskNetworkTypesRequest) Messages() map[string]string {
	return validate.MS{
		"int":      "{field}必须是整数",
		"maxLen":   "{field}长度不能大于 %v",
		"required": "{field}不能为空",
		"string":   "{field}必须是字符串",
	}
}
