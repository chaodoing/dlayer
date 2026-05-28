package validator

import (
	"github.com/gookit/validate"
)

// SafeTaskPlanRequest is the request validator for the safe_task_plan table.
type SafeTaskPlanRequest struct {
	Id       uint   `json:"id" form:"id" xml:"id" url:"id" validate:"required|uint" label:"主键"`
	Title    string `json:"title" form:"title" xml:"title" url:"title" validate:"required|string|maxLen:60" label:"任务计划标题"`
	Sip      string `json:"sip" form:"sip" xml:"sip" url:"sip" validate:"required|string|maxLen:30" label:"来源IP"`
	Sport    uint16 `json:"sport" form:"sport" xml:"sport" url:"sport" validate:"required|uint" label:"来源端口"`
	Dip      string `json:"dip" form:"dip" xml:"dip" url:"dip" validate:"required|string|maxLen:30" label:"目标IP"`
	Dport    uint16 `json:"dport" form:"dport" xml:"dport" url:"dport" validate:"required|uint" label:"目标端口"`
	Size     uint16 `json:"size" form:"size" xml:"size" url:"size" validate:"required|uint" label:"数据包大小"`
	Duration uint16 `json:"duration" form:"duration" xml:"duration" url:"duration" validate:"required|uint" label:"采集时长"`
	Disabled uint8  `json:"disabled" form:"disabled" xml:"disabled" url:"disabled" validate:"uint" label:"禁用状态"`
	Issue    uint16 `json:"issue" form:"issue" xml:"issue" url:"issue" validate:"uint" label:"下发的任务数量"`
	Unissue  uint16 `json:"unissue" form:"unissue" xml:"unissue" url:"unissue" validate:"uint" label:"未下发的任务数量"`
	Timeout  uint16 `json:"timeout" form:"timeout" xml:"timeout" url:"timeout" validate:"uint" label:"过期的任务数量"`
	Total    uint16 `json:"total" form:"total" xml:"total" url:"total" validate:"uint" label:"总计任务数量"`
	Sort     uint16 `json:"sort" form:"sort" xml:"sort" url:"sort" validate:"uint" label:"任务排序"`
}

// ConfigValidation configures gookit/validate scenes.
func (SafeTaskPlanRequest) ConfigValidation(v *validate.Validation) {
	v.WithScenes(validate.SValues{
		"insert": []string{"Title", "Sip", "Sport", "Dip", "Dport", "Size", "Duration", "Disabled", "Issue", "Unissue", "Timeout", "Total", "Sort"},
		"update": []string{"Id", "Title", "Sip", "Sport", "Dip", "Dport", "Size", "Duration", "Disabled", "Issue", "Unissue", "Timeout", "Total", "Sort"},
		"delete": []string{"Id"},
	})
}

// Messages defines Chinese validation messages.
func (SafeTaskPlanRequest) Messages() map[string]string {
	return validate.MS{
		"maxLen":   "{field}长度不能大于 %v",
		"required": "{field}不能为空",
		"string":   "{field}必须是字符串",
		"uint":     "{field}必须是非负整数",
	}
}
