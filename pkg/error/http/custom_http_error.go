package httpError

import (
	"net/http"
	"time"

	errorConstant "github.com/samannsr/vending-machine-control-system/pkg/error/constant"
)

func NewHttpValidationError(code int, message string, details map[string]string) HttpErr {
	return &httpErr{
		Title:     errorConstant.ErrValidationFailedTitle,
		Code:      code,
		Msg:       message,
		Details:   details,
		Status:    http.StatusBadRequest,
		Timestamp: time.Now(),
	}
}

func NewHttpBadRequestError(code int, message string, details map[string]string) HttpErr {
	return &httpErr{
		Title:     errorConstant.ErrBadRequestTitle,
		Code:      code,
		Msg:       message,
		Details:   details,
		Status:    http.StatusBadRequest,
		Timestamp: time.Now(),
	}
}

func NewHttpNotFoundError(code int, message string, details map[string]string) HttpErr {
	return &httpErr{
		Title:     errorConstant.ErrNotFoundTitle,
		Code:      code,
		Msg:       message,
		Details:   details,
		Status:    http.StatusNotFound,
		Timestamp: time.Now(),
	}
}

func NewHttpInternalServerError(code int, message string, details map[string]string) HttpErr {
	return &httpErr{
		Title:     errorConstant.ErrInternalServerErrorTitle,
		Code:      code,
		Msg:       message,
		Details:   details,
		Status:    http.StatusInternalServerError,
		Timestamp: time.Now(),
	}
}
