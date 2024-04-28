package responser

import (
	"fmt"
	"runtime"

	"github.com/go-playground/validator/v10"
)

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

func runtimeToString() string {
	_, file, line, ok := runtime.Caller(2)
	if ok {
		return fmt.Sprintf("Archivo %v, en línea %v.", file, line)
	}
	return "Ubicación desconocida"
}
