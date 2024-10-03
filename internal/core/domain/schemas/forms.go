package schema

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

type GenericForm[T any] struct {
	Data T
}

func FormValidator[T any](form GenericForm[T]) map[string][]string {
	validate := validator.New()
	if err := validate.Struct(form.Data); err != nil {
		errors := make(map[string][]string)
		for _, err := range err.(validator.ValidationErrors) {
			field := strings.ToLower(err.Field())
			var errorMsg string
			switch err.Tag() {
			case "required":
				errorMsg = "Campo obligatorio."
			case "email":
				errorMsg = "Ingresa un correo válido."
			case "numeric":
				errorMsg = "Ingresa un valor numérico."
			case "min":
				errorMsg = fmt.Sprintf("Mínimo %s carácteres.", err.Param())
			case "max":
				errorMsg = fmt.Sprintf("Máximo %s carácteres.", err.Param())
			default:
				errorMsg = err.Error()
			}
			errors[field] = append(errors[field], errorMsg)
		}
		return errors
	}

	return map[string][]string{}
}
