package responser

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func ConnectionError(err error, component string) *HTTPResponse {
	httpErr := NewHttpResponse(err, http.StatusInternalServerError)
	httpErr.details = fmt.Sprintf("El componente %s no puede establecer una conexión a la base de datos", component)
	return httpErr
}

func InvalidInputJson(err error, details string) *HTTPResponse {
	httpErr := NewHttpResponse(err, http.StatusBadRequest)
	httpErr.details = details
	return httpErr
}

func DataValidationError(err error) *HTTPResponse {
	var ve validator.ValidationErrors
	var customValidationErrorMsg string

	if errors.As(err, &ve) {
		for _, fe := range ve {
			customValidationErrorMsg = customErrorValidation(fe)
		}
	}

	validationError := errors.New("Error durante la validación de datos de entrada")
	httpErr := NewHttpResponse(validationError, http.StatusBadRequest)
	httpErr.details = customValidationErrorMsg
	return httpErr
}

func UnexpectedError(err error) *HTTPResponse {
	httpErr := NewHttpResponse(err, http.StatusInternalServerError)
	httpErr.details = "Ha ocurrido un error inesperado"
	return httpErr
}
