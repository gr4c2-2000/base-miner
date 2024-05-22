package common

import "github.com/go-playground/validator/v10"

var Validat *validator.Validate

func GetValidator() *validator.Validate {
	if Validat == nil {
		Validat = validator.New()
	}
	return Validat
}
