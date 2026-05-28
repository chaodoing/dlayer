package validator

import (
	"github.com/gookit/validate"
	"time"
)

// SafeCompanyRequest is the request validator for the safe_company table.
type SafeCompanyRequest struct {
	Id                  uint64    `json:"id" form:"id" xml:"id" url:"id" validate:"required|uint" label:"主键ID"`
	CompanyName         string    `json:"company_name" form:"company_name" xml:"company_name" url:"company_name" validate:"required|string|maxLen:200" label:"公司全称"`
	ShortName           string    `json:"short_name" form:"short_name" xml:"short_name" url:"short_name" validate:"required|string|maxLen:100" label:"公司简称"`
	CreditCode          string    `json:"credit_code" form:"credit_code" xml:"credit_code" url:"credit_code" validate:"required|string|maxLen:50" label:"统一社会信用代码（中国）/注册号"`
	LegalRepresentative string    `json:"legal_representative" form:"legal_representative" xml:"legal_representative" url:"legal_representative" validate:"required|string|maxLen:50" label:"法定代表人"`
	RegisteredCapital   float64   `json:"registered_capital" form:"registered_capital" xml:"registered_capital" url:"registered_capital" validate:"required|float" label:"注册资本（万元）"`
	EstablishmentDate   time.Time `json:"establishment_date" form:"establishment_date" xml:"establishment_date" url:"establishment_date" validate:"required" label:"成立日期"`
	BusinessScope       string    `json:"business_scope" form:"business_scope" xml:"business_scope" url:"business_scope" validate:"required|string" label:"经营范围"`
	CityCode            uint32    `json:"city_code" form:"city_code" xml:"city_code" url:"city_code" validate:"required|uint" label:"城市编码"`
	Address             string    `json:"address" form:"address" xml:"address" url:"address" validate:"required|string|maxLen:255" label:"详细地址"`
	ContactPhone        string    `json:"contact_phone" form:"contact_phone" xml:"contact_phone" url:"contact_phone" validate:"required|string|maxLen:20" label:"联系电话"`
	ContactEmail        string    `json:"contact_email" form:"contact_email" xml:"contact_email" url:"contact_email" validate:"required|string|maxLen:100|email" label:"联系邮箱"`
	Website             string    `json:"website" form:"website" xml:"website" url:"website" validate:"required|string|maxLen:150" label:"官网"`
	IndustryId          uint      `json:"industry_id" form:"industry_id" xml:"industry_id" url:"industry_id" validate:"required|uint" label:"所属行业ID（关联行业字典表）"`
	Status              uint8     `json:"status" form:"status" xml:"status" url:"status" validate:"uint" label:"状态: 1-正常, 0-停用, 2-注销"`
}

// ConfigValidation configures gookit/validate scenes.
func (SafeCompanyRequest) ConfigValidation(v *validate.Validation) {
	v.WithScenes(validate.SValues{
		"insert": []string{"CompanyName", "ShortName", "CreditCode", "LegalRepresentative", "RegisteredCapital", "EstablishmentDate", "BusinessScope", "CityCode", "Address", "ContactPhone", "ContactEmail", "Website", "IndustryId", "Status"},
		"update": []string{"Id", "CompanyName", "ShortName", "CreditCode", "LegalRepresentative", "RegisteredCapital", "EstablishmentDate", "BusinessScope", "CityCode", "Address", "ContactPhone", "ContactEmail", "Website", "IndustryId", "Status"},
		"delete": []string{"Id"},
	})
}

// Messages defines Chinese validation messages.
func (SafeCompanyRequest) Messages() map[string]string {
	return validate.MS{
		"email":    "{field}格式不正确",
		"float":    "{field}必须是浮点数",
		"maxLen":   "{field}长度不能大于 %v",
		"required": "{field}不能为空",
		"string":   "{field}必须是字符串",
		"uint":     "{field}必须是非负整数",
	}
}
