package vendingMachineDto

import (
	validator "github.com/go-ozzo/ozzo-validation"
	vendingMachineConstants "github.com/samannsr/vending-machine-control-system/internal/vending_machine/constanst"
)

type InsertCoinRequestDto struct {
	MachineID int `json:"machine_id"`
}

func (icDto *InsertCoinRequestDto) ValidateInsertCoinRequestDto() error {
	return validator.ValidateStruct(icDto,
		validator.Field(
			&icDto.MachineID,
			validator.Required,
			validator.Min(1),
		),
	)
}

type InsertCoinResponseDto struct {
	MachineID int    `json:"id"`
	Message   string `json:"message"`
	Status    string `json:"status"`
	Cola      int    `json:"cola"`
	Coffee    int    `json:"coffee"`
}

type SelectProductRequestDto struct {
	Product string `json:"product"`
}

func (icDto *SelectProductRequestDto) ValidateSelectProductRequestDto() error {
	return validator.ValidateStruct(icDto,
		validator.Field(
			&icDto.Product,
			validator.Required,
			validator.In(vendingMachineConstants.ColaType, vendingMachineConstants.CoffeeType),
		),
	)
}

type SelectProductResponseDto struct {
	MachineID int    `json:"id"`
	Message   string `json:"message"`
	Status    string `json:"status"`
}

type GetVendingMachineByIdRequestDto struct {
	MachineID int `json:"id"`
}

func (icDto *InsertCoinRequestDto) ValidateGetVendingMachineByIdRequestDto() error {
	return validator.ValidateStruct(icDto,
		validator.Field(
			&icDto.MachineID,
			validator.Required,
			validator.In(vendingMachineConstants.ColaType, vendingMachineConstants.CoffeeType),
		),
	)
}
