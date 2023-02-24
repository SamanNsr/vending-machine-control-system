package vendingMachineHttpController

import (
	vendingMachineDomain "github.com/samannsr/vending-machine-control-system/internal/vending_machine/domain"
	vendingMachineDto "github.com/samannsr/vending-machine-control-system/internal/vending_machine/dto"
	vendingMachineException "github.com/samannsr/vending-machine-control-system/internal/vending_machine/exception"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type controller struct {
	useCase vendingMachineDomain.UseCase
}

func NewController(uc vendingMachineDomain.UseCase) vendingMachineDomain.HttpController {
	return &controller{
		useCase: uc,
	}
}

func (c controller) GetVendingMachineById(ctx echo.Context) error {
	id := ctx.Param("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		return vendingMachineException.GetVendingMachineInvalidIdExc()
	}

	vDto := &vendingMachineDto.GetVendingMachineByIdRequestDto{MachineID: intId}

	res, err := c.useCase.GetVendingMachineById(ctx.Request().Context(), vDto)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, res)
}

func (c controller) InsertCoin(ctx echo.Context) error {
	iDto := new(vendingMachineDto.InsertCoinRequestDto)
	if err := ctx.Bind(iDto); err != nil {
		return vendingMachineException.InsertCoinBindingExc()
	}

	if err := iDto.ValidateInsertCoinRequestDto(); err != nil {
		return vendingMachineException.InsertCoinValidationExc(err)
	}

	res, err := c.useCase.InsertCoin(ctx.Request().Context(), iDto)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, res)
}

func (c controller) SelectProduct(ctx echo.Context) error {
	sDto := new(vendingMachineDto.SelectProductRequestDto)
	if err := ctx.Bind(sDto); err != nil {
		return vendingMachineException.SelectProductBindingExc()
	}

	if err := sDto.ValidateSelectProductRequestDto(); err != nil {
		return vendingMachineException.SelectProductValidationExc(err)
	}

	res, err := c.useCase.SelectProduct(ctx.Request().Context(), sDto)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, res)
}
