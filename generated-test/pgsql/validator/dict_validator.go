package validator

import (
	"github.com/gookit/validate"
)

// DictRequest is the request validator for the dict table.
type DictRequest struct {
	Id          int    `json:"id" form:"id" xml:"id" url:"id" validate:"required|int" label:"主键"`
	Title       string `json:"title" form:"title" xml:"title" url:"title" validate:"required|string|maxLen:50" label:"字典名称"`
	Name        string `json:"name" form:"name" xml:"name" url:"name" validate:"required|string|maxLen:50" label:"调用名称"`
	Sort        int    `json:"sort" form:"sort" xml:"sort" url:"sort" validate:"required|int" label:"显示顺序"`
	Description string `json:"description" form:"description" xml:"description" url:"description" validate:"required|string|maxLen:255" label:"字典备注"`
}

// ConfigValidation configures gookit/validate scenes.
func (DictRequest) ConfigValidation(v *validate.Validation) {
	v.WithScenes(validate.SValues{
		"insert": []string{"Title", "Name", "Sort", "Description"},
		"update": []string{"Id", "Title", "Name", "Sort", "Description"},
		"delete": []string{"Id"},
	})
}

// Messages defines Chinese validation messages.
func (DictRequest) Messages() map[string]string {
	return validate.MS{
		"int":      "{field}必须是整数",
		"maxLen":   "{field}长度不能大于 %v",
		"required": "{field}不能为空",
		"string":   "{field}必须是字符串",
	}
}
