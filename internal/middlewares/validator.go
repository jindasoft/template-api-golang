package middlewares

import (
	"github.com/go-playground/validator/v10"
)

func ValidateTemplate(fl validator.FieldLevel) bool {
	return true
}
