package validator

import (
	"github.com/gookit/validate"
	"gorm.io/datatypes"
	"time"
)

// RoadCommitRequest is the request validator for the road_commit table.
type RoadCommitRequest struct {
	Id             int            `json:"id" form:"id" xml:"id" url:"id" validate:"required|int" label:"主键ID"`
	Uid            int            `json:"uid" form:"uid" xml:"uid" url:"uid" validate:"required|int" label:"后台用户ID"`
	RoadId         int            `json:"road_id" form:"road_id" xml:"road_id" url:"road_id" validate:"required|int" label:"道路id"`
	TaskId         int            `json:"task_id" form:"task_id" xml:"task_id" url:"task_id" validate:"required|int" label:"任务ID"`
	StartTime      time.Time      `json:"start_time" form:"start_time" xml:"start_time" url:"start_time" validate:"required" label:"开始时间"`
	DoneTime       time.Time      `json:"done_time" form:"done_time" xml:"done_time" url:"done_time" validate:"required" label:"完成时间"`
	StartLongitude float32        `json:"start_longitude" form:"start_longitude" xml:"start_longitude" url:"start_longitude" validate:"required|float" label:"开始经度"`
	StartLatitude  float32        `json:"start_latitude" form:"start_latitude" xml:"start_latitude" url:"start_latitude" validate:"required|float" label:"开始维度"`
	DoneLongitude  float32        `json:"done_longitude" form:"done_longitude" xml:"done_longitude" url:"done_longitude" validate:"required|float" label:"完成经度"`
	DoneLatitude   float32        `json:"done_latitude" form:"done_latitude" xml:"done_latitude" url:"done_latitude" validate:"required|float" label:"完成维度"`
	SignalCount    int            `json:"signal_count" form:"signal_count" xml:"signal_count" url:"signal_count" validate:"required|int" label:"信号指标上传次数"`
	SpeedCount     int            `json:"speed_count" form:"speed_count" xml:"speed_count" url:"speed_count" validate:"required|int" label:"测试数据上传次数"`
	Terminal       string         `json:"terminal" form:"terminal" xml:"terminal" url:"terminal" validate:"required|string|maxLen:100" label:"终端设备信息"`
	Operator       string         `json:"operator" form:"operator" xml:"operator" url:"operator" validate:"required|string|maxLen:50" label:"运营商"`
	NetworkType    string         `json:"network_type" form:"network_type" xml:"network_type" url:"network_type" validate:"required|string|maxLen:50" label:"网络类型"`
	AvgUpload      float32        `json:"avg_upload" form:"avg_upload" xml:"avg_upload" url:"avg_upload" validate:"required|float" label:"平均上传速度(Mbps)"`
	AvgDownload    float32        `json:"avg_download" form:"avg_download" xml:"avg_download" url:"avg_download" validate:"required|float" label:"平均下载速度(Mbps)"`
	PeakUpload     float32        `json:"peak_upload" form:"peak_upload" xml:"peak_upload" url:"peak_upload" validate:"required|float" label:"峰值上传速度(Mbps)"`
	PeakDownload   float32        `json:"peak_download" form:"peak_download" xml:"peak_download" url:"peak_download" validate:"required|float" label:"峰值下载速度(Mbps)"`
	Ping           float32        `json:"ping" form:"ping" xml:"ping" url:"ping" validate:"required|float" label:"ping延时(ms)"`
	Jitter         float32        `json:"jitter" form:"jitter" xml:"jitter" url:"jitter" validate:"required|float" label:"抖动(ms)"`
	PacketLoss     float32        `json:"packet_loss" form:"packet_loss" xml:"packet_loss" url:"packet_loss" validate:"required|float" label:"丢包率(%)"`
	Delay          float32        `json:"delay" form:"delay" xml:"delay" url:"delay" validate:"required|float" label:"平均延时(ms)"`
	Signal         float32        `json:"signal" form:"signal" xml:"signal" url:"signal" validate:"required|float" label:"信号强度"`
	Sinr           float32        `json:"sinr" form:"sinr" xml:"sinr" url:"sinr" validate:"required|float" label:"信噪比"`
	ImagePhotoIds  datatypes.JSON `json:"image_photo_ids" form:"image_photo_ids" xml:"image_photo_ids" url:"image_photo_ids" validate:"required" label:"图片[开始ID,结束ID]"`
	VideoPhotoIds  datatypes.JSON `json:"video_photo_ids" form:"video_photo_ids" xml:"video_photo_ids" url:"video_photo_ids" validate:"required" label:"视频[开始ID,结束ID]"`
	Status         int            `json:"status" form:"status" xml:"status" url:"status" validate:"int" label:"提交数据状态 -1:已失效 0:进行中 1:已完成"`
}

// ConfigValidation configures gookit/validate scenes.
func (RoadCommitRequest) ConfigValidation(v *validate.Validation) {
	v.WithScenes(validate.SValues{
		"insert": []string{"Uid", "RoadId", "TaskId", "StartTime", "DoneTime", "StartLongitude", "StartLatitude", "DoneLongitude", "DoneLatitude", "SignalCount", "SpeedCount", "Terminal", "Operator", "NetworkType", "AvgUpload", "AvgDownload", "PeakUpload", "PeakDownload", "Ping", "Jitter", "PacketLoss", "Delay", "Signal", "Sinr", "ImagePhotoIds", "VideoPhotoIds", "Status"},
		"update": []string{"Id", "Uid", "RoadId", "TaskId", "StartTime", "DoneTime", "StartLongitude", "StartLatitude", "DoneLongitude", "DoneLatitude", "SignalCount", "SpeedCount", "Terminal", "Operator", "NetworkType", "AvgUpload", "AvgDownload", "PeakUpload", "PeakDownload", "Ping", "Jitter", "PacketLoss", "Delay", "Signal", "Sinr", "ImagePhotoIds", "VideoPhotoIds", "Status"},
		"delete": []string{"Id"},
	})
}

// Messages defines Chinese validation messages.
func (RoadCommitRequest) Messages() map[string]string {
	return validate.MS{
		"float":    "{field}必须是浮点数",
		"int":      "{field}必须是整数",
		"maxLen":   "{field}长度不能大于 %v",
		"required": "{field}不能为空",
		"string":   "{field}必须是字符串",
	}
}
