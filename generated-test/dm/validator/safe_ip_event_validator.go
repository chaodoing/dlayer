package validator

import (
	"github.com/gookit/validate"
	"time"
)

// SafeIpEventRequest is the request validator for the safe_ip_event table.
type SafeIpEventRequest struct {
	Id        int64     `json:"id" form:"id" xml:"id" url:"id" validate:"required|int" label:"主键"`
	AddressId int64     `json:"address_id" form:"address_id" xml:"address_id" url:"address_id" validate:"required|int" label:"IP地址ID"`
	Date      time.Time `json:"date" form:"date" xml:"date" url:"date" validate:"required" label:"日期"`
	EventType string    `json:"event_type" form:"event_type" xml:"event_type" url:"event_type" validate:"required|string|maxLen:64" label:"事件编码"`
	Attacks   int       `json:"attacks" form:"attacks" xml:"attacks" url:"attacks" validate:"required|int" label:"攻击次数"`
}

// ConfigValidation configures gookit/validate scenes.
func (SafeIpEventRequest) ConfigValidation(v *validate.Validation) {
	v.WithScenes(validate.SValues{
		"insert": []string{"AddressId", "Date", "EventType", "Attacks"},
		"update": []string{"Id", "AddressId", "Date", "EventType", "Attacks"},
		"delete": []string{"Id"},
	})
}

// Messages defines Chinese validation messages.
func (SafeIpEventRequest) Messages() map[string]string {
	return validate.MS{
		"int":      "{field}必须是整数",
		"maxLen":   "{field}长度不能大于 %v",
		"required": "{field}不能为空",
		"string":   "{field}必须是字符串",
	}
}
