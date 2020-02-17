package errors

import (
	"net/http"

	"github.com/micro/go-micro/v2/errors"
)

type Error = errors.Error

func BadRequest(id, format string, a ...interface{}) error {
	return errors.BadRequest(id, format, a...)
}

func Conflict(id, format string, a ...interface{}) error {
	return errors.Conflict(id, format, a...)
}

func Forbidden(id, format string, a ...interface{}) error {
	return errors.Forbidden(id, format, a...)
}

func InternalServerError(id, format string, a ...interface{}) error {
	return errors.InternalServerError(id, format, a...)
}

func MethodNotAllowed(id, format string, a ...interface{}) error {
	return errors.MethodNotAllowed(id, format, a...)
}

func NotFound(id, format string, a ...interface{}) error {
	return errors.NotFound(id, format, a...)
}

func NotImplemented(id string) *Error {
	code := http.StatusNotImplemented
	text := http.StatusText(code)

	return &Error{
		Id:     id,
		Code:   int32(code),
		Detail: text,
		Status: text,
	}
}

func Timeout(id, format string, a ...interface{}) error {
	return errors.Timeout(id, format, a...)
}

func Unauthorized(id, format string, a ...interface{}) error {
	return errors.Unauthorized(id, format, a...)
}
