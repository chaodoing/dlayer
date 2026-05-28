package validator

import (
	"github.com/gookit/validate"
)

// SafeUtsDictionaryRequest is the request validator for the safe_uts_dictionary table.
type SafeUtsDictionaryRequest struct {
	Id         uint   `json:"id" form:"id" xml:"id" url:"id" validate:"required|uint" label:"主键"`
	Name       string `json:"name" form:"name" xml:"name" url:"name" validate:"required|string|maxLen:20" label:"事件名称"`
	AlertType  string `json:"alert_type" form:"alert_type" xml:"alert_type" url:"alert_type" validate:"required|string|maxLen:16" label:"告警分类编码"`
	AlertLevel uint8  `json:"alert_level" form:"alert_level" xml:"alert_level" url:"alert_level" validate:"required|uint" label:"告警事件等级"`
	EventType  string `json:"event_type" form:"event_type" xml:"event_type" url:"event_type" validate:"required|string|maxLen:16" label:"告警事件编码"`
}

// ConfigValidation configures gookit/validate scenes.
func (SafeUtsDictionaryRequest) ConfigValidation(v *validate.Validation) {
	v.WithScenes(validate.SValues{
		"insert": []string{"Name", "AlertType", "AlertLevel", "EventType"},
		"update": []string{"Id", "Name", "AlertType", "AlertLevel", "EventType"},
		"delete": []string{"Id"},
	})
}

// Messages defines Chinese validation messages.
func (SafeUtsDictionaryRequest) Messages() map[string]string {
	return validate.MS{
		"maxLen":   "{field}长度不能大于 %v",
		"required": "{field}不能为空",
		"string":   "{field}必须是字符串",
		"uint":     "{field}必须是非负整数",
	}
}
