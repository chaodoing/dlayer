package validator

import (
	"github.com/gookit/validate"
)

// SafeDisposalRequest is the request validator for the safe_disposal table.
type SafeDisposalRequest struct {
	Id      uint64 `json:"id" form:"id" xml:"id" url:"id" validate:"required|uint" label:"主键"`
	EventId string `json:"event_id" form:"event_id" xml:"event_id" url:"event_id" validate:"required|string|maxLen:36" label:"事件ID"`
	Note    string `json:"note" form:"note" xml:"note" url:"note" validate:"required|string|maxLen:255" label:"处置说明"`
	AdminId uint   `json:"admin_id" form:"admin_id" xml:"admin_id" url:"admin_id" validate:"required|uint" label:"处置人"`
	Status  uint8  `json:"status" form:"status" xml:"status" url:"status" validate:"uint" label:"处置审核状态"`
	Type    uint8  `json:"type" form:"type" xml:"type" url:"type" validate:"uint" label:"处置类型1:DDOS 2:信安"`
}

// ConfigValidation configures gookit/validate scenes.
func (SafeDisposalRequest) ConfigValidation(v *validate.Validation) {
	v.WithScenes(validate.SValues{
		"insert": []string{"EventId", "Note", "AdminId", "Status", "Type"},
		"update": []string{"Id", "EventId", "Note", "AdminId", "Status", "Type"},
		"delete": []string{"Id"},
	})
}

// Messages defines Chinese validation messages.
func (SafeDisposalRequest) Messages() map[string]string {
	return validate.MS{
		"maxLen":   "{field}长度不能大于 %v",
		"required": "{field}不能为空",
		"string":   "{field}必须是字符串",
		"uint":     "{field}必须是非负整数",
	}
}
