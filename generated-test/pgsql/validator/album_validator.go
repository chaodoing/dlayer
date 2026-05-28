package validator

import (
	"github.com/gookit/validate"
)

// AlbumRequest is the request validator for the album table.
type AlbumRequest struct {
	Id   int    `json:"id" form:"id" xml:"id" url:"id" validate:"required|int" label:"主键ID"`
	Pid  *int   `json:"pid" form:"pid" xml:"pid" url:"pid" validate:"int" label:"父级ID"`
	Type int    `json:"type" form:"type" xml:"type" url:"type" validate:"int" label:"类型: [1=图片, 2=视频]"`
	Name string `json:"name" form:"name" xml:"name" url:"name" validate:"required|string|maxLen:32" label:"分类名称"`
}

// ConfigValidation configures gookit/validate scenes.
func (AlbumRequest) ConfigValidation(v *validate.Validation) {
	v.WithScenes(validate.SValues{
		"insert": []string{"Pid", "Type", "Name"},
		"update": []string{"Id", "Pid", "Type", "Name"},
		"delete": []string{"Id"},
	})
}

// Messages defines Chinese validation messages.
func (AlbumRequest) Messages() map[string]string {
	return validate.MS{
		"int":      "{field}必须是整数",
		"maxLen":   "{field}长度不能大于 %v",
		"required": "{field}不能为空",
		"string":   "{field}必须是字符串",
	}
}
