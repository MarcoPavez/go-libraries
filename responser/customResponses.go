package responser

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
)

func CustomResponseSetter([]func(name string, response func(error, int, ...map[string]interface{}) *response)) {
	Set("UnexpectedError", unexpectedError)
	Set("ConnectionError", connectionError)
	Set("DataValidationError", dataValidationError)
}

func unexpectedError(err error, status int, args ...map[string]interface{}) *response {
	httpErr := newResponse(err, status)
	httpErr.details = "Ha ocurrido un error inesperado"
	httpErr.optionals = args
	return httpErr
}
func connectionError(err error, status int, args ...map[string]interface{}) *response {
	httpErr := newResponse(err, status)
	httpErr.details = "El componente %s no puede establecer una conexión a la base de datos"
	return httpErr
}

func dataValidationError(err error, status int, args ...map[string]interface{}) *response {
	var ve validator.ValidationErrors
	var customMessage string

	if errors.As(err, &ve) {
		for _, fe := range ve {
			customMessage = customMessageValidationError(fe)
		}
	} else {
		customMessage = "Error en los datos ingresados"
	}

	validationError := errors.New("Error durante la validación de datos de entrada")
	httpErr := newResponse(validationError, status)
	httpErr.details = customMessage
	return httpErr
}

func customMessageValidationError(err validator.FieldError) string {

	fieldName := err.Field()
	tag := err.Tag()

	if message, exists := validationTags[tag]; exists {
		return fmt.Sprintf(message, fieldName)
	}
	return fmt.Sprintf("La regla de validación '%s' no existe", tag)
}
