package validator

import (
	"github.com/gookit/validate"
)

// SafeTaskPlanRequest is the request validator for the safe_task_plan table.
type SafeTaskPlanRequest struct {
	Id       int    `json:"id" form:"id" xml:"id" url:"id" validate:"required|int" label:"主键"`
	Title    string `json:"title" form:"title" xml:"title" url:"title" validate:"required|string|maxLen:240" label:"任务计划标题"`
	Sip      string `json:"sip" form:"sip" xml:"sip" url:"sip" validate:"required|string|maxLen:120" label:"来源IP"`
	Sport    int16  `json:"sport" form:"sport" xml:"sport" url:"sport" validate:"required|int" label:"来源端口"`
	Dip      string `json:"dip" form:"dip" xml:"dip" url:"dip" validate:"required|string|maxLen:120" label:"目标IP"`
	Dport    int16  `json:"dport" form:"dport" xml:"dport" url:"dport" validate:"required|int" label:"目标端口"`
	Size     int16  `json:"size" form:"size" xml:"size" url:"size" validate:"required|int" label:"数据包大小"`
	Duration int16  `json:"duration" form:"duration" xml:"duration" url:"duration" validate:"required|int" label:"采集时长"`
	Disabled int16  `json:"disabled" form:"disabled" xml:"disabled" url:"disabled" validate:"int" label:"禁用状态"`
	Issue    int16  `json:"issue" form:"issue" xml:"issue" url:"issue" validate:"int" label:"下发的任务数量"`
	Unissue  int16  `json:"unissue" form:"unissue" xml:"unissue" url:"unissue" validate:"int" label:"未下发的任务数量"`
	Timeout  int16  `json:"timeout" form:"timeout" xml:"timeout" url:"timeout" validate:"int" label:"过期的任务数量"`
	Total    int16  `json:"total" form:"total" xml:"total" url:"total" validate:"int" label:"总计任务数量"`
	Sort     int16  `json:"sort" form:"sort" xml:"sort" url:"sort" validate:"int" label:"任务排序"`
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
		"int":      "{field}必须是整数",
		"maxLen":   "{field}长度不能大于 %v",
		"required": "{field}不能为空",
		"string":   "{field}必须是字符串",
	}
}
