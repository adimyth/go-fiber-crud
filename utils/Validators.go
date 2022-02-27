package utils

import (
	"github.com/adimyth/go-fiber-crud/schema"
	"github.com/go-playground/validator"
)

var validate = validator.New()

func ValidateStruct(todo schema.ToDo) []*schema.ErrorResponse {
	var errors []*schema.ErrorResponse

	err := validate.Struct(todo)

	if err != nil {

		for _, err := range err.(validator.ValidationErrors) {
			var element schema.ErrorResponse
			element.Failed = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}
