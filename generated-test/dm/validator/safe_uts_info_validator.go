package validator

import (
	"github.com/gookit/validate"
)

// SafeUtsInfoRequest is the request validator for the safe_uts_info table.
type SafeUtsInfoRequest struct {
	UtsId int    `json:"uts_id" form:"uts_id" xml:"uts_id" url:"uts_id" validate:"required|int" label:"探针id"`
	Cpu   string `json:"cpu" form:"cpu" xml:"cpu" url:"cpu" validate:"required|string" label:"cpu信息"`
	Ram   string `json:"ram" form:"ram" xml:"ram" url:"ram" validate:"required|string" label:"内存信息"`
	Disk  string `json:"disk" form:"disk" xml:"disk" url:"disk" validate:"required|string" label:"磁盘信息"`
}

// ConfigValidation configures gookit/validate scenes.
func (SafeUtsInfoRequest) ConfigValidation(v *validate.Validation) {
	v.WithScenes(validate.SValues{
		"insert": []string{"UtsId", "Cpu", "Ram", "Disk"},
		"update": []string{"UtsId", "Cpu", "Ram", "Disk"},
		"delete": []string{},
	})
}

// Messages defines Chinese validation messages.
func (SafeUtsInfoRequest) Messages() map[string]string {
	return validate.MS{
		"int":      "{field}必须是整数",
		"required": "{field}不能为空",
		"string":   "{field}必须是字符串",
	}
}
