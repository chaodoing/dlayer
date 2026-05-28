package validator

import (
	"github.com/gookit/validate"
)

// SafeUtsRequest is the request validator for the safe_uts table.
type SafeUtsRequest struct {
	Id                 int    `json:"id" form:"id" xml:"id" url:"id" validate:"required|int" label:"主键"`
	PlatformId         int16  `json:"platform_id" form:"platform_id" xml:"platform_id" url:"platform_id" validate:"required|int" label:"来源平台"`
	Title              string `json:"title" form:"title" xml:"title" url:"title" validate:"required|string|maxLen:160" label:"探针设备标题"`
	Address            string `json:"address" form:"address" xml:"address" url:"address" validate:"required|string|maxLen:60" label:"IP地址"`
	DeviceId           string `json:"device_id" form:"device_id" xml:"device_id" url:"device_id" validate:"required|string|maxLen:144" label:"探针设备ID"`
	KafkaTopic         string `json:"kafka_topic" form:"kafka_topic" xml:"kafka_topic" url:"kafka_topic" validate:"required|string|maxLen:120" label:"kafka主题"`
	ElasticIndex       string `json:"elastic_index" form:"elastic_index" xml:"elastic_index" url:"elastic_index" validate:"required|string|maxLen:240" label:"es索引"`
	ElasticDetailIndex string `json:"elastic_detail_index" form:"elastic_detail_index" xml:"elastic_detail_index" url:"elastic_detail_index" validate:"required|string|maxLen:240" label:"对应详情索引"`
	Sort               int    `json:"sort" form:"sort" xml:"sort" url:"sort" validate:"required|int" label:"探针设备排序"`
	Disabled           int8   `json:"disabled" form:"disabled" xml:"disabled" url:"disabled" validate:"required|int" label:"禁用状态"`
}

// ConfigValidation configures gookit/validate scenes.
func (SafeUtsRequest) ConfigValidation(v *validate.Validation) {
	v.WithScenes(validate.SValues{
		"insert": []string{"PlatformId", "Title", "Address", "DeviceId", "KafkaTopic", "ElasticIndex", "ElasticDetailIndex", "Sort", "Disabled"},
		"update": []string{"Id", "PlatformId", "Title", "Address", "DeviceId", "KafkaTopic", "ElasticIndex", "ElasticDetailIndex", "Sort", "Disabled"},
		"delete": []string{"Id"},
	})
}

// Messages defines Chinese validation messages.
func (SafeUtsRequest) Messages() map[string]string {
	return validate.MS{
		"int":      "{field}必须是整数",
		"maxLen":   "{field}长度不能大于 %v",
		"required": "{field}不能为空",
		"string":   "{field}必须是字符串",
	}
}
