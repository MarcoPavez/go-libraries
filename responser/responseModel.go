package responser

import (
	"fmt"
)

type CustomResponseCreator interface {
	Error() string
	GetStatusCode() int
	GetResponseMessage() string
}

type HTTPResponse struct {
	cause      error
	details    string
	statusCode int
	location   string
}

func NewHttpResponse(e error, httpStatusCode int) *HTTPResponse {
	return &HTTPResponse{
		cause:      e,
		statusCode: httpStatusCode,
		location:   runtimeToString(),
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

func (err HTTPResponse) GetResponseMessage() string {
	if err.details == "" {
		return err.cause.Error()
	}

	if err.cause == nil {
		return err.details
	}
	return fmt.Sprintf("%s, %s", err.details, err.cause)
}
