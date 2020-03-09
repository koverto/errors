// Package errors provides common custom error types for sharing across
// microservices, wrapping github.com/micro/go-micro/v2/errors.
package errors

import (
	"net/http"

	"github.com/micro/go-micro/v2/errors"
)

// BadRequest generates a 400 error.
func BadRequest(id, format string, a ...interface{}) error {
	return errors.BadRequest(id, format, a...)
}

// Unauthorized generates a 401 error.
func Unauthorized(id, format string, a ...interface{}) error {
	return errors.Unauthorized(id, format, a...)
}

// Forbidden generates a 403 error.
func Forbidden(id, format string, a ...interface{}) error {
	return errors.Forbidden(id, format, a...)
}

// NotFound generates a 404 error.
func NotFound(id, format string, a ...interface{}) error {
	return errors.NotFound(id, format, a...)
}

// MethodNotAllowed generates a 405 error.
func MethodNotAllowed(id, format string, a ...interface{}) error {
	return errors.MethodNotAllowed(id, format, a...)
}

// Timeout generates a 408 error.
func Timeout(id, format string, a ...interface{}) error {
	return errors.Timeout(id, format, a...)
}

// Conflict generates a 409 error.
func Conflict(id, format string, a ...interface{}) error {
	return errors.Conflict(id, format, a...)
}

// InternalServerError generates a 500 error.
func InternalServerError(id, format string, a ...interface{}) error {
	return errors.InternalServerError(id, format, a...)
}

// NotImplemented generates a 501 error.
func NotImplemented(id string) *errors.Error {
	code := http.StatusNotImplemented
	text := http.StatusText(code)

	return &errors.Error{
		Id:     id,
		Code:   int32(code),
		Detail: text,
		Status: text,
	}
}
