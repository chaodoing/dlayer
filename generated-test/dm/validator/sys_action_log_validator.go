package validator

import (
	"github.com/gookit/validate"
)

// SysActionLogRequest is the request validator for the sys_action_log table.
type SysActionLogRequest struct {
	Id      int64   `json:"id" form:"id" xml:"id" url:"id" validate:"required|int" label:"主键"`
	Method  string  `json:"method" form:"method" xml:"method" url:"method" validate:"required|string|maxLen:24" label:"请求方式"`
	Path    string  `json:"path" form:"path" xml:"path" url:"path" validate:"required|string|maxLen:1020" label:"请求地址"`
	Enctype string  `json:"enctype" form:"enctype" xml:"enctype" url:"enctype" validate:"required|string|maxLen:240" label:"表单类型"`
	Query   string  `json:"query" form:"query" xml:"query" url:"query" validate:"required|string|maxLen:1020" label:"请求url参数"`
	Body    string  `json:"body" form:"body" xml:"body" url:"body" validate:"required|string|maxLen:32767" label:"请求内容"`
	Remote  string  `json:"remote" form:"remote" xml:"remote" url:"remote" validate:"required|string|maxLen:240" label:"远程地址"`
	Uid     float64 `json:"uid" form:"uid" xml:"uid" url:"uid" validate:"float" label:"用户id"`
}

// ConfigValidation configures gookit/validate scenes.
func (SysActionLogRequest) ConfigValidation(v *validate.Validation) {
	v.WithScenes(validate.SValues{
		"insert": []string{"Method", "Path", "Enctype", "Query", "Body", "Remote", "Uid"},
		"update": []string{"Id", "Method", "Path", "Enctype", "Query", "Body", "Remote", "Uid"},
		"delete": []string{"Id"},
	})
}

// Messages defines Chinese validation messages.
func (SysActionLogRequest) Messages() map[string]string {
	return validate.MS{
		"float":    "{field}必须是浮点数",
		"int":      "{field}必须是整数",
		"maxLen":   "{field}长度不能大于 %v",
		"required": "{field}不能为空",
		"string":   "{field}必须是字符串",
	}
}
