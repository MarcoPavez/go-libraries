package responser

import "fmt"

type CustomResponseCreator interface {
	// Para implementar la interfaz de error
	Error() string
	// Getter del codigo de estatus http
	GetStatusCode() int
	// Getter del codigo interno definido por la empresa
	GetInternalCode() string
	// getter del mensaje personalizado de la respuesta
	GetResponseMessage() string
}

// Estructura que implementa mensajes personalizados para responder
// las request HTTP cuando ocurre un error.
type HTTPResponse struct {
	// Error arrojado por el proceso del endpoint
	cause error
	// Mensaje custom relacionado al error
	details string
	// Codigo de estado de la respuesta
	statusCode int
	// Codigo interno de la respuesta
	internalCode string
	// Indica el archivo y línea de código desde la cual se invoca a la funcion de error
	location string
}

func HTTPErrorInstance(e error, httpStatusCode int, tannerInternalCode string) *HTTPResponse {
	return &HTTPResponse{
		cause:        e,
		statusCode:   httpStatusCode,
		internalCode: tannerInternalCode,
		location:     runtimeToString(),
	}
}

func (err HTTPResponse) Error() string {
	if err.cause == nil {
		return err.details
	}
	return fmt.Sprintf("%s %s", err.location, err.cause.Error())
}

func (err HTTPResponse) GetStatusCode() int {
	return err.statusCode
}

func (err HTTPResponse) GetInternalCode() string {
	return err.internalCode
}

func (err HTTPResponse) GetResponseMessage() string {
	if err.details == "" {
		return err.cause.Error()
	}

	if err.cause == nil {
		return err.details
	}
	return fmt.Sprintf("%s, %s", err.details, err.cause)
}
