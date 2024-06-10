package validate

import (
	"github.com/go-playground/validator/v10"
)

func ValidatePayload(data interface{}) error {
	validator := validator.New()

	err := validator.Struct(data)

	if err != nil {
		return err
	}

	return nil
}
