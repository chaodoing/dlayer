package validator

import (
	"github.com/gookit/validate"
)

// SafeDisposalRequest is the request validator for the safe_disposal table.
type SafeDisposalRequest struct {
	Id      int64  `json:"id" form:"id" xml:"id" url:"id" validate:"required|int" label:"主键"`
	EventId string `json:"event_id" form:"event_id" xml:"event_id" url:"event_id" validate:"required|string|maxLen:144" label:"事件ID"`
	Note    string `json:"note" form:"note" xml:"note" url:"note" validate:"required|string|maxLen:1020" label:"处置说明"`
	AdminId int64  `json:"admin_id" form:"admin_id" xml:"admin_id" url:"admin_id" validate:"required|int" label:"处置人"`
	Status  int    `json:"status" form:"status" xml:"status" url:"status" validate:"int" label:"处置审核状态"`
	Type    int    `json:"type" form:"type" xml:"type" url:"type" validate:"int" label:"处置类型1:DDOS 2:信安"`
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
		"int":      "{field}必须是整数",
		"maxLen":   "{field}长度不能大于 %v",
		"required": "{field}不能为空",
		"string":   "{field}必须是字符串",
	}
}
