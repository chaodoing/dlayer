package validator

import (
	"github.com/gookit/validate"
)

// SysRuleRequest is the request validator for the sys_rule table.
type SysRuleRequest struct {
	Id       int64  `json:"id" form:"id" xml:"id" url:"id" validate:"required|int" label:"主键id"`
	Pid      int64  `json:"pid" form:"pid" xml:"pid" url:"pid" validate:"int" label:"上级菜单"`
	Title    string `json:"title" form:"title" xml:"title" url:"title" validate:"required|string|maxLen:120" label:"菜单标题"`
	Name     string `json:"name" form:"name" xml:"name" url:"name" validate:"required|string|maxLen:240" label:"唯一名称 可用于权限检查"`
	Icon     string `json:"icon" form:"icon" xml:"icon" url:"icon" validate:"required|string|maxLen:240" label:"菜单图标"`
	Method   string `json:"method" form:"method" xml:"method" url:"method" validate:"required|string|maxLen:24" label:"请求方法"`
	Href     string `json:"href" form:"href" xml:"href" url:"href" validate:"required|string|maxLen:480" label:"URL地址"`
	Target   string `json:"target" form:"target" xml:"target" url:"target" validate:"required|string|maxLen:28" label:"打开方式"`
	Mode     int    `json:"mode" form:"mode" xml:"mode" url:"mode" validate:"int" label:"类型 0 目录 1 菜单 2 操作"`
	Sort     int64  `json:"sort" form:"sort" xml:"sort" url:"sort" validate:"required|int" label:"菜单排序"`
	Visible  int    `json:"visible" form:"visible" xml:"visible" url:"visible" validate:"int" label:"显示的"`
	Disabled int    `json:"disabled" form:"disabled" xml:"disabled" url:"disabled" validate:"int" label:"禁用的"`
}

// ConfigValidation configures gookit/validate scenes.
func (SysRuleRequest) ConfigValidation(v *validate.Validation) {
	v.WithScenes(validate.SValues{
		"insert": []string{"Pid", "Title", "Name", "Icon", "Method", "Href", "Target", "Mode", "Sort", "Visible", "Disabled"},
		"update": []string{"Id", "Pid", "Title", "Name", "Icon", "Method", "Href", "Target", "Mode", "Sort", "Visible", "Disabled"},
		"delete": []string{"Id"},
	})
}

// Messages defines Chinese validation messages.
func (SysRuleRequest) Messages() map[string]string {
	return validate.MS{
		"int":      "{field}必须是整数",
		"maxLen":   "{field}长度不能大于 %v",
		"required": "{field}不能为空",
		"string":   "{field}必须是字符串",
	}
}
