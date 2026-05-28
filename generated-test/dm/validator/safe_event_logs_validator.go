package validator

import (
	"github.com/gookit/validate"
	"time"
)

// SafeEventLogsRequest is the request validator for the safe_event_logs table.
type SafeEventLogsRequest struct {
	Id         int64     `json:"id" form:"id" xml:"id" url:"id" validate:"required|int" label:"主键"`
	Date       time.Time `json:"date" form:"date" xml:"date" url:"date" validate:"required" label:"日期"`
	EventType  string    `json:"event_type" form:"event_type" xml:"event_type" url:"event_type" validate:"required|string|maxLen:64" label:"所属事件"`
	PlatformId int64     `json:"platform_id" form:"platform_id" xml:"platform_id" url:"platform_id" validate:"required|int" label:"来源平台ID"`
	Total      int       `json:"total" form:"total" xml:"total" url:"total" validate:"required|int" label:"总计数量"`
	Successed  int       `json:"successed" form:"successed" xml:"successed" url:"successed" validate:"required|int" label:"成功数量"`
	Failed     int       `json:"failed" form:"failed" xml:"failed" url:"failed" validate:"required|int" label:"失败数量"`
}

// ConfigValidation configures gookit/validate scenes.
func (SafeEventLogsRequest) ConfigValidation(v *validate.Validation) {
	v.WithScenes(validate.SValues{
		"insert": []string{"Date", "EventType", "PlatformId", "Total", "Successed", "Failed"},
		"update": []string{"Id", "Date", "EventType", "PlatformId", "Total", "Successed", "Failed"},
		"delete": []string{"Id"},
	})
}

// Messages defines Chinese validation messages.
func (SafeEventLogsRequest) Messages() map[string]string {
	return validate.MS{
		"int":      "{field}必须是整数",
		"maxLen":   "{field}长度不能大于 %v",
		"required": "{field}不能为空",
		"string":   "{field}必须是字符串",
	}
}
