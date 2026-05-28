package validator

import (
	"github.com/gookit/validate"
	"gorm.io/datatypes"
)

// MassSpeedsRequest is the request validator for the mass_speeds table.
type MassSpeedsRequest struct {
	TestId   int            `json:"test_id" form:"test_id" xml:"test_id" url:"test_id" validate:"required|int" label:"众测ID"`
	Upload   datatypes.JSON `json:"upload" form:"upload" xml:"upload" url:"upload" validate:"required" label:"上传速率"`
	Download datatypes.JSON `json:"download" form:"download" xml:"download" url:"download" validate:"required" label:"下载速率"`
}

// ConfigValidation configures gookit/validate scenes.
func (MassSpeedsRequest) ConfigValidation(v *validate.Validation) {
	v.WithScenes(validate.SValues{
		"insert": []string{"TestId", "Upload", "Download"},
		"update": []string{"TestId", "Upload", "Download"},
		"delete": []string{},
	})
}

// Messages defines Chinese validation messages.
func (MassSpeedsRequest) Messages() map[string]string {
	return validate.MS{
		"int":      "{field}必须是整数",
		"required": "{field}不能为空",
	}
}
