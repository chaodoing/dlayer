package validator

import (
	"github.com/gookit/validate"
)

// PlaceScenesRequest is the request validator for the place_scenes table.
type PlaceScenesRequest struct {
	Id       int    `json:"id" form:"id" xml:"id" url:"id" validate:"required|int" label:"场景ID"`
	Name     string `json:"name" form:"name" xml:"name" url:"name" validate:"required|string|maxLen:50" label:"场景名称"`
	ParentId *int   `json:"parent_id" form:"parent_id" xml:"parent_id" url:"parent_id" validate:"int" label:"上级场景ID"`
	Level    int    `json:"level" form:"level" xml:"level" url:"level" validate:"required|int" label:"场景级别"`
	Sort     int    `json:"sort" form:"sort" xml:"sort" url:"sort" validate:"required|int" label:"场景排序"`
	Disabled bool   `json:"disabled" form:"disabled" xml:"disabled" url:"disabled" validate:"bool" label:"禁用状态"`
}

// ConfigValidation configures gookit/validate scenes.
func (PlaceScenesRequest) ConfigValidation(v *validate.Validation) {
	v.WithScenes(validate.SValues{
		"insert": []string{"Name", "ParentId", "Level", "Sort", "Disabled"},
		"update": []string{"Id", "Name", "ParentId", "Level", "Sort", "Disabled"},
		"delete": []string{"Id"},
	})
}

// Messages defines Chinese validation messages.
func (PlaceScenesRequest) Messages() map[string]string {
	return validate.MS{
		"bool":     "{field}必须是布尔值",
		"int":      "{field}必须是整数",
		"maxLen":   "{field}长度不能大于 %v",
		"required": "{field}不能为空",
		"string":   "{field}必须是字符串",
	}
}
