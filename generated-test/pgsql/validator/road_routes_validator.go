package validator

import (
	"github.com/gookit/validate"
)

// RoadRoutesRequest is the request validator for the road_routes table.
type RoadRoutesRequest struct {
	RoadSectionId int     `json:"road_section_id" form:"road_section_id" xml:"road_section_id" url:"road_section_id" validate:"required|int" label:"路段信息ID"`
	RoadId        int     `json:"road_id" form:"road_id" xml:"road_id" url:"road_id" validate:"required|int" label:"道路ID"`
	Longitude     float32 `json:"longitude" form:"longitude" xml:"longitude" url:"longitude" validate:"required|float" label:"经度"`
	Latitude      float32 `json:"latitude" form:"latitude" xml:"latitude" url:"latitude" validate:"required|float" label:"纬度"`
}

// ConfigValidation configures gookit/validate scenes.
func (RoadRoutesRequest) ConfigValidation(v *validate.Validation) {
	v.WithScenes(validate.SValues{
		"insert": []string{"RoadSectionId", "RoadId", "Longitude", "Latitude"},
		"update": []string{"RoadSectionId", "RoadId", "Longitude", "Latitude"},
		"delete": []string{},
	})
}

// Messages defines Chinese validation messages.
func (RoadRoutesRequest) Messages() map[string]string {
	return validate.MS{
		"float":    "{field}必须是浮点数",
		"int":      "{field}必须是整数",
		"required": "{field}不能为空",
	}
}
