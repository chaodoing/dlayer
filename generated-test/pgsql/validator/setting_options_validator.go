package validator

import (
	"github.com/gookit/validate"
	"gorm.io/datatypes"
)

// SettingOptionsRequest is the request validator for the setting_options table.
type SettingOptionsRequest struct {
	Id          int            `json:"id" form:"id" xml:"id" url:"id" validate:"required|int" label:"主键"`
	ConfigId    int            `json:"config_id" form:"config_id" xml:"config_id" url:"config_id" validate:"required|int" label:"配置ID"`
	Title       string         `json:"title" form:"title" xml:"title" url:"title" validate:"required|string|maxLen:50" label:"配置标题"`
	Code        string         `json:"code" form:"code" xml:"code" url:"code" validate:"required|string|maxLen:60" label:"配置编码"`
	Value       string         `json:"value" form:"value" xml:"value" url:"value" validate:"required|string" label:"配置值"`
	Options     datatypes.JSON `json:"options" form:"options" xml:"options" url:"options" validate:"required" label:"配置项"`
	Type        string         `json:"type" form:"type" xml:"type" url:"type" validate:"required|string|maxLen:32" label:"配置类型"`
	Sort        int            `json:"sort" form:"sort" xml:"sort" url:"sort" validate:"required|int" label:"排序"`
	Description string         `json:"description" form:"description" xml:"description" url:"description" validate:"required|string|maxLen:120" label:"配置说明"`
	Disabled    int            `json:"disabled" form:"disabled" xml:"disabled" url:"disabled" validate:"int" label:"锁定状态"`
}

// ConfigValidation configures gookit/validate scenes.
func (SettingOptionsRequest) ConfigValidation(v *validate.Validation) {
	v.WithScenes(validate.SValues{
		"insert": []string{"ConfigId", "Title", "Code", "Value", "Options", "Type", "Sort", "Description", "Disabled"},
		"update": []string{"Id", "ConfigId", "Title", "Code", "Value", "Options", "Type", "Sort", "Description", "Disabled"},
		"delete": []string{"Id"},
	})
}

// Messages defines Chinese validation messages.
func (SettingOptionsRequest) Messages() map[string]string {
	return validate.MS{
		"int":      "{field}必须是整数",
		"maxLen":   "{field}长度不能大于 %v",
		"required": "{field}不能为空",
		"string":   "{field}必须是字符串",
	}
}
