package vendingMachineDto

import (
	validator "github.com/go-ozzo/ozzo-validation"
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
