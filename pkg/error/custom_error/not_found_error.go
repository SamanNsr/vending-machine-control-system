package customError

import (
	"errors"
)

type notFoundError struct {
	CustomError
}

func (e *notFoundError) IsNotFoundError() bool {
	return true
}

type NotFoundError interface {
	CustomError
	IsNotFoundError() bool
}

func IsNotFoundError(e error) bool {
	var notFoundError NotFoundError
	if errors.As(e, &notFoundError) {
		return notFoundError.IsNotFoundError()
	}
	return false
}

func NewNotFoundError(message string, code int, details map[string]string) error {
	return &notFoundError{
		CustomError: NewCustomError(nil, code, message, details),
	}
}

func NewNotFoundErrorWrap(err error, message string, code int, details map[string]string) error {
	return &notFoundError{
		CustomError: NewCustomError(err, code, message, details),
	}
}
