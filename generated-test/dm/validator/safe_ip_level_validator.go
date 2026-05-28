package validator

import (
	"github.com/gookit/validate"
	"time"
)

// SafeIpLevelRequest is the request validator for the safe_ip_level table.
type SafeIpLevelRequest struct {
	Id        int64     `json:"id" form:"id" xml:"id" url:"id" validate:"required|int" label:"主键"`
	AddressId int64     `json:"address_id" form:"address_id" xml:"address_id" url:"address_id" validate:"required|int" label:"IP地址ID"`
	Date      time.Time `json:"date" form:"date" xml:"date" url:"date" validate:"required" label:"日期"`
	High      int       `json:"high" form:"high" xml:"high" url:"high" validate:"required|int" label:"高"`
	Medium    int       `json:"medium" form:"medium" xml:"medium" url:"medium" validate:"required|int" label:"中"`
	Low       int       `json:"low" form:"low" xml:"low" url:"low" validate:"required|int" label:"低"`
}

// ConfigValidation configures gookit/validate scenes.
func (SafeIpLevelRequest) ConfigValidation(v *validate.Validation) {
	v.WithScenes(validate.SValues{
		"insert": []string{"AddressId", "Date", "High", "Medium", "Low"},
		"update": []string{"Id", "AddressId", "Date", "High", "Medium", "Low"},
		"delete": []string{"Id"},
	})
}

// Messages defines Chinese validation messages.
func (SafeIpLevelRequest) Messages() map[string]string {
	return validate.MS{
		"int":      "{field}必须是整数",
		"required": "{field}不能为空",
	}
}
