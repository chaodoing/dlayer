package validator

import (
	"github.com/gookit/validate"
	"time"
)

// SafeTodayAlarmRequest is the request validator for the safe_today_alarm table.
type SafeTodayAlarmRequest struct {
	Id        int64     `json:"id" form:"id" xml:"id" url:"id" validate:"required|int" label:"主键"`
	Date      time.Time `json:"date" form:"date" xml:"date" url:"date" validate:"required" label:"日期"`
	EventType string    `json:"event_type" form:"event_type" xml:"event_type" url:"event_type" validate:"required|string|maxLen:64" label:"事件ID"`
	Total     int64     `json:"total" form:"total" xml:"total" url:"total" validate:"int" label:"事件数量"`
}

// ConfigValidation configures gookit/validate scenes.
func (SafeTodayAlarmRequest) ConfigValidation(v *validate.Validation) {
	v.WithScenes(validate.SValues{
		"insert": []string{"Date", "EventType", "Total"},
		"update": []string{"Id", "Date", "EventType", "Total"},
		"delete": []string{"Id"},
	})
}

// Messages defines Chinese validation messages.
func (SafeTodayAlarmRequest) Messages() map[string]string {
	return validate.MS{
		"int":      "{field}必须是整数",
		"maxLen":   "{field}长度不能大于 %v",
		"required": "{field}不能为空",
		"string":   "{field}必须是字符串",
	}
}
