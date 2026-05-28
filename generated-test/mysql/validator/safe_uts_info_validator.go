package validator

import (
	"github.com/gookit/validate"
	"gorm.io/datatypes"
)

// SafeUtsInfoRequest is the request validator for the safe_uts_info table.
type SafeUtsInfoRequest struct {
	UtsId uint           `json:"uts_id" form:"uts_id" xml:"uts_id" url:"uts_id" validate:"required|uint" label:"探针id"`
	Cpu   datatypes.JSON `json:"cpu" form:"cpu" xml:"cpu" url:"cpu" validate:"required" label:"cpu信息"`
	Ram   datatypes.JSON `json:"ram" form:"ram" xml:"ram" url:"ram" validate:"required" label:"内存信息"`
	Disk  datatypes.JSON `json:"disk" form:"disk" xml:"disk" url:"disk" validate:"required" label:"磁盘信息"`
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
		"required": "{field}不能为空",
		"uint":     "{field}必须是非负整数",
	}
}
