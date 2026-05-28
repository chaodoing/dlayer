package validator

import (
	"github.com/gookit/validate"
)

// TaskZonesRequest is the request validator for the task_zones table.
type TaskZonesRequest struct {
	TaskId    int `json:"task_id" form:"task_id" xml:"task_id" url:"task_id" validate:"required|int" label:"任务ID"`
	ZoneCode  int `json:"zone_code" form:"zone_code" xml:"zone_code" url:"zone_code" validate:"required|int" label:"地区编码"`
	ZoneLevel int `json:"zone_level" form:"zone_level" xml:"zone_level" url:"zone_level" validate:"required|int" label:"地区等级"`
}

// ConfigValidation configures gookit/validate scenes.
func (TaskZonesRequest) ConfigValidation(v *validate.Validation) {
	v.WithScenes(validate.SValues{
		"insert": []string{"TaskId", "ZoneCode", "ZoneLevel"},
		"update": []string{"TaskId", "ZoneCode", "ZoneLevel"},
		"delete": []string{},
	})
}

// Messages defines Chinese validation messages.
func (TaskZonesRequest) Messages() map[string]string {
	return validate.MS{
		"int":      "{field}必须是整数",
		"required": "{field}不能为空",
	}
}
