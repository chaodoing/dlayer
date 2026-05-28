package validator

import (
	"github.com/gookit/validate"
)

// RoadCommitIndicatorsRequest is the request validator for the road_commit_indicators table.
type RoadCommitIndicatorsRequest struct {
	Id         int     `json:"id" form:"id" xml:"id" url:"id" validate:"required|int" label:"主键ID"`
	Uid        int     `json:"uid" form:"uid" xml:"uid" url:"uid" validate:"required|int" label:"前台用户ID"`
	CommitId   int     `json:"commit_id" form:"commit_id" xml:"commit_id" url:"commit_id" validate:"required|int" label:"提交ID"`
	RoadId     int     `json:"road_id" form:"road_id" xml:"road_id" url:"road_id" validate:"required|int" label:"道路ID"`
	TaskId     int     `json:"task_id" form:"task_id" xml:"task_id" url:"task_id" validate:"required|int" label:"路测任务ID"`
	Longitude  float32 `json:"longitude" form:"longitude" xml:"longitude" url:"longitude" validate:"required|float" label:"经度"`
	Latitude   float32 `json:"latitude" form:"latitude" xml:"latitude" url:"latitude" validate:"required|float" label:"纬度"`
	ZoneCode   int     `json:"zone_code" form:"zone_code" xml:"zone_code" url:"zone_code" validate:"required|int" label:"地区代码"`
	Location   string  `json:"location" form:"location" xml:"location" url:"location" validate:"required|string|maxLen:120" label:"提交位置"`
	Sinr       float32 `json:"sinr" form:"sinr" xml:"sinr" url:"sinr" validate:"required|float" label:"平均信噪比"`
	Signal     float32 `json:"signal" form:"signal" xml:"signal" url:"signal" validate:"required|float" label:"平均信号值"`
	Delay      float32 `json:"delay" form:"delay" xml:"delay" url:"delay" validate:"required|float" label:"平均延迟"`
	Ping       float32 `json:"ping" form:"ping" xml:"ping" url:"ping" validate:"required|float" label:"ping延时(ms)"`
	Jitter     float32 `json:"jitter" form:"jitter" xml:"jitter" url:"jitter" validate:"required|float" label:"平均抖动(ms)"`
	PacketLoss float32 `json:"packet_loss" form:"packet_loss" xml:"packet_loss" url:"packet_loss" validate:"required|float" label:"平均丢包率"`
}

// ConfigValidation configures gookit/validate scenes.
func (RoadCommitIndicatorsRequest) ConfigValidation(v *validate.Validation) {
	v.WithScenes(validate.SValues{
		"insert": []string{"Uid", "CommitId", "RoadId", "TaskId", "Longitude", "Latitude", "ZoneCode", "Location", "Sinr", "Signal", "Delay", "Ping", "Jitter", "PacketLoss"},
		"update": []string{"Id", "Uid", "CommitId", "RoadId", "TaskId", "Longitude", "Latitude", "ZoneCode", "Location", "Sinr", "Signal", "Delay", "Ping", "Jitter", "PacketLoss"},
		"delete": []string{"Id"},
	})
}

// Messages defines Chinese validation messages.
func (RoadCommitIndicatorsRequest) Messages() map[string]string {
	return validate.MS{
		"float":    "{field}必须是浮点数",
		"int":      "{field}必须是整数",
		"maxLen":   "{field}长度不能大于 %v",
		"required": "{field}不能为空",
		"string":   "{field}必须是字符串",
	}
}
