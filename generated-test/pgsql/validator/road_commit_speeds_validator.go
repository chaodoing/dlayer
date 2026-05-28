package validator

import (
	"github.com/gookit/validate"
)

// RoadCommitSpeedsRequest is the request validator for the road_commit_speeds table.
type RoadCommitSpeedsRequest struct {
	Id           int     `json:"id" form:"id" xml:"id" url:"id" validate:"required|int" label:"主键ID"`
	Uid          int     `json:"uid" form:"uid" xml:"uid" url:"uid" validate:"required|int" label:"前台用户ID"`
	CommitId     int     `json:"commit_id" form:"commit_id" xml:"commit_id" url:"commit_id" validate:"required|int" label:"提交ID"`
	RoadId       int     `json:"road_id" form:"road_id" xml:"road_id" url:"road_id" validate:"required|int" label:"道路ID"`
	IndicatorId  int     `json:"indicator_id" form:"indicator_id" xml:"indicator_id" url:"indicator_id" validate:"required|int" label:"关联的指标数据ID"`
	TaskId       int     `json:"task_id" form:"task_id" xml:"task_id" url:"task_id" validate:"required|int" label:"路测任务ID"`
	Longitude    float32 `json:"longitude" form:"longitude" xml:"longitude" url:"longitude" validate:"required|float" label:"经度"`
	Latitude     float32 `json:"latitude" form:"latitude" xml:"latitude" url:"latitude" validate:"required|float" label:"纬度"`
	ZoneCode     int     `json:"zone_code" form:"zone_code" xml:"zone_code" url:"zone_code" validate:"required|int" label:"地区代码"`
	Location     string  `json:"location" form:"location" xml:"location" url:"location" validate:"required|string|maxLen:120" label:"提交位置"`
	AvgUpload    float32 `json:"avg_upload" form:"avg_upload" xml:"avg_upload" url:"avg_upload" validate:"required|float" label:"平均上传速率(Mbps)"`
	AvgDownload  float32 `json:"avg_download" form:"avg_download" xml:"avg_download" url:"avg_download" validate:"required|float" label:"平均下载速率(Mbps)"`
	PeakUpload   float32 `json:"peak_upload" form:"peak_upload" xml:"peak_upload" url:"peak_upload" validate:"required|float" label:"峰值上传速率(Mbps)"`
	PeakDownload float32 `json:"peak_download" form:"peak_download" xml:"peak_download" url:"peak_download" validate:"required|float" label:"峰值下载速率(Mbps)"`
}

// ConfigValidation configures gookit/validate scenes.
func (RoadCommitSpeedsRequest) ConfigValidation(v *validate.Validation) {
	v.WithScenes(validate.SValues{
		"insert": []string{"Uid", "CommitId", "RoadId", "IndicatorId", "TaskId", "Longitude", "Latitude", "ZoneCode", "Location", "AvgUpload", "AvgDownload", "PeakUpload", "PeakDownload"},
		"update": []string{"Id", "Uid", "CommitId", "RoadId", "IndicatorId", "TaskId", "Longitude", "Latitude", "ZoneCode", "Location", "AvgUpload", "AvgDownload", "PeakUpload", "PeakDownload"},
		"delete": []string{"Id"},
	})
}

// Messages defines Chinese validation messages.
func (RoadCommitSpeedsRequest) Messages() map[string]string {
	return validate.MS{
		"float":    "{field}必须是浮点数",
		"int":      "{field}必须是整数",
		"maxLen":   "{field}长度不能大于 %v",
		"required": "{field}不能为空",
		"string":   "{field}必须是字符串",
	}
}
