package validator

import (
	"github.com/gookit/validate"
)

// SysAlbumRequest is the request validator for the sys_album table.
type SysAlbumRequest struct {
	Id   uint   `json:"id" form:"id" xml:"id" url:"id" validate:"required|uint" label:"主键ID"`
	Cid  uint   `json:"cid" form:"cid" xml:"cid" url:"cid" validate:"uint" label:"类目ID"`
	Aid  uint   `json:"aid" form:"aid" xml:"aid" url:"aid" validate:"uint" label:"管理员ID"`
	Uid  uint   `json:"uid" form:"uid" xml:"uid" url:"uid" validate:"uint" label:"用户ID"`
	Type uint8  `json:"type" form:"type" xml:"type" url:"type" validate:"uint" label:"文件类型: [10=图片, 20=视频]"`
	Name string `json:"name" form:"name" xml:"name" url:"name" validate:"required|string|maxLen:100" label:"文件名称"`
	Uri  string `json:"uri" form:"uri" xml:"uri" url:"uri" validate:"required|string|maxLen:200" label:"文件路径"`
	Ext  string `json:"ext" form:"ext" xml:"ext" url:"ext" validate:"required|string|maxLen:10" label:"文件扩展"`
	Size uint   `json:"size" form:"size" xml:"size" url:"size" validate:"uint" label:"文件大小"`
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
		"maxLen":   "{field}长度不能大于 %v",
		"required": "{field}不能为空",
		"string":   "{field}必须是字符串",
		"uint":     "{field}必须是非负整数",
	}
}
