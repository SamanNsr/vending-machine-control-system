package httpError

import (
	"net/http"

	errorConstant "github.com/samannsr/vending-machine-control-system/pkg/error/constant"
	customError "github.com/samannsr/vending-machine-control-system/pkg/error/custom_error"
)

func ParseError(err error) HttpErr {
	customErr := customError.AsCustomError(err)
	if customErr == nil {
		internalServerErrorCode := errorConstant.ErrorLists.InternalServerError
		err =
			customError.NewInternalServerErrorWrap(err, internalServerErrorCode.Msg, internalServerErrorCode.Code, nil)
		customErr = customError.AsCustomError(err)
		return NewHttpError(http.StatusInternalServerError,
			customErr.Code(),
			errorConstant.ErrInternalServerErrorTitle,
			customErr.Error(),
			customErr.Details())
	}

	if err != nil {
		switch {
		case customError.IsValidationError(err):
			return NewHttpValidationError(customErr.Code(), customErr.Message(), customErr.Details())

		case customError.IsBadRequestError(err):
			return NewHttpBadRequestError(customErr.Code(), customErr.Message(), customErr.Details())

		case customError.IsNotFoundError(err):
			return NewHttpNotFoundError(customErr.Code(), customErr.Message(), customErr.Details())

		case customError.IsInternalServerError(err):
			return NewHttpInternalServerError(customErr.Code(), customErr.Message(), customErr.Details())

		case customError.IsCustomError(err):
			return NewHttpError(http.StatusInternalServerError, customErr.Code(),
				http.StatusText(http.StatusInternalServerError),
				customErr.Message(), customErr.Details())
		default:
			return NewHttpInternalServerError(customErr.Code(), customErr.Message(), customErr.Details())
		}
	}

	return nil
}
