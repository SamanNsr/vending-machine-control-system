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
	MachineID int    `json:"id"`
	Product   string `json:"product"`
}

func (spDto *SelectProductRequestDto) ValidateSelectProductRequestDto() error {
	return validator.ValidateStruct(spDto,
		validator.Field(
			&spDto.MachineID,
			validator.Required,
			validator.Min(1),
		),
		validator.Field(
			&spDto.Product,
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

func (gmDto *GetVendingMachineByIdRequestDto) ValidateGetVendingMachineByIdRequestDto() error {
	return validator.ValidateStruct(gmDto,
		validator.Field(
			&gmDto.MachineID,
			validator.Required,
			validator.In(vendingMachineConstants.ColaType, vendingMachineConstants.CoffeeType),
		),
	)
}
