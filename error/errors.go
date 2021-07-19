package error

import (
	"fmt"
)

type RuntimeError struct {
	msg string
}

func (e *RuntimeError) Error() string {
	return e.msg
}

func NewRuntimeError(msg string) error {
	return &RuntimeError{msg: msg}
}

type HttpRequestError struct {
	StatusCode int
}

func (e *HttpRequestError) Error() string {
	return fmt.Sprintf("http request error. status code: %d", e.StatusCode)
}

func NewHttpRequestError(statusCode int) error {
	return &HttpRequestError{StatusCode: statusCode}
}
