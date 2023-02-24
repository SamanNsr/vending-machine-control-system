package vendingMachineConfigurator

import (
	"context"
	vendingMachineConstants "github.com/samannsr/vending-machine-control-system/internal/vending_machine/constanst"
	vendingMachineHttpController "github.com/samannsr/vending-machine-control-system/internal/vending_machine/delivery/http"
	vendingMachineDomain "github.com/samannsr/vending-machine-control-system/internal/vending_machine/domain"
	vendingMachineUseCase "github.com/samannsr/vending-machine-control-system/internal/vending_machine/usecase"
	"github.com/samannsr/vending-machine-control-system/pkg/infrastructure"
)

type configurator struct {
	ic *infrastructure.IContainer
}

func NewConfigurator(ic *infrastructure.IContainer) vendingMachineDomain.Configurator {
	return &configurator{ic: ic}
}

func (c *configurator) Configure(ctx context.Context) error {
	useCase := vendingMachineUseCase.NewUseCase(DefaultVMs)

	// http
	httpRouterGp := c.ic.EchoHttpServer.GetEchoInstance().Group(c.ic.EchoHttpServer.GetBasePath())
	httpController := vendingMachineHttpController.NewController(useCase)
	vendingMachineHttpController.NewRouter(httpController).Register(httpRouterGp)

	return nil
}

var DefaultVMs = []vendingMachineDomain.VendingMachine{
	{
		ID: 1,
		Inventory: &vendingMachineDomain.Inventory{
			Cola:   10,
			Coffee: 20,
		},
		Status: vendingMachineConstants.StatusIdle,
		Coins:  0,
	},
	{
		ID: 2,
		Inventory: &vendingMachineDomain.Inventory{
			Cola:   5,
			Coffee: 15,
		},
		Status: vendingMachineConstants.StatusIdle,
		Coins:  0,
	},
	{
		ID: 2,
		Inventory: &vendingMachineDomain.Inventory{
			Cola:   3,
			Coffee: 1,
		},
		Status: vendingMachineConstants.StatusIdle,
		Coins:  0,
	},
}
