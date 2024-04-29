package responser

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func unexpectedError(err error, status int, args ...map[string]interface{}) *response {
	httpErr := newResponse(err, status)
	httpErr.optionals = args
	httpErr.details = "Ha ocurrido un error inesperado"
	return httpErr
}

func CustomResponseSetter() {
	Set("UnexpectedError", unexpectedError)
	Get("UnexpectedError", errors.New("Lol"), http.StatusOK)
}

func ConnectionError(err error, component string) *response {
	httpErr := newResponse(err, http.StatusInternalServerError)
	httpErr.details = fmt.Sprintf("El componente %s no puede establecer una conexión a la base de datos", component)
	return httpErr
}

func InvalidInputJson(err error, details string) *response {
	httpErr := newResponse(err, http.StatusBadRequest)
	httpErr.details = details
	return httpErr
}

func DataValidationError(err error) *response {
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
	httpErr := newResponse(validationError, http.StatusBadRequest)
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
