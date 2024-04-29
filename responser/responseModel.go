package responser

import (
	"fmt"
)

type ResponseCreator interface {
	Error() string
	GetStatusCode() int
	GetResponseMessage() string
}

type response struct {
	cause      error
	details    string
	statusCode int
	location   string
	optionals  []map[string]interface{}
}

func newResponse(e error, httpStatusCode int) *response {
	return &response{
		cause:      e,
		statusCode: httpStatusCode,
		location:   runtimeToString(),
	}
}

func (err response) Error() string {
	if err.cause == nil {
		return err.details
	}
	return fmt.Sprintf("%s %s", err.location, err.cause.Error())
}

func (err response) GetStatusCode() int {
	return err.statusCode
}

func (err response) GetResponseMessage() string {
	if err.details == "" {
		return err.cause.Error()
	}

	if err.cause == nil {
		return err.details
	}
	return fmt.Sprintf("%s, %s", err.details, err.cause)
}
