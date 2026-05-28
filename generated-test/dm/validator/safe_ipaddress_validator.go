package validator

import (
	"github.com/gookit/validate"
	"time"
)

// SafeIpaddressRequest is the request validator for the safe_ipaddress table.
type SafeIpaddressRequest struct {
	Id             int64     `json:"id" form:"id" xml:"id" url:"id" validate:"required|int" label:"主键"`
	CompanyId      float64   `json:"company_id" form:"company_id" xml:"company_id" url:"company_id" validate:"required|float" label:"关联company_info.id"`
	IpAddress      string    `json:"ip_address" form:"ip_address" xml:"ip_address" url:"ip_address" validate:"required|string|maxLen:180" label:"IP地址（兼容IPv4/IPv6）"`
	IpVersion      int       `json:"ip_version" form:"ip_version" xml:"ip_version" url:"ip_version" validate:"int" label:"4=IPv4, 6=IPv6"`
	IpType         int       `json:"ip_type" form:"ip_type" xml:"ip_type" url:"ip_type" validate:"required|int" label:"1=官网,2=API,3=办公出口,4=CDN,5=数据库"`
	Purpose        string    `json:"purpose" form:"purpose" xml:"purpose" url:"purpose" validate:"required|string|maxLen:400" label:"用途说明"`
	IsActive       int       `json:"is_active" form:"is_active" xml:"is_active" url:"is_active" validate:"int" label:"是否启用"`
	LastVerifiedAt time.Time `json:"last_verified_at" form:"last_verified_at" xml:"last_verified_at" url:"last_verified_at" validate:"required" label:"最后验证时间"`
	VerifiedBy     float64   `json:"verified_by" form:"verified_by" xml:"verified_by" url:"verified_by" validate:"required|float" label:"验证人"`
}

// ConfigValidation configures gookit/validate scenes.
func (SafeIpaddressRequest) ConfigValidation(v *validate.Validation) {
	v.WithScenes(validate.SValues{
		"insert": []string{"CompanyId", "IpAddress", "IpVersion", "IpType", "Purpose", "IsActive", "LastVerifiedAt", "VerifiedBy"},
		"update": []string{"Id", "CompanyId", "IpAddress", "IpVersion", "IpType", "Purpose", "IsActive", "LastVerifiedAt", "VerifiedBy"},
		"delete": []string{"Id"},
	})
}

// Messages defines Chinese validation messages.
func (SafeIpaddressRequest) Messages() map[string]string {
	return validate.MS{
		"float":    "{field}必须是浮点数",
		"int":      "{field}必须是整数",
		"maxLen":   "{field}长度不能大于 %v",
		"required": "{field}不能为空",
		"string":   "{field}必须是字符串",
	}
}
