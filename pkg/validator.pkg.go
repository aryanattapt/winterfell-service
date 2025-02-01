package pkg

import (
	"github.com/go-playground/validator/v10"
)

func ValidateForm(err error) map[string]interface{} {
	errors := make(map[string]interface{})
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			field := LowercaseFirstChar(e.StructField())
			errors[field] = map[string]interface{}{
				"message": getMessageForTag(e),
			}
		}
	}
	return errors
}

func getMessageForTag(e validator.FieldError) string {
	var fieldName = LowercaseFirstChar(e.StructField())
	switch e.Tag() {
	case "required":
		return fieldName + " is required. For example: 'exampleValue'."
	case "email":
		return fieldName + " must be a valid email address. For example: 'user@example.com'."
	case "e164":
		return fieldName + " must be a valid phone number. For example: '+1234567890'."
	case "uuid":
		return fieldName + " must be a valid UUID. For example: '123e4567-e89b-12d3-a456-426614174000'."
	case "url":
		return fieldName + " must be a valid URL. For example: 'https://www.example.com'."
	case "http_url":
		return fieldName + " must be a valid HTTP URL. For example: 'http://www.example.com'."
	case "https_url":
		return fieldName + " must be a valid HTTPS URL. For example: 'https://www.example.com'."
	case "date":
		return fieldName + " must be a valid date. For example: '2024-09-12'."
	case "datetime":
		return fieldName + " must be a valid datetime. For example: '2024-09-12T15:04:05Z'."
	case "ascii":
		return fieldName + " must contain only ASCII characters. For example: 'HelloWorld123'."
	case "alphanum":
		return fieldName + " must be alphanumeric. For example: 'Hello123'."
	case "alpha":
		return fieldName + " must contain only alphabetic characters. For example: 'HelloWorld'."
	case "number":
		return fieldName + " must be a number. For example: '12345'."
	case "len":
		return fieldName + " length must be exactly as defined by the rule. For example: '12345'."
	case "min":
		return fieldName + " length must be at least as defined by the rule. For example: '12345'."
	case "max":
		return fieldName + " length must be at most as defined by the rule. For example: '12345'."
	case "gte":
		return fieldName + " must be greater than or equal to as defined by the rule. For example: '100'."
	case "lte":
		return fieldName + " must be less than or equal to as defined by the rule. For example: '100'."
	case "gt":
		return fieldName + " must be greater than as defined by the rule. For example: '101'."
	case "lt":
		return fieldName + " must be less than as defined by the rule. For example: '99'."
	default:
		return e.Error()
	}
}
