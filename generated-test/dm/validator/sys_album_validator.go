package validator

import (
	"github.com/gookit/validate"
)

// SysAlbumRequest is the request validator for the sys_album table.
type SysAlbumRequest struct {
	Id   int64  `json:"id" form:"id" xml:"id" url:"id" validate:"required|int" label:"主键ID"`
	Cid  int64  `json:"cid" form:"cid" xml:"cid" url:"cid" validate:"int" label:"类目ID"`
	Aid  int64  `json:"aid" form:"aid" xml:"aid" url:"aid" validate:"int" label:"管理员ID"`
	Uid  int64  `json:"uid" form:"uid" xml:"uid" url:"uid" validate:"int" label:"用户ID"`
	Type int    `json:"type" form:"type" xml:"type" url:"type" validate:"int" label:"文件类型: [10=图片, 20=视频]"`
	Name string `json:"name" form:"name" xml:"name" url:"name" validate:"required|string|maxLen:400" label:"文件名称"`
	Uri  string `json:"uri" form:"uri" xml:"uri" url:"uri" validate:"required|string|maxLen:800" label:"文件路径"`
	Ext  string `json:"ext" form:"ext" xml:"ext" url:"ext" validate:"required|string|maxLen:40" label:"文件扩展"`
	Size int64  `json:"size" form:"size" xml:"size" url:"size" validate:"int" label:"文件大小"`
}

// ConfigValidation configures gookit/validate scenes.
func (SysAlbumRequest) ConfigValidation(v *validate.Validation) {
	v.WithScenes(validate.SValues{
		"insert": []string{"Cid", "Aid", "Uid", "Type", "Name", "Uri", "Ext", "Size"},
		"update": []string{"Id", "Cid", "Aid", "Uid", "Type", "Name", "Uri", "Ext", "Size"},
		"delete": []string{"Id"},
	})
}

// Messages defines Chinese validation messages.
func (SysAlbumRequest) Messages() map[string]string {
	return validate.MS{
		"int":      "{field}必须是整数",
		"maxLen":   "{field}长度不能大于 %v",
		"required": "{field}不能为空",
		"string":   "{field}必须是字符串",
	}
}
