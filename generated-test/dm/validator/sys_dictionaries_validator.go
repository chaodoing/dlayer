package validator

import (
	"github.com/gookit/validate"
)

// SysDictionariesRequest is the request validator for the sys_dictionaries table.
type SysDictionariesRequest struct {
	Id       int64  `json:"id" form:"id" xml:"id" url:"id" validate:"required|int" label:"主键"`
	DictId   int64  `json:"dict_id" form:"dict_id" xml:"dict_id" url:"dict_id" validate:"required|int" label:"字典类型id"`
	Title    string `json:"title" form:"title" xml:"title" url:"title" validate:"required|string|maxLen:200" label:"字典名称"`
	Value    string `json:"value" form:"value" xml:"value" url:"value" validate:"required|string" label:"字典内容"`
	Sort     int    `json:"sort" form:"sort" xml:"sort" url:"sort" validate:"required|int" label:"显示顺序"`
	Disabled int    `json:"disabled" form:"disabled" xml:"disabled" url:"disabled" validate:"int" label:"禁用状态"`
}

// ConfigValidation configures gookit/validate scenes.
func (SysDictionariesRequest) ConfigValidation(v *validate.Validation) {
	v.WithScenes(validate.SValues{
		"insert": []string{"DictId", "Title", "Value", "Sort", "Disabled"},
		"update": []string{"Id", "DictId", "Title", "Value", "Sort", "Disabled"},
		"delete": []string{"Id"},
	})
}

// Messages defines Chinese validation messages.
func (SysDictionariesRequest) Messages() map[string]string {
	return validate.MS{
		"int":      "{field}必须是整数",
		"maxLen":   "{field}长度不能大于 %v",
		"required": "{field}不能为空",
		"string":   "{field}必须是字符串",
	}
}
