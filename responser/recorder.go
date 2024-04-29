package responser

import "log"

var responsesRepository = make(map[string]func(error, int, ...map[string]interface{}) *response)

func Set(name string, response func(error, int, ...map[string]interface{}) *response) {
	if _, exists := responsesRepository[name]; !exists {
		responsesRepository[name] = response
	} else {
		log.Panicf("El registro de la respuesta '%s' falló. La respuesta ya existe", name)
	}
}

func Get(name string, err error, statusCode int) *response {
	if response, exists := responsesRepository[name]; exists {
		return response(err, statusCode)
	}
	return newResponse(err, statusCode)
}

var validationTags = make(map[string]string)

func NewValidationTag(tag string, message string) {
	validationTags[tag] = message
}
