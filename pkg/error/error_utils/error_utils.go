package errorUtils

import (
	"errors"
	validator "github.com/go-ozzo/ozzo-validation"
	customError "github.com/samannsr/vending-machine-control-system/pkg/error/custom_error"
)

func ValidationErrorHandler(err error) (map[string]string, error) {
	var customErr validator.Errors
	if errors.As(err, &customErr) {
		details := make(map[string]string)
		for k, v := range customErr {
			details[k] = v.Error()
		}
		return details, nil
	}
	// TODO : get internal error from constant.
	return nil, customError.NewInternalServerErrorWrap(err, "internal error", 8585, nil)
}
