package validator

import (
	"github.com/gookit/validate"
	"time"
)

// SafeTodayThreatRequest is the request validator for the safe_today_threat table.
type SafeTodayThreatRequest struct {
	Id     int64     `json:"id" form:"id" xml:"id" url:"id" validate:"required|int" label:"主键"`
	Date   time.Time `json:"date" form:"date" xml:"date" url:"date" validate:"required" label:"日期"`
	High   int64     `json:"high" form:"high" xml:"high" url:"high" validate:"int" label:"高危事件数量"`
	Medium int64     `json:"medium" form:"medium" xml:"medium" url:"medium" validate:"int" label:"中危事件数量"`
	Low    int64     `json:"low" form:"low" xml:"low" url:"low" validate:"int" label:"低危事件数量"`
}

// ConfigValidation configures gookit/validate scenes.
func (SafeTodayThreatRequest) ConfigValidation(v *validate.Validation) {
	v.WithScenes(validate.SValues{
		"insert": []string{"Date", "High", "Medium", "Low"},
		"update": []string{"Id", "Date", "High", "Medium", "Low"},
		"delete": []string{"Id"},
	})
}

// Messages defines Chinese validation messages.
func (SafeTodayThreatRequest) Messages() map[string]string {
	return validate.MS{
		"int":      "{field}必须是整数",
		"required": "{field}不能为空",
	}
}
