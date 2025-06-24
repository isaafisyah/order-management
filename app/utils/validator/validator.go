package validator

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func ValidateStruct(s interface{}) error {
	err := validate.Struct(s)
	if err == nil {
		return nil
	}

	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		var messages []string
		for _, e := range validationErrors {
			// Ubah format error jadi lebih user friendly
			messages = append(messages, fmt.Sprintf("Field %s is %s", e.Field(), e.Tag()))
		}
		return fmt.Errorf(strings.Join(messages, ", "))
	}

	return err
}