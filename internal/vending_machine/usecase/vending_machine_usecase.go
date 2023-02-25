package vendingMachineUsecase

import (
	"context"
	vendingMachineConstants "github.com/samannsr/vending-machine-control-system/internal/vending_machine/constanst"
	vendingMachineDomain "github.com/samannsr/vending-machine-control-system/internal/vending_machine/domain"
	vendingMachineDto "github.com/samannsr/vending-machine-control-system/internal/vending_machine/dto"
	vendingMachineException "github.com/samannsr/vending-machine-control-system/internal/vending_machine/exception"
	"sync"
)

type InMemoryVendingMachineStorage struct {
	vm *vendingMachineDomain.VendingMachine
	mu sync.RWMutex
}

type useCase struct {
	storage []*InMemoryVendingMachineStorage
}

func NewUseCase(vm []vendingMachineDomain.VendingMachine) vendingMachineDomain.UseCase {
	u := &useCase{
		storage: make([]*InMemoryVendingMachineStorage, len(vm)),
	}

	for i := range vm {
		v := vm[i]
		u.storage[i] = &InMemoryVendingMachineStorage{
			vm: &v,
			mu: sync.RWMutex{},
		}
	}

	return u
}

func (uc *useCase) findVendingMachineById(ctx context.Context, id int) (*InMemoryVendingMachineStorage, error) {
	for _, vmStorage := range uc.storage {
		if vmStorage.vm.ID == id {
			return vmStorage, nil
		}
	}
	return nil, vendingMachineException.VendingMachineNotFoundExc()
}

func (uc *useCase) GetVendingMachineById(ctx context.Context, dto *vendingMachineDto.GetVendingMachineByIdRequestDto) (*vendingMachineDomain.VendingMachine, error) {
	vmStorage, handledErr := uc.findVendingMachineById(ctx, dto.MachineID)
	if handledErr != nil {
		return nil, handledErr
	}
	return vmStorage.vm, handledErr
}

func (uc *useCase) InsertCoin(ctx context.Context, dto *vendingMachineDto.InsertCoinRequestDto) (*vendingMachineDto.InsertCoinResponseDto, error) {
	vmStorage, handledErr := uc.findVendingMachineById(ctx, dto.MachineID)
	if handledErr != nil {
		return nil, handledErr
	}

	// Check if the machine is idle
	if vmStorage.vm.Status != vendingMachineConstants.StatusIdle {
		return nil, vendingMachineException.VendingMachineNotOkStatusBadRequestExc()
	}

	vmStorage.mu.Lock()
	defer vmStorage.mu.Unlock()

	// Check if the machine is idle
	if vmStorage.vm.Status != vendingMachineConstants.StatusIdle {
		return nil, vendingMachineException.VendingMachineNotOkStatusBadRequestExc()
	}

	if vmStorage.vm.Inventory.Coffee <= 0 && vmStorage.vm.Inventory.Cola <= 0 {
		return nil, vendingMachineException.VendingMachineNoInventoryBadRequestExc()
	}
	// Add the coin to the vending machine
	vmStorage.vm.Coins += 1

	// Update the vending machine status
	vmStorage.vm.Status = vendingMachineConstants.StatusSelecting

	// Return the response DTO
	return &vendingMachineDto.InsertCoinResponseDto{
		Message:   "Coin inserted successfully",
		MachineID: dto.MachineID,
		Coffee:    vmStorage.vm.Inventory.Coffee,
		Cola:      vmStorage.vm.Inventory.Cola,
		Status:    vmStorage.vm.Status,
	}, nil
}

func (uc *useCase) SelectProduct(ctx context.Context, dto *vendingMachineDto.SelectProductRequestDto) (*vendingMachineDto.SelectProductResponseDto, error) {
	// Find the vending machine by ID
	storageVm, handledErr := uc.findVendingMachineById(ctx, dto.MachineID)
	if handledErr != nil {
		return nil, handledErr
	}

	// Acquire write lock on the vending machine to update its state
	storageVm.mu.Lock()
	defer storageVm.mu.Unlock()

	// Check that the vending machine is in the Idle state
	if storageVm.vm.Status != vendingMachineConstants.StatusSelecting {
		return nil, vendingMachineException.VendingMachineNotOkStatusBadRequestExc()
	}

	// Determine the selected product and its price
	switch dto.Product {
	case vendingMachineConstants.ColaType:
		if storageVm.vm.Inventory.Cola < 1 {
			return nil, vendingMachineException.VendingMachineNoInventoryBadRequestExc()
		}
		storageVm.vm.Inventory.Cola--
		storageVm.vm.Status = vendingMachineConstants.StatusIdle
		return &vendingMachineDto.SelectProductResponseDto{
			MachineID: dto.MachineID,
			Status:    storageVm.vm.Status,
			Message:   "Enjoy your colaaa :))))))",
		}, nil
	case vendingMachineConstants.CoffeeType:
		if storageVm.vm.Inventory.Coffee < 1 {
			return nil, vendingMachineException.VendingMachineNoInventoryBadRequestExc()
		}
		storageVm.vm.Inventory.Coffee--
		storageVm.vm.Status = vendingMachineConstants.StatusIdle
		return &vendingMachineDto.SelectProductResponseDto{
			MachineID: dto.MachineID,
			Status:    storageVm.vm.Status,
			Message:   "Enjoy your coffeee :))))))",
		}, nil
	default:
		return nil, vendingMachineException.VendingMachineInvalidProductBadRequestExc()
	}

}
