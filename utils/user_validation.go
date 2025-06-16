package utils

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/viveksingh-01/jarvis-auth/models"
)

var userValidator = validator.New()

func ValidateUser(user models.User) error {
	err := userValidator.Struct(user)
	if validationErrs, ok := err.(validator.ValidationErrors); ok {
		for _, fieldErr := range validationErrs {
			switch fieldErr.Tag() {
			case "required":
				return fmt.Errorf("%s is required", fieldErr.Field())
			case "min":
				return fmt.Errorf("%s must be at least %s characters long", fieldErr.Field(), fieldErr.Param())
			case "max":
				return fmt.Errorf("%s must be no more than %s characters", fieldErr.Field(), fieldErr.Param())
			default:
				return fmt.Errorf("validation failed on field %s", fieldErr.Field())
			}
		}
		return err
	}
	return nil
}
