package validator

import (
	"github.com/gookit/validate"
)

// DepartmentZonesRequest is the request validator for the department_zones table.
type DepartmentZonesRequest struct {
	DepartmentId int `json:"department_id" form:"department_id" xml:"department_id" url:"department_id" validate:"required|int" label:"部门ID"`
	ZoneCode     int `json:"zone_code" form:"zone_code" xml:"zone_code" url:"zone_code" validate:"required|int" label:"地区ID"`
}

// ConfigValidation configures gookit/validate scenes.
func (DepartmentZonesRequest) ConfigValidation(v *validate.Validation) {
	v.WithScenes(validate.SValues{
		"insert": []string{"DepartmentId", "ZoneCode"},
		"update": []string{"DepartmentId", "ZoneCode"},
		"delete": []string{},
	})
}

// Messages defines Chinese validation messages.
func (DepartmentZonesRequest) Messages() map[string]string {
	return validate.MS{
		"int":      "{field}必须是整数",
		"required": "{field}不能为空",
	}
}
