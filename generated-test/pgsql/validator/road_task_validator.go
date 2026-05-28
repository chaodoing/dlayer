package validator

import (
	"github.com/gookit/validate"
)

// RoadTaskRequest is the request validator for the road_task table.
type RoadTaskRequest struct {
	Id             int    `json:"id" form:"id" xml:"id" url:"id" validate:"required|int" label:"路测任务ID"`
	Uid            int    `json:"uid" form:"uid" xml:"uid" url:"uid" validate:"required|int" label:"创建者用户ID"`
	DepartmentId   int    `json:"department_id" form:"department_id" xml:"department_id" url:"department_id" validate:"required|int" label:"分配的测试部门"`
	RoadId         int    `json:"road_id" form:"road_id" xml:"road_id" url:"road_id" validate:"required|int" label:"道路ID"`
	Name           string `json:"name" form:"name" xml:"name" url:"name" validate:"required|string|maxLen:50" label:"任务名称"`
	Operator       string `json:"operator" form:"operator" xml:"operator" url:"operator" validate:"required|string|maxLen:50" label:"要求测试运营商[电信,移动,广电,联通]"`
	NetworkType    string `json:"network_type" form:"network_type" xml:"network_type" url:"network_type" validate:"required|string|maxLen:50" label:"要求测试网络类型[WIFI,4G,5G]"`
	Status         int    `json:"status" form:"status" xml:"status" url:"status" validate:"int" label:"任务状态: 0已取消 1 进行中 2已完成"`
	SignalDistance int    `json:"signal_distance" form:"signal_distance" xml:"signal_distance" url:"signal_distance" validate:"required|int" label:"指标上传距离:单位M"`
	RateDistance   int    `json:"rate_distance" form:"rate_distance" xml:"rate_distance" url:"rate_distance" validate:"required|int" label:"上下行上传速率:单位M"`
	Radius         int    `json:"radius" form:"radius" xml:"radius" url:"radius" validate:"required|int" label:"误差半径"`
	HasImage       bool   `json:"has_image" form:"has_image" xml:"has_image" url:"has_image" validate:"bool" label:"是否上传图片"`
	HasVideo       bool   `json:"has_video" form:"has_video" xml:"has_video" url:"has_video" validate:"bool" label:"是否上传视频"`
}

// ConfigValidation configures gookit/validate scenes.
func (RoadTaskRequest) ConfigValidation(v *validate.Validation) {
	v.WithScenes(validate.SValues{
		"insert": []string{"Uid", "DepartmentId", "RoadId", "Name", "Operator", "NetworkType", "Status", "SignalDistance", "RateDistance", "Radius", "HasImage", "HasVideo"},
		"update": []string{"Id", "Uid", "DepartmentId", "RoadId", "Name", "Operator", "NetworkType", "Status", "SignalDistance", "RateDistance", "Radius", "HasImage", "HasVideo"},
		"delete": []string{"Id"},
	})
}

// Messages defines Chinese validation messages.
func (RoadTaskRequest) Messages() map[string]string {
	return validate.MS{
		"bool":     "{field}必须是布尔值",
		"int":      "{field}必须是整数",
		"maxLen":   "{field}长度不能大于 %v",
		"required": "{field}不能为空",
		"string":   "{field}必须是字符串",
	}
}
