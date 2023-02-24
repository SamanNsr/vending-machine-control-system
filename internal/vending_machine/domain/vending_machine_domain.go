package vendingMachineDomain

import (
	"context"
	"github.com/labstack/echo/v4"
	vendingMachineDto "github.com/samannsr/vending-machine-control-system/internal/vending_machine/dto"
	"time"
)

type VendingMachine struct {
	ID         int `json:"id"`
	Status     string
	Inventory  *Inventory
	Coins      int
	LastAccess time.Time
}

type Inventory struct {
	Cola   int
	Coffee int
}

type Configurator interface {
	Configure(ctx context.Context) error
}

type UseCase interface {
	InsertCoin(ctx context.Context, dto *vendingMachineDto.InsertCoinRequestDto) (*vendingMachineDto.InsertCoinResponseDto, error)
	SelectProduct(ctx context.Context, dto *vendingMachineDto.SelectProductRequestDto) (*vendingMachineDto.SelectProductResponseDto, error)
	GetVendingMachineById(ctx context.Context, dto *vendingMachineDto.GetVendingMachineByIdRequestDto) (*VendingMachine, error)
}

type Repository interface {
	FindVendingMachineByID(vmID int) (*VendingMachine, error)
	UpdateVendingMachine(vm *VendingMachine) error
}

type HttpController interface {
	InsertCoin(c echo.Context) error
	SelectProduct(c echo.Context) error
}
