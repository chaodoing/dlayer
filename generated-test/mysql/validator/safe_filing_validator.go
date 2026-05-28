package validator

import (
	"github.com/gookit/validate"
	"time"
)

// SafeFilingRequest is the request validator for the safe_filing table.
type SafeFilingRequest struct {
	Id              uint64    `json:"id" form:"id" xml:"id" url:"id" validate:"required|uint" label:"备案记录ID"`
	CompanyId       uint64    `json:"company_id" form:"company_id" xml:"company_id" url:"company_id" validate:"required|uint" label:"关联公司ID（外键→company.id）"`
	FilingType      int8      `json:"filing_type" form:"filing_type" xml:"filing_type" url:"filing_type" validate:"required|int" label:"备案类型: 1=ICP, 2=公安, 3=工商变更, 4=行业专项"`
	FilingNumber    string    `json:"filing_number" form:"filing_number" xml:"filing_number" url:"filing_number" validate:"required|string|maxLen:100" label:"备案号（例：京ICP备12345678号）"`
	FilingSubject   string    `json:"filing_subject" form:"filing_subject" xml:"filing_subject" url:"filing_subject" validate:"required|string|maxLen:200" label:"备案主体（与工商名可能不同）"`
	Status          int8      `json:"status" form:"status" xml:"status" url:"status" validate:"int" label:"状态: 0=草稿,1=审核中,2=已通过,3=驳回,4=已注销,5=过期"`
	ApplyDate       time.Time `json:"apply_date" form:"apply_date" xml:"apply_date" url:"apply_date" validate:"required" label:"申请日期"`
	ApproveDate     time.Time `json:"approve_date" form:"approve_date" xml:"approve_date" url:"approve_date" validate:"required" label:"批准日期"`
	ValidUntil      time.Time `json:"valid_until" form:"valid_until" xml:"valid_until" url:"valid_until" validate:"required" label:"有效期至（NULL=长期有效）"`
	IsCurrent       uint8     `json:"is_current" form:"is_current" xml:"is_current" url:"is_current" validate:"uint" label:"是否为当前有效备案（用于同一公司多备案场景）"`
	CertificatePath string    `json:"certificate_path" form:"certificate_path" xml:"certificate_path" url:"certificate_path" validate:"required|string|maxLen:500" label:"证书文件OSS路径"`
	Remark          string    `json:"remark" form:"remark" xml:"remark" url:"remark" validate:"required|string|maxLen:500" label:"备注（驳回原因/特殊说明）"`
	AuditUserId     uint64    `json:"audit_user_id" form:"audit_user_id" xml:"audit_user_id" url:"audit_user_id" validate:"required|uint" label:"审核人ID"`
	AuditAt         time.Time `json:"audit_at" form:"audit_at" xml:"audit_at" url:"audit_at" validate:"required" label:"审核时间"`
	CreatedBy       uint64    `json:"created_by" form:"created_by" xml:"created_by" url:"created_by" validate:"required|uint" label:"创建人"`
}

// ConfigValidation configures gookit/validate scenes.
func (SafeFilingRequest) ConfigValidation(v *validate.Validation) {
	v.WithScenes(validate.SValues{
		"insert": []string{"CompanyId", "FilingType", "FilingNumber", "FilingSubject", "Status", "ApplyDate", "ApproveDate", "ValidUntil", "IsCurrent", "CertificatePath", "Remark", "AuditUserId", "AuditAt", "CreatedBy"},
		"update": []string{"Id", "CompanyId", "FilingType", "FilingNumber", "FilingSubject", "Status", "ApplyDate", "ApproveDate", "ValidUntil", "IsCurrent", "CertificatePath", "Remark", "AuditUserId", "AuditAt", "CreatedBy"},
		"delete": []string{"Id"},
	})
}

// Messages defines Chinese validation messages.
func (SafeFilingRequest) Messages() map[string]string {
	return validate.MS{
		"int":      "{field}必须是整数",
		"maxLen":   "{field}长度不能大于 %v",
		"required": "{field}不能为空",
		"string":   "{field}必须是字符串",
		"uint":     "{field}必须是非负整数",
	}
}
