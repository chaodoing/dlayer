package validator

import (
	"github.com/gookit/validate"
)

// SafeUtsDdosRequest is the request validator for the safe_uts_ddos table.
type SafeUtsDdosRequest struct {
	Id         int    `json:"id" form:"id" xml:"id" url:"id" validate:"required|int" label:"主键"`
	AddressId  int64  `json:"address_id" form:"address_id" xml:"address_id" url:"address_id" validate:"required|int" label:"受攻击目标"`
	Sip        string `json:"sip" form:"sip" xml:"sip" url:"sip" validate:"required|string" label:"攻击来源IP"`
	StartTime  int    `json:"start_time" form:"start_time" xml:"start_time" url:"start_time" validate:"required|int" label:"开始攻击时间"`
	StopTime   int    `json:"stop_time" form:"stop_time" xml:"stop_time" url:"stop_time" validate:"int" label:"结束攻击时间"`
	Domestic   int8   `json:"domestic" form:"domestic" xml:"domestic" url:"domestic" validate:"required|int" label:"是否境内攻击"`
	Duration   int    `json:"duration" form:"duration" xml:"duration" url:"duration" validate:"required|int" label:"攻击时长单位:秒"`
	AttackType string `json:"attack_type" form:"attack_type" xml:"attack_type" url:"attack_type" validate:"required|string|maxLen:48" label:"攻击类型"`
}

// ConfigValidation configures gookit/validate scenes.
func (SafeUtsDdosRequest) ConfigValidation(v *validate.Validation) {
	v.WithScenes(validate.SValues{
		"insert": []string{"AddressId", "Sip", "StartTime", "StopTime", "Domestic", "Duration", "AttackType"},
		"update": []string{"Id", "AddressId", "Sip", "StartTime", "StopTime", "Domestic", "Duration", "AttackType"},
		"delete": []string{"Id"},
	})
}

// Messages defines Chinese validation messages.
func (SafeUtsDdosRequest) Messages() map[string]string {
	return validate.MS{
		"int":      "{field}必须是整数",
		"maxLen":   "{field}长度不能大于 %v",
		"required": "{field}不能为空",
		"string":   "{field}必须是字符串",
	}
}
