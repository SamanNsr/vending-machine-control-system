package customError

import (
	"errors"
)

type internalServerError struct {
	CustomError
}

func (e *internalServerError) IsInternalServerError() bool {
	return true
}

type InternalServerError interface {
	CustomError
	IsInternalServerError() bool
}

func IsInternalServerError(e error) bool {
	var internalError InternalServerError
	if errors.As(e, &internalError) {
		return internalError.IsInternalServerError()
	}
	return false
}

func NewInternalServerError(message string, code int, details map[string]string) error {
	return &internalServerError{
		CustomError: NewCustomError(nil, code, message, details),
	}
}

func NewInternalServerErrorWrap(err error, message string, code int, details map[string]string) error {
	return &internalServerError{
		CustomError: NewCustomError(err, code, message, details),
	}
}
