package validator

import (
	"github.com/gookit/validate"
)

// ExportFilesRequest is the request validator for the export_files table.
type ExportFilesRequest struct {
	Id       int    `json:"id" form:"id" xml:"id" url:"id" validate:"required|int" label:"导出文件ID"`
	Uid      int    `json:"uid" form:"uid" xml:"uid" url:"uid" validate:"required|int" label:"后台用户ID"`
	File     string `json:"file" form:"file" xml:"file" url:"file" validate:"required|string|maxLen:255" label:"文件名称"`
	State    int    `json:"state" form:"state" xml:"state" url:"state" validate:"required|int" label:"状态 0:导出中 1:已完成 2:有错误"`
	Progress int    `json:"progress" form:"progress" xml:"progress" url:"progress" validate:"int" label:"导出进度"`
	Message  string `json:"message" form:"message" xml:"message" url:"message" validate:"required|string|maxLen:200" label:"错误消息"`
	Complete bool   `json:"complete" form:"complete" xml:"complete" url:"complete" validate:"bool" label:"下载完毕"`
	Unlink   bool   `json:"unlink" form:"unlink" xml:"unlink" url:"unlink" validate:"bool" label:"删除文件"`
	Md5      string `json:"md5" form:"md5" xml:"md5" url:"md5" validate:"required|string|maxLen:32" label:"文件MD5"`
}

// ConfigValidation configures gookit/validate scenes.
func (ExportFilesRequest) ConfigValidation(v *validate.Validation) {
	v.WithScenes(validate.SValues{
		"insert": []string{"Uid", "File", "State", "Progress", "Message", "Complete", "Unlink", "Md5"},
		"update": []string{"Id", "Uid", "File", "State", "Progress", "Message", "Complete", "Unlink", "Md5"},
		"delete": []string{"Id"},
	})
}

// Messages defines Chinese validation messages.
func (ExportFilesRequest) Messages() map[string]string {
	return validate.MS{
		"bool":     "{field}必须是布尔值",
		"int":      "{field}必须是整数",
		"maxLen":   "{field}长度不能大于 %v",
		"required": "{field}不能为空",
		"string":   "{field}必须是字符串",
	}
}
