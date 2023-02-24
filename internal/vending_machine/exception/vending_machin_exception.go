package vendingMachineException

import (
	errorConstant "github.com/samannsr/vending-machine-control-system/pkg/error/constant"
	customErrors "github.com/samannsr/vending-machine-control-system/pkg/error/custom_error"
	errorUtils "github.com/samannsr/vending-machine-control-system/pkg/error/error_utils"
)

func VendingMachineNotFoundExc() error {
	return customErrors.NewNotFoundError("vending machine not found", errorConstant.ErrorLists.NotFoundError.Code, nil)
}

func VendingMachineNotOkStatusBadRequestExc() error {
	return customErrors.NewBadRequestError("vending machine status conflict", errorConstant.ErrorLists.BadRequestError.Code, nil)
}

func VendingMachineNoInventoryBadRequestExc() error {
	return customErrors.NewBadRequestError("", errorConstant.ErrorLists.BadRequestError.Code, nil)
}

func VendingMachineInvalidProductBadRequestExc() error {
	return customErrors.NewBadRequestError("invalid product type", errorConstant.ErrorLists.BadRequestError.Code, nil)
}

func GetVendingMachineInvalidIdExc() error {
	return customErrors.NewBadRequestError("invalid id", errorConstant.ErrorLists.BadRequestError.Code, nil)
}

func InsertCoinBindingExc() error {
	return customErrors.NewBadRequestError(errorConstant.ErrorLists.BadRequestError.Msg, errorConstant.ErrorLists.BadRequestError.Code, nil)
}

func InsertCoinValidationExc(err error) error {
	ve, ie := errorUtils.ValidationErrorHandler(err)
	if ie != nil {
		return ie
	}

	validationErrorCode := errorConstant.ErrorLists.ValidationError
	return customErrors.NewValidationError(validationErrorCode.Msg, validationErrorCode.Code, ve)
}

func SelectProductBindingExc() error {
	return customErrors.NewBadRequestError(errorConstant.ErrorLists.BadRequestError.Msg, errorConstant.ErrorLists.BadRequestError.Code, nil)
}

func SelectProductValidationExc(err error) error {
	ve, ie := errorUtils.ValidationErrorHandler(err)
	if ie != nil {
		return ie
	}

	validationErrorCode := errorConstant.ErrorLists.ValidationError
	return customErrors.NewValidationError(validationErrorCode.Msg, validationErrorCode.Code, ve)
}
