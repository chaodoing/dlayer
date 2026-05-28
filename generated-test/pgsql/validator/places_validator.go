package validator

import (
	"github.com/gookit/validate"
)

// PlacesRequest is the request validator for the places table.
type PlacesRequest struct {
	Id          int     `json:"id" form:"id" xml:"id" url:"id" validate:"required|int" label:"主键"`
	ZoneCode    int     `json:"zone_code" form:"zone_code" xml:"zone_code" url:"zone_code" validate:"required|int" label:"地区编码"`
	SceneId     int     `json:"scene_id" form:"scene_id" xml:"scene_id" url:"scene_id" validate:"required|int" label:"地点场景ID"`
	Name        string  `json:"name" form:"name" xml:"name" url:"name" validate:"required|string|maxLen:50" label:"场景名称"`
	Address     string  `json:"address" form:"address" xml:"address" url:"address" validate:"required|string|maxLen:80" label:"场景具体地址"`
	Longitude   float32 `json:"longitude" form:"longitude" xml:"longitude" url:"longitude" validate:"required|float" label:"经度"`
	Latitude    float32 `json:"latitude" form:"latitude" xml:"latitude" url:"latitude" validate:"required|float" label:"维度"`
	Description string  `json:"description" form:"description" xml:"description" url:"description" validate:"required|string|maxLen:60" label:"描述信息"`
	Disabled    bool    `json:"disabled" form:"disabled" xml:"disabled" url:"disabled" validate:"bool" label:"禁用状态"`
	PingCount   int     `json:"ping_count" form:"ping_count" xml:"ping_count" url:"ping_count" validate:"required|int" label:"测试次数"`
}

// ConfigValidation configures gookit/validate scenes.
func (PlacesRequest) ConfigValidation(v *validate.Validation) {
	v.WithScenes(validate.SValues{
		"insert": []string{"ZoneCode", "SceneId", "Name", "Address", "Longitude", "Latitude", "Description", "Disabled", "PingCount"},
		"update": []string{"Id", "ZoneCode", "SceneId", "Name", "Address", "Longitude", "Latitude", "Description", "Disabled", "PingCount"},
		"delete": []string{"Id"},
	})
}

// Messages defines Chinese validation messages.
func (PlacesRequest) Messages() map[string]string {
	return validate.MS{
		"bool":     "{field}必须是布尔值",
		"float":    "{field}必须是浮点数",
		"int":      "{field}必须是整数",
		"maxLen":   "{field}长度不能大于 %v",
		"required": "{field}不能为空",
		"string":   "{field}必须是字符串",
	}
}
