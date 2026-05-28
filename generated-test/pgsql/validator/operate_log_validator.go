package validator

import (
	"github.com/gookit/validate"
	"gorm.io/datatypes"
)

// OperateLogRequest is the request validator for the operate_log table.
type OperateLogRequest struct {
	Id      int            `json:"id" form:"id" xml:"id" url:"id" validate:"required|int" label:"主键"`
	Method  string         `json:"method" form:"method" xml:"method" url:"method" validate:"required|string|maxLen:10" label:"请求方式"`
	Path    string         `json:"path" form:"path" xml:"path" url:"path" validate:"required|string|maxLen:255" label:"请求地址"`
	Enctype string         `json:"enctype" form:"enctype" xml:"enctype" url:"enctype" validate:"required|string|maxLen:60" label:"表单类型"`
	Query   string         `json:"query" form:"query" xml:"query" url:"query" validate:"required|string|maxLen:255" label:"请求url参数"`
	Body    datatypes.JSON `json:"body" form:"body" xml:"body" url:"body" validate:"required" label:"请求内容"`
	Remote  string         `json:"remote" form:"remote" xml:"remote" url:"remote" validate:"required|string|maxLen:60" label:"远程地址"`
	Uid     int            `json:"uid" form:"uid" xml:"uid" url:"uid" validate:"int" label:"用户id"`
}

// ConfigValidation configures gookit/validate scenes.
func (OperateLogRequest) ConfigValidation(v *validate.Validation) {
	v.WithScenes(validate.SValues{
		"insert": []string{"Method", "Path", "Enctype", "Query", "Body", "Remote", "Uid"},
		"update": []string{"Id", "Method", "Path", "Enctype", "Query", "Body", "Remote", "Uid"},
		"delete": []string{"Id"},
	})
}

// Messages defines Chinese validation messages.
func (OperateLogRequest) Messages() map[string]string {
	return validate.MS{
		"int":      "{field}必须是整数",
		"maxLen":   "{field}长度不能大于 %v",
		"required": "{field}不能为空",
		"string":   "{field}必须是字符串",
	}
}
