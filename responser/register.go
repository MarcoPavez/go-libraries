package responser

var ErrorFactoryMap = map[string]func(error, int) *HTTPResponse{}

func NewRegister(key string, constructor func(error, int) *HTTPResponse) {
	if _, exists := ErrorFactoryMap[key]; !exists {
		ErrorFactoryMap[key] = constructor
	}
}

func Create(key string, err error, statusCode int) *HTTPResponse {
	if constructor, exists := ErrorFactoryMap[key]; exists {
		return constructor(err, statusCode)
	}
	return NewHttpResponse(err, statusCode)
}
