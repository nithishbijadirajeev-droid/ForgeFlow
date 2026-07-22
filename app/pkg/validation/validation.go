package validation

import (
	"errors"
	"strings"

	"github.com/go-playground/validator/v10"
)

type FieldError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func FormatValidationErrors(err error) interface{} {
	var validationErrs validator.ValidationErrors

	if !errors.As(err, &validationErrs) {
		return err.Error()
	}

	var formatted []FieldError

	for _, e := range validationErrs {
		field := strings.ToLower(e.Field())

		message := e.Error()

		switch e.Tag() {
		case "required":
			message = field + " is required"
		case "email":
			message = "invalid email address"
		case "min":
			message = field + " must be at least " + e.Param() + " characters"
		case "max":
			message = field + " must be at most " + e.Param() + " characters"
		case "url":
			message = "invalid URL"
		}

		formatted = append(formatted, FieldError{
			Field:   field,
			Message: message,
		})
	}

	return formatted
}
