package validator

import (
	"github.com/gookit/validate"
)

// RoadCommitServerRequest is the request validator for the road_commit_server table.
type RoadCommitServerRequest struct {
	RoadCommitId int    `json:"road_commit_id" form:"road_commit_id" xml:"road_commit_id" url:"road_commit_id" validate:"required|int" label:"路测提交ID"`
	DlLoaded     int    `json:"dl_loaded" form:"dl_loaded" xml:"dl_loaded" url:"dl_loaded" validate:"required|int" label:"下行已接收数据量"`
	UlLoaded     int    `json:"ul_loaded" form:"ul_loaded" xml:"ul_loaded" url:"ul_loaded" validate:"required|int" label:"上行已发送数据量"`
	IpAddress    string `json:"ip_address" form:"ip_address" xml:"ip_address" url:"ip_address" validate:"required|string|maxLen:100" label:"IP地址"`
	VersionCode  int    `json:"version_code" form:"version_code" xml:"version_code" url:"version_code" validate:"required|int" label:"App内部版本号"`
	VersionName  string `json:"version_name" form:"version_name" xml:"version_name" url:"version_name" validate:"required|string|maxLen:20" label:"App展示版本号"`
}

// ConfigValidation configures gookit/validate scenes.
func (RoadCommitServerRequest) ConfigValidation(v *validate.Validation) {
	v.WithScenes(validate.SValues{
		"insert": []string{"RoadCommitId", "DlLoaded", "UlLoaded", "IpAddress", "VersionCode", "VersionName"},
		"update": []string{"RoadCommitId", "DlLoaded", "UlLoaded", "IpAddress", "VersionCode", "VersionName"},
		"delete": []string{},
	})
}

// Messages defines Chinese validation messages.
func (RoadCommitServerRequest) Messages() map[string]string {
	return validate.MS{
		"int":      "{field}必须是整数",
		"maxLen":   "{field}长度不能大于 %v",
		"required": "{field}不能为空",
		"string":   "{field}必须是字符串",
	}
}
