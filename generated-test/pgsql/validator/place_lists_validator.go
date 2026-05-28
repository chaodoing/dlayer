package validator

import (
	"github.com/gookit/validate"
)

// PlaceListsRequest is the request validator for the place_lists table.
type PlaceListsRequest struct {
	Id       int    `json:"id" form:"id" xml:"id" url:"id" validate:"required|int" label:"主键"`
	Name     string `json:"name" form:"name" xml:"name" url:"name" validate:"required|string|maxLen:50" label:"清单名称"`
	Disabled bool   `json:"disabled" form:"disabled" xml:"disabled" url:"disabled" validate:"bool" label:"禁用状态"`
}

// ConfigValidation configures gookit/validate scenes.
func (PlaceListsRequest) ConfigValidation(v *validate.Validation) {
	v.WithScenes(validate.SValues{
		"insert": []string{"Name", "Disabled"},
		"update": []string{"Id", "Name", "Disabled"},
		"delete": []string{"Id"},
	})
}

// Messages defines Chinese validation messages.
func (PlaceListsRequest) Messages() map[string]string {
	return validate.MS{
		"bool":     "{field}必须是布尔值",
		"int":      "{field}必须是整数",
		"maxLen":   "{field}长度不能大于 %v",
		"required": "{field}不能为空",
		"string":   "{field}必须是字符串",
	}
}
