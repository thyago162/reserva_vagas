package internalerrors

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
)

var ErrValidation = errors.New("Validation error")
var validate = validator.New()

func ValidateStruct(obj interface{}) error {
	err := validate.Struct(obj)
	if err == nil {
		return nil
	}

	validationErrors, ok := err.(validator.ValidationErrors)
	if !ok || len(validationErrors) == 0 {
		return fmt.Errorf("%w, %s", ErrValidation, err.Error())
	}

	validationError := validationErrors[0]
	switch validationError.Tag() {
	case "required":
		return fmt.Errorf("%w: %s is required", ErrValidation, validationError.StructField())
	case "min":
		return fmt.Errorf("%w: %s is required with min %s", ErrValidation, validationError.StructField(), validationError.Param())
	case "max":
		return fmt.Errorf("%w: %s is required with max %s", ErrValidation, validationError.StructField(), validationError.Param())
	case "email":
		return fmt.Errorf("%w: %s is invalid", ErrValidation, validationError.StructField())
	default:
		return fmt.Errorf("%w: %s failed validation on rule %s", ErrValidation, validationError.StructField(), validationError.Param())
	}
}
