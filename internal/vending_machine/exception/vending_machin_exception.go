package articleException

import (
	errorConstant "github.com/samannsr/vending-machine-control-system/pkg/error/constant"
	customErrors "github.com/samannsr/vending-machine-control-system/pkg/error/custom_error"
)

func FindVmByIdNotFoundExc() error {
	return customErrors.NewNotFoundError(errorConstant.ErrorLists.NotFoundError.Msg, errorConstant.ErrorLists.NotFoundError.Code, nil)
}
