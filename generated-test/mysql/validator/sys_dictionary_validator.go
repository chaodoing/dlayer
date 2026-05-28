package validator

import (
	"github.com/gookit/validate"
)

// SysDictionaryRequest is the request validator for the sys_dictionary table.
type SysDictionaryRequest struct {
	Id    uint   `json:"id" form:"id" xml:"id" url:"id" validate:"required|uint" label:"主键"`
	Title string `json:"title" form:"title" xml:"title" url:"title" validate:"required|string|maxLen:50" label:"字典名称"`
	Name  string `json:"name" form:"name" xml:"name" url:"name" validate:"required|string|maxLen:50" label:"调用名称"`
	Sort  uint   `json:"sort" form:"sort" xml:"sort" url:"sort" validate:"required|uint" label:"显示顺序"`
	Note  string `json:"note" form:"note" xml:"note" url:"note" validate:"required|string|maxLen:255" label:"字典备注"`
}

// ConfigValidation configures gookit/validate scenes.
func (SysDictionaryRequest) ConfigValidation(v *validate.Validation) {
	v.WithScenes(validate.SValues{
		"insert": []string{"Title", "Name", "Sort", "Note"},
		"update": []string{"Id", "Title", "Name", "Sort", "Note"},
		"delete": []string{"Id"},
	})
}

// Messages defines Chinese validation messages.
func (SysDictionaryRequest) Messages() map[string]string {
	return validate.MS{
		"maxLen":   "{field}长度不能大于 %v",
		"required": "{field}不能为空",
		"string":   "{field}必须是字符串",
		"uint":     "{field}必须是非负整数",
	}
}
