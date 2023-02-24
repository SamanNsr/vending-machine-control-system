package vendingMachineUseCaseFixtures

import (
	vendingMachineConstants "github.com/samannsr/vending-machine-control-system/internal/vending_machine/constanst"
	vendingMachineDomain "github.com/samannsr/vending-machine-control-system/internal/vending_machine/domain"
	"time"
)

var (
	VM1 = vendingMachineDomain.VendingMachine{ID: 1, Status: vendingMachineConstants.StatusIdle, Inventory: &vendingMachineDomain.Inventory{Cola: 5, Coffee: 10}, Coins: 0, LastAccess: time.Now()}
	VM2 = vendingMachineDomain.VendingMachine{ID: 2, Status: vendingMachineConstants.StatusIdle, Inventory: &vendingMachineDomain.Inventory{Cola: 0, Coffee: 8}, Coins: 0, LastAccess: time.Now()}
	VM3 = vendingMachineDomain.VendingMachine{ID: 3, Status: vendingMachineConstants.StatusSelecting, Inventory: &vendingMachineDomain.Inventory{Cola: 2, Coffee: 8}, Coins: 0, LastAccess: time.Now()}
	VM4 = vendingMachineDomain.VendingMachine{ID: 4, Status: vendingMachineConstants.StatusDelivering, Inventory: &vendingMachineDomain.Inventory{Cola: 2, Coffee: 8}, Coins: 0, LastAccess: time.Now()}
)
