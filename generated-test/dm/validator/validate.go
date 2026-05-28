package validator

import "github.com/gookit/validate"

const (
	SceneInsert = "insert"
	SceneUpdate = "update"
	SceneDelete = "delete"
)

// Validate checks a generated request struct with gookit/validate.
func Validate(value any, scenes ...string) error {
	v := validate.Struct(value)
	if v.Validate(scenes...) {
		return nil
	}
	return v.Errors.ErrOrNil()
}
