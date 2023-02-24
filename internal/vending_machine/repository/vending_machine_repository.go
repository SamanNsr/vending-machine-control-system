package vendingMachineRepository

import (
	vendingMachineDomain "github.com/samannsr/vending-machine-control-system/internal/vending_machine/domain"
	articleException "github.com/samannsr/vending-machine-control-system/internal/vending_machine/exception"
	"sync"
)

type InMemoryVendingMachineStorage struct {
	vm *vendingMachineDomain.VendingMachine
	mu sync.RWMutex
}

type repository struct {
	storage []*InMemoryVendingMachineStorage
}

func NewRepository(vm []vendingMachineDomain.VendingMachine) vendingMachineDomain.Repository {
	r := &repository{
		storage: make([]*InMemoryVendingMachineStorage, len(vm)),
	}

	for i, v := range vm {
		r.storage[i] = &InMemoryVendingMachineStorage{
			vm: &v,
			mu: sync.RWMutex{},
		}
	}

	return r
}

func (r *repository) FindVendingMachineByID(id int) (*vendingMachineDomain.VendingMachine, error) {
	for _, s := range r.storage {
		if s.vm.ID == id {
			vmCopy := *s.vm
			return &vmCopy, nil
		}
	}
	return nil, articleException.FindVmByIdNotFoundExc()
}

func (r *repository) UpdateVendingMachine(vm *vendingMachineDomain.VendingMachine) error {
	return nil
}
