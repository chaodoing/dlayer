package validator

import (
	"github.com/gookit/validate"
	"time"
)

// TaskPlacesRequest is the request validator for the task_places table.
type TaskPlacesRequest struct {
	TaskId        int        `json:"task_id" form:"task_id" xml:"task_id" url:"task_id" validate:"required|int" label:"任务ID"`
	PlaceId       int        `json:"place_id" form:"place_id" xml:"place_id" url:"place_id" validate:"required|int" label:"场所ID"`
	PingCount     int        `json:"ping_count" form:"ping_count" xml:"ping_count" url:"ping_count" validate:"required|int" label:"场所测试次数"`
	CompleteCount int        `json:"complete_count" form:"complete_count" xml:"complete_count" url:"complete_count" validate:"int" label:"完成次数"`
	CompletedAt   *time.Time `json:"completed_at" form:"completed_at" xml:"completed_at" url:"completed_at" label:"完成时间"`
}

// ConfigValidation configures gookit/validate scenes.
func (TaskPlacesRequest) ConfigValidation(v *validate.Validation) {
	v.WithScenes(validate.SValues{
		"insert": []string{"TaskId", "PlaceId", "PingCount", "CompleteCount", "CompletedAt"},
		"update": []string{"TaskId", "PlaceId", "PingCount", "CompleteCount", "CompletedAt"},
		"delete": []string{},
	})
}

// Messages defines Chinese validation messages.
func (TaskPlacesRequest) Messages() map[string]string {
	return validate.MS{
		"int":      "{field}必须是整数",
		"required": "{field}不能为空",
	}
}
