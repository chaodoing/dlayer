package validator

import (
	"github.com/gookit/validate"
)

// SafeUtsTrafficLogRequest is the request validator for the safe_uts_traffic_log table.
type SafeUtsTrafficLogRequest struct {
	Id            uint64 `json:"id" form:"id" xml:"id" url:"id" validate:"required|uint" label:"日志id"`
	UtsId         uint   `json:"uts_id" form:"uts_id" xml:"uts_id" url:"uts_id" validate:"required|uint" label:"关联的探针设备"`
	Bps           uint64 `json:"bps" form:"bps" xml:"bps" url:"bps" validate:"required|uint" label:"接收流量比特率大小"`
	Pps           uint64 `json:"pps" form:"pps" xml:"pps" url:"pps" validate:"required|uint" label:"接收流量的每秒包数"`
	TotalBytes    uint64 `json:"total_bytes" form:"total_bytes" xml:"total_bytes" url:"total_bytes" validate:"required|uint" label:"接收流量总字节数大小"`
	TotalPackets  uint64 `json:"total_packets" form:"total_packets" xml:"total_packets" url:"total_packets" validate:"required|uint" label:"接收流量总包数大小"`
	InterfaceName string `json:"interface_name" form:"interface_name" xml:"interface_name" url:"interface_name" validate:"required|string|maxLen:40" label:"如果非单条接口信息 即为总流量信息则interfacename值为\"\""`
	Time          uint64 `json:"time" form:"time" xml:"time" url:"time" validate:"required|uint" label:"上传时间"`
}

// ConfigValidation configures gookit/validate scenes.
func (SafeUtsTrafficLogRequest) ConfigValidation(v *validate.Validation) {
	v.WithScenes(validate.SValues{
		"insert": []string{"UtsId", "Bps", "Pps", "TotalBytes", "TotalPackets", "InterfaceName", "Time"},
		"update": []string{"Id", "UtsId", "Bps", "Pps", "TotalBytes", "TotalPackets", "InterfaceName", "Time"},
		"delete": []string{"Id"},
	})
}

// Messages defines Chinese validation messages.
func (SafeUtsTrafficLogRequest) Messages() map[string]string {
	return validate.MS{
		"maxLen":   "{field}长度不能大于 %v",
		"required": "{field}不能为空",
		"string":   "{field}必须是字符串",
		"uint":     "{field}必须是非负整数",
	}
}
