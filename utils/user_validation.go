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
			if fieldErr.Tag() == "required" {
				return fmt.Errorf("%s is required", fieldErr.Field())
			}
		}
		return err
	}
	return nil
}
