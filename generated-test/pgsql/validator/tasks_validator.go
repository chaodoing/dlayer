package validator

import (
	"github.com/gookit/validate"
	"time"
)

// TasksRequest is the request validator for the tasks table.
type TasksRequest struct {
	Id           int       `json:"id" form:"id" xml:"id" url:"id" validate:"required|int" label:"主键"`
	ParentId     *int      `json:"parent_id" form:"parent_id" xml:"parent_id" url:"parent_id" validate:"int" label:"上级ID"`
	DepartmentId int       `json:"department_id" form:"department_id" xml:"department_id" url:"department_id" validate:"required|int" label:"所属部门ID"`
	PlaceListId  int       `json:"place_list_id" form:"place_list_id" xml:"place_list_id" url:"place_list_id" validate:"required|int" label:"场所清单id"`
	Name         string    `json:"name" form:"name" xml:"name" url:"name" validate:"required|string|maxLen:50" label:"任务名称"`
	StartTime    time.Time `json:"start_time" form:"start_time" xml:"start_time" url:"start_time" validate:"required" label:"开始时间"`
	DoneTime     time.Time `json:"done_time" form:"done_time" xml:"done_time" url:"done_time" validate:"required" label:"完成时间"`
	HasVideo     bool      `json:"has_video" form:"has_video" xml:"has_video" url:"has_video" validate:"bool" label:"是否上传视频"`
	HasImage     bool      `json:"has_image" form:"has_image" xml:"has_image" url:"has_image" validate:"bool" label:"是否上传图片"`
}

// ConfigValidation configures gookit/validate scenes.
func (TasksRequest) ConfigValidation(v *validate.Validation) {
	v.WithScenes(validate.SValues{
		"insert": []string{"ParentId", "DepartmentId", "PlaceListId", "Name", "StartTime", "DoneTime", "HasVideo", "HasImage"},
		"update": []string{"Id", "ParentId", "DepartmentId", "PlaceListId", "Name", "StartTime", "DoneTime", "HasVideo", "HasImage"},
		"delete": []string{"Id"},
	})
}

// Messages defines Chinese validation messages.
func (TasksRequest) Messages() map[string]string {
	return validate.MS{
		"bool":     "{field}必须是布尔值",
		"int":      "{field}必须是整数",
		"maxLen":   "{field}长度不能大于 %v",
		"required": "{field}不能为空",
		"string":   "{field}必须是字符串",
	}
}
