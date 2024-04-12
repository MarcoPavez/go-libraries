package responser

import (
	"errors"
	"fmt"
	"runtime"

	"github.com/go-playground/validator/v10"
)

func ConnectionError(e error, component string) *HTTPResponse {
	httpErr := HTTPErrorInstance(e, StatusCodeInternalServerError, InternalStatusCodeConnectionError)
	httpErr.details = fmt.Sprintf("El componente %s no puede establecer una conexión a la base de datos", component)
	return httpErr
}

func InvalidJSONError(e error, component, example string) *HTTPResponse {
	httpErr := HTTPErrorInstance(e, StatusCodeBadRequestError, InternalStatusCodeBadInputData)
	httpErr.details = fmt.Sprintf("La información ingresada no corresponde al valor esperado en %s. Se esperaba %s", component, example)
	return httpErr
}

func BadLinkToFileError(e error) *HTTPResponse {
	httpErr := HTTPErrorInstance(e, StatusCodeBadRequestError, InternalStatusCodeBadInputData)
	httpErr.details = "No se puede descargar el documento. Verifica que el enlace sea válido"
	return httpErr
}

func BadCustomerCreationError(e error) *HTTPResponse {
	httpErr := HTTPErrorInstance(e, StatusCodeInternalServerError, InternalStatusCodeBadInputData)
	httpErr.details = "Error en la creación del cliente"
	return httpErr
}

func DataValidationError(e error) *HTTPResponse {
	var ve validator.ValidationErrors
	var customValidationErrorMsg string

	if errors.As(e, &ve) {
		for _, fe := range ve {
			customValidationErrorMsg = customErrorValidation(fe)
		}
	}

	validationError := errors.New("Error durante la validación de datos de entrada")
	httpErr := HTTPErrorInstance(validationError, StatusCodeBadRequestError, InternalStatusCodeBadInputData)
	httpErr.details = customValidationErrorMsg
	return httpErr
}

func UnexpectedError(e error) *HTTPResponse {
	httpErr := HTTPErrorInstance(e, StatusCodeInternalServerError, InternalStatusCodeServerError)
	httpErr.details = "Ha ocurrido un error inesperado"
	return httpErr
}

func CustomerNotFoundError() *HTTPResponse {
	httpErr := HTTPErrorInstance(nil, StatusCodeNotFoundError, InternalStatusCodeNotFound)
	httpErr.details = "El RUT ingresado no se encuentra en la base de datos"
	return httpErr
}

func CustomerFounded(e error, details string) *HTTPResponse {
	httpErr := HTTPErrorInstance(e, StatusCodeBadRequestError, InternalStatusCodeBadInputData)
	httpErr.details = details
	return httpErr
}

func CustomerNotAuthorizedByCompliance(e error, details string) *HTTPResponse {
	httpErr := HTTPErrorInstance(e, StatusCodeBadRequestError, InternalStatusCodeCustomerNotAuthorized)
	httpErr.details = details
	return httpErr
}

func ComplianceMustVerify(e error, details string) *HTTPResponse {
	httpErr := HTTPErrorInstance(e, StatusCodeBadRequestError, InternalStatusCodeComplianceMustCheck)
	httpErr.details = details
	return httpErr
}

// Funcion auxiliar para obtener el archivo y linea donde fue creado un error
func runtimeToString() string {
	_, file, line, ok := runtime.Caller(2)
	if ok {
		return fmt.Sprintf("Archivo %v, en línea %v.", file, line)
	}
	return "Ubicacion desconocida."
}

// Errores de validación de datos
func customErrorValidation(fe validator.FieldError) string {
	field := fe.Field()
	tag := fe.Tag()
	switch {
	case tag == "required":
		return fmt.Sprintf("El campo '%s' no puede ser null o estar vacío", field)
	case tag == "required_if":
		return fmt.Sprintf("El campo '%s' no puede ser null o estar vacío según el valor ingresado en el campo 'CustomerType'", field)
	case tag == "alpha":
		return fmt.Sprintf("El campo '%s' sólo acepta letras", field)
	case tag == "number":
		return fmt.Sprintf("El campo '%s' sólo acepta números", field)
	case (tag == "gte" || tag == "lte" || tag == "oneof"):
		return fmt.Sprintf("El campo '%s' no acepta el valor ingresado. Para conocer dichos valores, consulta el diccionario de datos", field)
	case (tag == "datetime"):
		return fmt.Sprintf("La fecha ingresada en el campo '%s' no es válida. El formato debe ser dd/mm/aaaa", field)
	case (tag == "email"):
		return fmt.Sprintf("El email ingresado en el campo '%s' no es válido", field)
	default:
		return fe.Error()
	}
}
