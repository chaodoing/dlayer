package validator

import (
	"github.com/gookit/validate"
)

// RoadRequest is the request validator for the road table.
type RoadRequest struct {
	Id          int     `json:"id" form:"id" xml:"id" url:"id" validate:"required|int" label:"道路ID"`
	Name        string  `json:"name" form:"name" xml:"name" url:"name" validate:"required|string|maxLen:60" label:"道路名称"`
	Code        string  `json:"code" form:"code" xml:"code" url:"code" validate:"required|string|maxLen:30" label:"道路编码"`
	Type        int     `json:"type" form:"type" xml:"type" url:"type" validate:"required|int" label:"道路类型 1国道 2 省道 3 高铁"`
	Distance    float32 `json:"distance" form:"distance" xml:"distance" url:"distance" validate:"required|float" label:"道路总长度(公里)"`
	Uid         int     `json:"uid" form:"uid" xml:"uid" url:"uid" validate:"required|int" label:"创建后台用户ID"`
	Status      int     `json:"status" form:"status" xml:"status" url:"status" validate:"required|int" label:"状态(1:正常,2:维护,3:禁用)"`
	Description string  `json:"description" form:"description" xml:"description" url:"description" validate:"required|string|maxLen:120" label:"描述信息"`
}

// ConfigValidation configures gookit/validate scenes.
func (RoadRequest) ConfigValidation(v *validate.Validation) {
	v.WithScenes(validate.SValues{
		"insert": []string{"Name", "Code", "Type", "Distance", "Uid", "Status", "Description"},
		"update": []string{"Id", "Name", "Code", "Type", "Distance", "Uid", "Status", "Description"},
		"delete": []string{"Id"},
	})
}

// Messages defines Chinese validation messages.
func (RoadRequest) Messages() map[string]string {
	return validate.MS{
		"float":    "{field}必须是浮点数",
		"int":      "{field}必须是整数",
		"maxLen":   "{field}长度不能大于 %v",
		"required": "{field}不能为空",
		"string":   "{field}必须是字符串",
	}
}
