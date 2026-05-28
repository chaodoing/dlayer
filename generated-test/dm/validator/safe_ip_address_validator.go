package validator

import (
	"github.com/gookit/validate"
)

// SafeIpAddressRequest is the request validator for the safe_ip_address table.
type SafeIpAddressRequest struct {
	Id        int64   `json:"id" form:"id" xml:"id" url:"id" validate:"required|int" label:"主键"`
	Address   string  `json:"address" form:"address" xml:"address" url:"address" validate:"required|string|maxLen:200" label:"IP地址"`
	IpVersion int8    `json:"ip_version" form:"ip_version" xml:"ip_version" url:"ip_version" validate:"required|int" label:"4=IPv4;6=ipv6"`
	Country   string  `json:"country" form:"country" xml:"country" url:"country" validate:"required|string|maxLen:200" label:"国家"`
	Province  string  `json:"province" form:"province" xml:"province" url:"province" validate:"required|string|maxLen:200" label:"省"`
	City      string  `json:"city" form:"city" xml:"city" url:"city" validate:"required|string|maxLen:200" label:"市"`
	Isp       string  `json:"isp" form:"isp" xml:"isp" url:"isp" validate:"required|string|maxLen:200" label:"运营商"`
	Lon       float64 `json:"lon" form:"lon" xml:"lon" url:"lon" validate:"required|float" label:"经度"`
	Lat       float64 `json:"lat" form:"lat" xml:"lat" url:"lat" validate:"required|float" label:"维度"`
}

// ConfigValidation configures gookit/validate scenes.
func (SafeIpAddressRequest) ConfigValidation(v *validate.Validation) {
	v.WithScenes(validate.SValues{
		"insert": []string{"Address", "IpVersion", "Country", "Province", "City", "Isp", "Lon", "Lat"},
		"update": []string{"Id", "Address", "IpVersion", "Country", "Province", "City", "Isp", "Lon", "Lat"},
		"delete": []string{"Id"},
	})
}

// Messages defines Chinese validation messages.
func (SafeIpAddressRequest) Messages() map[string]string {
	return validate.MS{
		"float":    "{field}必须是浮点数",
		"int":      "{field}必须是整数",
		"maxLen":   "{field}长度不能大于 %v",
		"required": "{field}不能为空",
		"string":   "{field}必须是字符串",
	}
}
