package validator

import (
	"github.com/gookit/validate"
)

// PlaceContentsRequest is the request validator for the place_contents table.
type PlaceContentsRequest struct {
	PlaceListId int `json:"place_list_id" form:"place_list_id" xml:"place_list_id" url:"place_list_id" validate:"required|int" label:"清单ID"`
	PlaceId     int `json:"place_id" form:"place_id" xml:"place_id" url:"place_id" validate:"required|int" label:"场所ID"`
}

// ConfigValidation configures gookit/validate scenes.
func (PlaceContentsRequest) ConfigValidation(v *validate.Validation) {
	v.WithScenes(validate.SValues{
		"insert": []string{"PlaceListId", "PlaceId"},
		"update": []string{"PlaceListId", "PlaceId"},
		"delete": []string{},
	})
}

// Messages defines Chinese validation messages.
func (PlaceContentsRequest) Messages() map[string]string {
	return validate.MS{
		"int":      "{field}必须是整数",
		"required": "{field}不能为空",
	}
}
