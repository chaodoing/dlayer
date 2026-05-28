package validator

import (
	"github.com/gookit/validate"
	"gorm.io/datatypes"
)

// RoadMassCommitIndicatorSignalsRequest is the request validator for the road_mass_commit_indicator_signals table.
type RoadMassCommitIndicatorSignalsRequest struct {
	Id          int            `json:"id" form:"id" xml:"id" url:"id" validate:"required|int" label:"主键ID"`
	IndicatorId int            `json:"indicator_id" form:"indicator_id" xml:"indicator_id" url:"indicator_id" validate:"required|int" label:"路测众测指标数据ID"`
	Type        string         `json:"type" form:"type" xml:"type" url:"type" validate:"required|string|maxLen:10" label:"网络类型"`
	T           int            `json:"t" form:"t" xml:"t" url:"t" validate:"required|int" label:"采集时间"`
	Timestamp   int            `json:"timestamp" form:"timestamp" xml:"timestamp" url:"timestamp" validate:"required|int" label:"记录生成时间"`
	Band        string         `json:"band" form:"band" xml:"band" url:"band" validate:"required|string|maxLen:30" label:"频段信息"`
	Freq        datatypes.JSON `json:"freq" form:"freq" xml:"freq" url:"freq" validate:"required" label:"频率[下行频率 MHz, 上行频率 MHz]"`
	Level       int            `json:"level" form:"level" xml:"level" url:"level" validate:"required|int" label:"信号等级"`
	Sinr        int            `json:"sinr" form:"sinr" xml:"sinr" url:"sinr" validate:"required|int" label:"信噪比"`
	Rsrp        int            `json:"rsrp" form:"rsrp" xml:"rsrp" url:"rsrp" validate:"required|int" label:"信号接收功率"`
	Rsrq        int            `json:"rsrq" form:"rsrq" xml:"rsrq" url:"rsrq" validate:"required|int" label:"信号质量"`
	Arfcn       int            `json:"arfcn" form:"arfcn" xml:"arfcn" url:"arfcn" validate:"required|int" label:"频点编号"`
}

// ConfigValidation configures gookit/validate scenes.
func (RoadMassCommitIndicatorSignalsRequest) ConfigValidation(v *validate.Validation) {
	v.WithScenes(validate.SValues{
		"insert": []string{"IndicatorId", "Type", "T", "Timestamp", "Band", "Freq", "Level", "Sinr", "Rsrp", "Rsrq", "Arfcn"},
		"update": []string{"Id", "IndicatorId", "Type", "T", "Timestamp", "Band", "Freq", "Level", "Sinr", "Rsrp", "Rsrq", "Arfcn"},
		"delete": []string{"Id"},
	})
}

// Messages defines Chinese validation messages.
func (RoadMassCommitIndicatorSignalsRequest) Messages() map[string]string {
	return validate.MS{
		"int":      "{field}必须是整数",
		"maxLen":   "{field}长度不能大于 %v",
		"required": "{field}不能为空",
		"string":   "{field}必须是字符串",
	}
}
