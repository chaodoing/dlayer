package validator

import (
	"github.com/gookit/validate"
	"time"
)

// SafeTaskPlanUtsRequest is the request validator for the safe_task_plan_uts table.
type SafeTaskPlanUtsRequest struct {
	Id     uint      `json:"id" form:"id" xml:"id" url:"id" validate:"required|uint" label:"主键"`
	UtsId  uint      `json:"uts_id" form:"uts_id" xml:"uts_id" url:"uts_id" validate:"required|uint" label:"探针ID"`
	TaskId uint      `json:"task_id" form:"task_id" xml:"task_id" url:"task_id" validate:"required|uint" label:"任务id"`
	Stime  time.Time `json:"stime" form:"stime" xml:"stime" url:"stime" validate:"required" label:"开始时间"`
	Etime  time.Time `json:"etime" form:"etime" xml:"etime" url:"etime" validate:"required" label:"结束时间"`
	Issued uint8     `json:"issued" form:"issued" xml:"issued" url:"issued" validate:"uint" label:"下发状态"`
	Pcap   string    `json:"pcap" form:"pcap" xml:"pcap" url:"pcap" validate:"required|string|maxLen:255" label:"文件名称"`
}

// ConfigValidation configures gookit/validate scenes.
func (SafeTaskPlanUtsRequest) ConfigValidation(v *validate.Validation) {
	v.WithScenes(validate.SValues{
		"insert": []string{"UtsId", "TaskId", "Stime", "Etime", "Issued", "Pcap"},
		"update": []string{"Id", "UtsId", "TaskId", "Stime", "Etime", "Issued", "Pcap"},
		"delete": []string{"Id"},
	})
}

// Messages defines Chinese validation messages.
func (SafeTaskPlanUtsRequest) Messages() map[string]string {
	return validate.MS{
		"maxLen":   "{field}长度不能大于 %v",
		"required": "{field}不能为空",
		"string":   "{field}必须是字符串",
		"uint":     "{field}必须是非负整数",
	}
}
