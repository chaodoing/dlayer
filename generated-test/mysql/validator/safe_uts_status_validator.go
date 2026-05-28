package validator

import (
	"github.com/gookit/validate"
)

// SafeUtsStatusRequest is the request validator for the safe_uts_status table.
type SafeUtsStatusRequest struct {
	UtsId   uint   `json:"uts_id" form:"uts_id" xml:"uts_id" url:"uts_id" validate:"required|uint" label:"探针设备表ID"`
	Online  uint8  `json:"online" form:"online" xml:"online" url:"online" validate:"uint" label:"在线状态"`
	Status  uint8  `json:"status" form:"status" xml:"status" url:"status" validate:"required|uint" label:"采集状态\n1:normal (正常)\n2:warning (警告，但是可以正常工作) \n3:error (错误，无法正常工作)"`
	Total   uint64 `json:"total" form:"total" xml:"total" url:"total" validate:"required|uint" label:"总计流量"`
	Packet  int64  `json:"packet" form:"packet" xml:"packet" url:"packet" validate:"required|int" label:"总计数据包"`
	Cpu     uint8  `json:"cpu" form:"cpu" xml:"cpu" url:"cpu" validate:"uint" label:"CPU占用率"`
	Memory  uint8  `json:"memory" form:"memory" xml:"memory" url:"memory" validate:"required|uint" label:"内存占用率"`
	Disk    uint8  `json:"disk" form:"disk" xml:"disk" url:"disk" validate:"uint" label:"磁盘信息"`
	Runtime uint64 `json:"runtime" form:"runtime" xml:"runtime" url:"runtime" validate:"required|uint" label:"系统运行时长"`
}

// ConfigValidation configures gookit/validate scenes.
func (SafeUtsStatusRequest) ConfigValidation(v *validate.Validation) {
	v.WithScenes(validate.SValues{
		"insert": []string{"UtsId", "Online", "Status", "Total", "Packet", "Cpu", "Memory", "Disk", "Runtime"},
		"update": []string{"UtsId", "Online", "Status", "Total", "Packet", "Cpu", "Memory", "Disk", "Runtime"},
		"delete": []string{},
	})
}

// Messages defines Chinese validation messages.
func (SafeUtsStatusRequest) Messages() map[string]string {
	return validate.MS{
		"int":      "{field}必须是整数",
		"required": "{field}不能为空",
		"uint":     "{field}必须是非负整数",
	}
}
