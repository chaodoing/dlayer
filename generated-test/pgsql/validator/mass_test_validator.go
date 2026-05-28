package validator

import (
	"github.com/gookit/validate"
)

// MassTestRequest is the request validator for the mass_test table.
type MassTestRequest struct {
	Id                  int     `json:"id" form:"id" xml:"id" url:"id" validate:"required|int" label:"主键ID"`
	PlaceId             int     `json:"place_id" form:"place_id" xml:"place_id" url:"place_id" validate:"required|int" label:"场景ID"`
	UserId              int     `json:"user_id" form:"user_id" xml:"user_id" url:"user_id" validate:"required|int" label:"用户ID"`
	Operator            string  `json:"operator" form:"operator" xml:"operator" url:"operator" validate:"required|string|maxLen:50" label:"运营商"`
	ZoneCode            int     `json:"zone_code" form:"zone_code" xml:"zone_code" url:"zone_code" validate:"required|int" label:"城市编码"`
	NetworkType         string  `json:"network_type" form:"network_type" xml:"network_type" url:"network_type" validate:"required|string|maxLen:50" label:"网络类型"`
	AvgUpload           float32 `json:"avg_upload" form:"avg_upload" xml:"avg_upload" url:"avg_upload" validate:"required|float" label:"平均上传速率"`
	AvgDownload         float32 `json:"avg_download" form:"avg_download" xml:"avg_download" url:"avg_download" validate:"required|float" label:"平均下载速率"`
	PeakUpload          float32 `json:"peak_upload" form:"peak_upload" xml:"peak_upload" url:"peak_upload" validate:"required|float" label:"峰值上传速率"`
	PeakDownload        float32 `json:"peak_download" form:"peak_download" xml:"peak_download" url:"peak_download" validate:"required|float" label:"峰值下载速率"`
	DataState           int     `json:"data_state" form:"data_state" xml:"data_state" url:"data_state" validate:"int" label:"数据状态 1:正常 0:失效"`
	SignalStrength      float32 `json:"signal_strength" form:"signal_strength" xml:"signal_strength" url:"signal_strength" validate:"required|float" label:"信号强度"`
	Sinr                float32 `json:"sinr" form:"sinr" xml:"sinr" url:"sinr" validate:"required|float" label:"信噪比"`
	LossRate            float32 `json:"loss_rate" form:"loss_rate" xml:"loss_rate" url:"loss_rate" validate:"required|float" label:"丢包率"`
	Delay               float32 `json:"delay" form:"delay" xml:"delay" url:"delay" validate:"float" label:"平均延迟"`
	Jitter              float32 `json:"jitter" form:"jitter" xml:"jitter" url:"jitter" validate:"required|float" label:"抖动"`
	Ping                float32 `json:"ping" form:"ping" xml:"ping" url:"ping" validate:"required|float" label:"ping值"`
	Terminal            string  `json:"terminal" form:"terminal" xml:"terminal" url:"terminal" validate:"required|string|maxLen:50" label:"终端型号"`
	Longitude           float32 `json:"longitude" form:"longitude" xml:"longitude" url:"longitude" validate:"required|float" label:"上传经度"`
	Latitude            float32 `json:"latitude" form:"latitude" xml:"latitude" url:"latitude" validate:"required|float" label:"上传纬度"`
	ImagePhotoId        *int    `json:"image_photo_id" form:"image_photo_id" xml:"image_photo_id" url:"image_photo_id" validate:"int" label:"图片文件相册ID"`
	VideoPhotoId        *int    `json:"video_photo_id" form:"video_photo_id" xml:"video_photo_id" url:"video_photo_id" validate:"int" label:"视频文件相册ID"`
	Location            string  `json:"location" form:"location" xml:"location" url:"location" validate:"required|string|maxLen:100" label:"位置名称"`
	LocationDescription string  `json:"location_description" form:"location_description" xml:"location_description" url:"location_description" validate:"required|string|maxLen:100" label:"位置描述"`
	LocationType        int     `json:"location_type" form:"location_type" xml:"location_type" url:"location_type" validate:"required|int" label:"位置类型 0:正常 1:地下 2:电梯"`
}

// ConfigValidation configures gookit/validate scenes.
func (MassTestRequest) ConfigValidation(v *validate.Validation) {
	v.WithScenes(validate.SValues{
		"insert": []string{"PlaceId", "UserId", "Operator", "ZoneCode", "NetworkType", "AvgUpload", "AvgDownload", "PeakUpload", "PeakDownload", "DataState", "SignalStrength", "Sinr", "LossRate", "Delay", "Jitter", "Ping", "Terminal", "Longitude", "Latitude", "ImagePhotoId", "VideoPhotoId", "Location", "LocationDescription", "LocationType"},
		"update": []string{"Id", "PlaceId", "UserId", "Operator", "ZoneCode", "NetworkType", "AvgUpload", "AvgDownload", "PeakUpload", "PeakDownload", "DataState", "SignalStrength", "Sinr", "LossRate", "Delay", "Jitter", "Ping", "Terminal", "Longitude", "Latitude", "ImagePhotoId", "VideoPhotoId", "Location", "LocationDescription", "LocationType"},
		"delete": []string{"Id"},
	})
}

// Messages defines Chinese validation messages.
func (MassTestRequest) Messages() map[string]string {
	return validate.MS{
		"float":    "{field}必须是浮点数",
		"int":      "{field}必须是整数",
		"maxLen":   "{field}长度不能大于 %v",
		"required": "{field}不能为空",
		"string":   "{field}必须是字符串",
	}
}
