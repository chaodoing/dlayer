package validator

import (
	"github.com/gookit/validate"
	"gorm.io/datatypes"
)

// RoadMassCommitSpeedDetailsRequest is the request validator for the road_mass_commit_speed_details table.
type RoadMassCommitSpeedDetailsRequest struct {
	CommitSpeedId int            `json:"commit_speed_id" form:"commit_speed_id" xml:"commit_speed_id" url:"commit_speed_id" validate:"required|int" label:"路测众测提交上下行速率ID"`
	Upload        datatypes.JSON `json:"upload" form:"upload" xml:"upload" url:"upload" validate:"required" label:"上传速率"`
	Download      datatypes.JSON `json:"download" form:"download" xml:"download" url:"download" validate:"required" label:"下载速率"`
}

// ConfigValidation configures gookit/validate scenes.
func (RoadMassCommitSpeedDetailsRequest) ConfigValidation(v *validate.Validation) {
	v.WithScenes(validate.SValues{
		"insert": []string{"CommitSpeedId", "Upload", "Download"},
		"update": []string{"CommitSpeedId", "Upload", "Download"},
		"delete": []string{},
	})
}

// Messages defines Chinese validation messages.
func (RoadMassCommitSpeedDetailsRequest) Messages() map[string]string {
	return validate.MS{
		"int":      "{field}必须是整数",
		"required": "{field}不能为空",
	}
}
