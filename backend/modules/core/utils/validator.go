package core

import (
	core "github.com/aashu10sh/authapi/modules/core/entities/errors"
	"github.com/go-playground/validator/v10"
)

func ValidateStruct[T any](object T) []*core.ErrorResponse {
	var errors []*core.ErrorResponse
	validate := validator.New()
	err := validate.Struct(object)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element core.ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}
