package validator

import (
	"github.com/gookit/validate"
	"gorm.io/datatypes"
)

// RoadSectionsRequest is the request validator for the road_sections table.
type RoadSectionsRequest struct {
	Id          int            `json:"id" form:"id" xml:"id" url:"id" validate:"required|int" label:"道路路段"`
	RoadId      int            `json:"road_id" form:"road_id" xml:"road_id" url:"road_id" validate:"required|int" label:"所属道路ID"`
	Name        string         `json:"name" form:"name" xml:"name" url:"name" validate:"required|string|maxLen:50" label:"路段名称"`
	Sequence    int            `json:"sequence" form:"sequence" xml:"sequence" url:"sequence" validate:"required|int" label:"路段序号"`
	Routes      datatypes.JSON `json:"routes" form:"routes" xml:"routes" url:"routes" validate:"required" label:"途径点"`
	Length      float32        `json:"length" form:"length" xml:"length" url:"length" validate:"required|float" label:"路段长度(km)"`
	Description string         `json:"description" form:"description" xml:"description" url:"description" validate:"required|string|maxLen:100" label:"路段描述"`
}

// ConfigValidation configures gookit/validate scenes.
func (RoadSectionsRequest) ConfigValidation(v *validate.Validation) {
	v.WithScenes(validate.SValues{
		"insert": []string{"RoadId", "Name", "Sequence", "Routes", "Length", "Description"},
		"update": []string{"Id", "RoadId", "Name", "Sequence", "Routes", "Length", "Description"},
		"delete": []string{"Id"},
	})
}

// Messages defines Chinese validation messages.
func (RoadSectionsRequest) Messages() map[string]string {
	return validate.MS{
		"float":    "{field}必须是浮点数",
		"int":      "{field}必须是整数",
		"maxLen":   "{field}长度不能大于 %v",
		"required": "{field}不能为空",
		"string":   "{field}必须是字符串",
	}
}
