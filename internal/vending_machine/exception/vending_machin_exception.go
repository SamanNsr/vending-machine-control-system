package vendingMachineException

import (
	errorConstant "github.com/samannsr/vending-machine-control-system/pkg/error/constant"
	customErrors "github.com/samannsr/vending-machine-control-system/pkg/error/custom_error"
)

func VendingMachineNotFoundExc() error {
	return customErrors.NewNotFoundError(errorConstant.ErrorLists.NotFoundError.Msg, errorConstant.ErrorLists.NotFoundError.Code, nil)
}

func VendingMachineNotOkStatusBadRequestExc() error {
	return customErrors.NewBadRequestError(errorConstant.ErrorLists.BadRequestError.Msg, errorConstant.ErrorLists.BadRequestError.Code, nil)
}

func VendingMachineNoInventoryBadRequestExc() error {
	return customErrors.NewBadRequestError(errorConstant.ErrorLists.BadRequestError.Msg, errorConstant.ErrorLists.BadRequestError.Code, nil)
}

func VendingMachineInvalidProductBadRequestExc() error {
	return customErrors.NewBadRequestError(errorConstant.ErrorLists.BadRequestError.Msg, errorConstant.ErrorLists.BadRequestError.Code, nil)
}
