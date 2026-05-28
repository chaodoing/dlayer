package validator

import (
	"github.com/gookit/validate"
)

// SafeEventsRequest is the request validator for the safe_events table.
type SafeEventsRequest struct {
	Id          uint64 `json:"id" form:"id" xml:"id" url:"id" validate:"required|uint" label:"主键"`
	NetType     uint8  `json:"net_type" form:"net_type" xml:"net_type" url:"net_type" validate:"uint" label:"网络类型"`
	EventType   string `json:"event_type" form:"event_type" xml:"event_type" url:"event_type" validate:"required|string|maxLen:16" label:"事件编码"`
	SrcIpv4     string `json:"src_ipv4" form:"src_ipv4" xml:"src_ipv4" url:"src_ipv4" validate:"required|string|maxLen:15" label:"来源IP"`
	SrcPort     uint32 `json:"src_port" form:"src_port" xml:"src_port" url:"src_port" validate:"required|uint" label:"来源端口"`
	SrcIpv6     string `json:"src_ipv6" form:"src_ipv6" xml:"src_ipv6" url:"src_ipv6" validate:"required|string|maxLen:64" label:"来源IPV6"`
	DistIpv4    string `json:"dist_ipv4" form:"dist_ipv4" xml:"dist_ipv4" url:"dist_ipv4" validate:"required|string|maxLen:15" label:"目标IPv4"`
	DistPort    uint32 `json:"dist_port" form:"dist_port" xml:"dist_port" url:"dist_port" validate:"required|uint" label:"目标端口"`
	DistIpv6    string `json:"dist_ipv6" form:"dist_ipv6" xml:"dist_ipv6" url:"dist_ipv6" validate:"required|string|maxLen:64" label:"目标IPv6"`
	PlatformId  uint16 `json:"platform_id" form:"platform_id" xml:"platform_id" url:"platform_id" validate:"required|uint" label:"来源平台"`
	EventLevel  uint8  `json:"event_level" form:"event_level" xml:"event_level" url:"event_level" validate:"uint" label:"事件等级0未定义 1低2中3高"`
	DiscoveryAt uint   `json:"discovery_at" form:"discovery_at" xml:"discovery_at" url:"discovery_at" validate:"required|uint" label:"发现时间"`
}

// ConfigValidation configures gookit/validate scenes.
func (SafeEventsRequest) ConfigValidation(v *validate.Validation) {
	v.WithScenes(validate.SValues{
		"insert": []string{"NetType", "EventType", "SrcIpv4", "SrcPort", "SrcIpv6", "DistIpv4", "DistPort", "DistIpv6", "PlatformId", "EventLevel", "DiscoveryAt"},
		"update": []string{"Id", "NetType", "EventType", "SrcIpv4", "SrcPort", "SrcIpv6", "DistIpv4", "DistPort", "DistIpv6", "PlatformId", "EventLevel", "DiscoveryAt"},
		"delete": []string{"Id"},
	})
}

// Messages defines Chinese validation messages.
func (SafeEventsRequest) Messages() map[string]string {
	return validate.MS{
		"maxLen":   "{field}长度不能大于 %v",
		"required": "{field}不能为空",
		"string":   "{field}必须是字符串",
		"uint":     "{field}必须是非负整数",
	}
}
