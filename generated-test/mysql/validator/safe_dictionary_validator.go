package validator

import (
	"github.com/gookit/validate"
)

// SafeDictionaryRequest is the request validator for the safe_dictionary table.
type SafeDictionaryRequest struct {
	Id        uint   `json:"id" form:"id" xml:"id" url:"id" validate:"required|uint" label:"主键"`
	ParentId  uint   `json:"parent_id" form:"parent_id" xml:"parent_id" url:"parent_id" validate:"uint" label:"上级告警事件ID"`
	EventType string `json:"event_type" form:"event_type" xml:"event_type" url:"event_type" validate:"required|string|maxLen:16" label:"告警事件编码"`
	EventName string `json:"event_name" form:"event_name" xml:"event_name" url:"event_name" validate:"required|string|maxLen:30" label:"告警事件名称"`
	Disabled  uint8  `json:"disabled" form:"disabled" xml:"disabled" url:"disabled" validate:"uint" label:"禁用状态"`
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
		"maxLen":   "{field}长度不能大于 %v",
		"required": "{field}不能为空",
		"string":   "{field}必须是字符串",
		"uint":     "{field}必须是非负整数",
	}
}
