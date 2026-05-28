package validator

import (
	"github.com/gookit/validate"
)

// SafeDictionaryRequest is the request validator for the safe_dictionary table.
type SafeDictionaryRequest struct {
	Id        int64  `json:"id" form:"id" xml:"id" url:"id" validate:"required|int" label:"主键"`
	ParentId  int64  `json:"parent_id" form:"parent_id" xml:"parent_id" url:"parent_id" validate:"int" label:"上级告警事件ID"`
	EventType string `json:"event_type" form:"event_type" xml:"event_type" url:"event_type" validate:"required|string|maxLen:64" label:"告警事件编码"`
	EventName string `json:"event_name" form:"event_name" xml:"event_name" url:"event_name" validate:"required|string|maxLen:120" label:"告警事件名称"`
	Disabled  int    `json:"disabled" form:"disabled" xml:"disabled" url:"disabled" validate:"int" label:"禁用状态"`
}

// ConfigValidation configures gookit/validate scenes.
func (SafeDictionaryRequest) ConfigValidation(v *validate.Validation) {
	v.WithScenes(validate.SValues{
		"insert": []string{"ParentId", "EventType", "EventName", "Disabled"},
		"update": []string{"Id", "ParentId", "EventType", "EventName", "Disabled"},
		"delete": []string{"Id"},
	})
}

// Messages defines Chinese validation messages.
func (SafeDictionaryRequest) Messages() map[string]string {
	return validate.MS{
		"int":      "{field}必须是整数",
		"maxLen":   "{field}长度不能大于 %v",
		"required": "{field}不能为空",
		"string":   "{field}必须是字符串",
	}
}
