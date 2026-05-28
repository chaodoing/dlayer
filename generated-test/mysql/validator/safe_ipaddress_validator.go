package validator

import (
	"github.com/gookit/validate"
	"time"
)

// SafeIpaddressRequest is the request validator for the safe_ipaddress table.
type SafeIpaddressRequest struct {
	Id             uint64    `json:"id" form:"id" xml:"id" url:"id" validate:"required|uint" label:"主键"`
	CompanyId      uint64    `json:"company_id" form:"company_id" xml:"company_id" url:"company_id" validate:"required|uint" label:"关联company_info.id"`
	IpAddress      string    `json:"ip_address" form:"ip_address" xml:"ip_address" url:"ip_address" validate:"required|string|maxLen:45" label:"IP地址（兼容IPv4/IPv6）"`
	IpVersion      uint8     `json:"ip_version" form:"ip_version" xml:"ip_version" url:"ip_version" validate:"uint" label:"4=IPv4, 6=IPv6"`
	IpType         uint8     `json:"ip_type" form:"ip_type" xml:"ip_type" url:"ip_type" validate:"required|uint" label:"1=官网,2=API,3=办公出口,4=CDN,5=数据库"`
	Purpose        string    `json:"purpose" form:"purpose" xml:"purpose" url:"purpose" validate:"required|string|maxLen:100" label:"用途说明"`
	IsActive       uint8     `json:"is_active" form:"is_active" xml:"is_active" url:"is_active" validate:"uint" label:"是否启用"`
	LastVerifiedAt time.Time `json:"last_verified_at" form:"last_verified_at" xml:"last_verified_at" url:"last_verified_at" validate:"required" label:"最后验证时间"`
	VerifiedBy     uint64    `json:"verified_by" form:"verified_by" xml:"verified_by" url:"verified_by" validate:"required|uint" label:"验证人"`
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
		"maxLen":   "{field}长度不能大于 %v",
		"required": "{field}不能为空",
		"string":   "{field}必须是字符串",
		"uint":     "{field}必须是非负整数",
	}
}
