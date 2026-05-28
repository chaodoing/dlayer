package validator

import (
	"github.com/gookit/validate"
)

// SysAlbumCategoryRequest is the request validator for the sys_album_category table.
type SysAlbumCategoryRequest struct {
	Id   int64  `json:"id" form:"id" xml:"id" url:"id" validate:"required|int" label:"主键ID"`
	Pid  int64  `json:"pid" form:"pid" xml:"pid" url:"pid" validate:"int" label:"父级ID"`
	Type int    `json:"type" form:"type" xml:"type" url:"type" validate:"int" label:"类型: [10=图片, 20=视频]"`
	Name string `json:"name" form:"name" xml:"name" url:"name" validate:"required|string|maxLen:128" label:"分类名称"`
}

// ConfigValidation configures gookit/validate scenes.
func (SysAlbumCategoryRequest) ConfigValidation(v *validate.Validation) {
	v.WithScenes(validate.SValues{
		"insert": []string{"Id", "Pid", "Type", "Name"},
		"update": []string{"Id", "Pid", "Type", "Name"},
		"delete": []string{"Id"},
	})
}

// Messages defines Chinese validation messages.
func (SysAlbumCategoryRequest) Messages() map[string]string {
	return validate.MS{
		"int":      "{field}必须是整数",
		"maxLen":   "{field}长度不能大于 %v",
		"required": "{field}不能为空",
		"string":   "{field}必须是字符串",
	}
}
