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

	for i, v := range vm {
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

func (uc *useCase) GetVendingMachineById(ctx context.Context, dto *vendingMachineDto.InsertCoinRequestDto) (*vendingMachineDomain.VendingMachine, error) {
	vmStorage, handledErr := uc.findVendingMachineById(ctx, dto.MachineID)
	return vmStorage.vm, handledErr
}

func (uc *useCase) InsertCoin(ctx context.Context, dto *vendingMachineDto.InsertCoinRequestDto) (*vendingMachineDto.InsertCoinResponseDto, error) {
	vmStorage, handledErr := uc.findVendingMachineById(ctx, dto.MachineID)
	if handledErr != nil {
		return nil, handledErr
	}
	vmStorage.mu.Lock()
	defer vmStorage.mu.Unlock()

	// Check if the machine is idle
	if vmStorage.vm.Status != vendingMachineConstants.StatusIdle {
		return nil, vendingMachineException.VendingMachineNotIdleBadRequestExc()
	}

	if vmStorage.vm.Inventory.Coffee <= 0 || vmStorage.vm.Inventory.Cola <= 0 {
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

	// Vending machine not found
	return nil, vendingMachineException.VendingMachineNotFoundExc()
}
