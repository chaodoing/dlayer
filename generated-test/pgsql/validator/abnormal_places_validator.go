package validator

import (
	"github.com/gookit/validate"
	"gorm.io/datatypes"
)

// AbnormalPlacesRequest is the request validator for the abnormal_places table.
type AbnormalPlacesRequest struct {
	Id          int            `json:"id" form:"id" xml:"id" url:"id" validate:"required|int" label:"自增ID"`
	PlaceId     int            `json:"place_id" form:"place_id" xml:"place_id" url:"place_id" validate:"required|int" label:"场所ID"`
	SceneId     int            `json:"scene_id" form:"scene_id" xml:"scene_id" url:"scene_id" validate:"required|int" label:"场景二级ID"`
	Uid         int            `json:"uid" form:"uid" xml:"uid" url:"uid" validate:"required|int" label:"用户ID"`
	Reporter    string         `json:"reporter" form:"reporter" xml:"reporter" url:"reporter" validate:"required|string|maxLen:20" label:"报告人姓名"`
	Mobile      string         `json:"mobile" form:"mobile" xml:"mobile" url:"mobile" validate:"required|string|maxLen:16" label:"报告人电话"`
	Description string         `json:"description" form:"description" xml:"description" url:"description" validate:"required|string|maxLen:50" label:"描述"`
	Reason      int            `json:"reason" form:"reason" xml:"reason" url:"reason" validate:"required|int" label:"原因类型 0:未知原因 1:无信号 2:无法测试 3:场所名称错误 4:区县归类错误 5:经纬度错误 6:场所不存在 7:其它异常情况 "`
	Done        bool           `json:"done" form:"done" xml:"done" url:"done" validate:"required|bool" label:"修改完成状态"`
	Content     datatypes.JSON `json:"content" form:"content" xml:"content" url:"content" validate:"required" label:"上报修改内容"`
}

// ConfigValidation configures gookit/validate scenes.
func (AbnormalPlacesRequest) ConfigValidation(v *validate.Validation) {
	v.WithScenes(validate.SValues{
		"insert": []string{"PlaceId", "SceneId", "Uid", "Reporter", "Mobile", "Description", "Reason", "Done", "Content"},
		"update": []string{"Id", "PlaceId", "SceneId", "Uid", "Reporter", "Mobile", "Description", "Reason", "Done", "Content"},
		"delete": []string{"Id"},
	})
}

// Messages defines Chinese validation messages.
func (AbnormalPlacesRequest) Messages() map[string]string {
	return validate.MS{
		"bool":     "{field}必须是布尔值",
		"int":      "{field}必须是整数",
		"maxLen":   "{field}长度不能大于 %v",
		"required": "{field}不能为空",
		"string":   "{field}必须是字符串",
	}
}
