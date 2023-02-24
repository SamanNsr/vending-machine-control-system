package usecaseUnitTest

import (
	"context"
	vendingMachineDomain "github.com/samannsr/vending-machine-control-system/internal/vending_machine/domain"
	vendingMachineDto "github.com/samannsr/vending-machine-control-system/internal/vending_machine/dto"
	vendingMachineUseCase "github.com/samannsr/vending-machine-control-system/internal/vending_machine/usecase"
	vendingMachineUseCaseFixtures "github.com/samannsr/vending-machine-control-system/internal/vending_machine/usecase/tests/fixtures"
	customError "github.com/samannsr/vending-machine-control-system/pkg/error/custom_error"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type testSuite struct {
	suite.Suite
}

func TestRunSuite(t *testing.T) {
	tSuite := new(testSuite)
	suite.Run(t, tSuite)
}

func (suite *testSuite) TestGetVendingMachineByIdSuccess() {
	ctx := context.Background()
	var vms = make([]vendingMachineDomain.VendingMachine, 0)
	vms = append(vms, vendingMachineUseCaseFixtures.VM1, vendingMachineUseCaseFixtures.VM2)
	uc := vendingMachineUseCase.NewUseCase(vms)

	result, err := uc.GetVendingMachineById(ctx, &vendingMachineDto.InsertCoinRequestDto{
		MachineID: vendingMachineUseCaseFixtures.VM1.ID,
	})
	assert.NoError(suite.T(), err)
	if assert.NotNil(suite.T(), result) {
		assert.Equal(suite.T(), result.ID, vendingMachineUseCaseFixtures.VM1.ID)
		assert.Equal(suite.T(), result.Inventory.Coffee, vendingMachineUseCaseFixtures.VM1.Inventory.Coffee)
	}
}

func (suite *testSuite) TestGetVendingMachineByIdFailedMachineNotFound() {
	// Create new vending machines
	ctx := context.Background()
	var vms = make([]vendingMachineDomain.VendingMachine, 0)
	vms = append(vms, vendingMachineUseCaseFixtures.VM1, vendingMachineUseCaseFixtures.VM2)
	uc := vendingMachineUseCase.NewUseCase(vms)

	_, err := uc.GetVendingMachineById(ctx, &vendingMachineDto.InsertCoinRequestDto{
		MachineID: vendingMachineUseCaseFixtures.VM3.ID,
	})

	assert.Error(suite.T(), err)
	assert.True(suite.T(), customError.IsNotFoundError(err))
}

func (suite *testSuite) TestInsertCoinSuccess() {
	// Create new vending machines
	ctx := context.Background()
	var vms = make([]vendingMachineDomain.VendingMachine, 0)
	vms = append(vms, vendingMachineUseCaseFixtures.VM1, vendingMachineUseCaseFixtures.VM2)
	uc := vendingMachineUseCase.NewUseCase(vms)

	result, err := uc.InsertCoin(ctx, &vendingMachineDto.InsertCoinRequestDto{
		MachineID: vendingMachineUseCaseFixtures.VM1.ID,
	})
	assert.NoError(suite.T(), err)
	if assert.NotNil(suite.T(), result.Message) {
		assert.Equal(suite.T(), result.MachineID, vendingMachineUseCaseFixtures.VM1.ID)
		assert.Equal(suite.T(), result.Coffee, vendingMachineUseCaseFixtures.VM1.Inventory.Coffee)
		assert.Equal(suite.T(), result.Cola, vendingMachineUseCaseFixtures.VM1.Inventory.Cola)
	}
}

func (suite *testSuite) TestInsertCoinFailedStatusNotIdle() {
	// Create new vending machines
	ctx := context.Background()
	var vms = make([]vendingMachineDomain.VendingMachine, 0)
	vms = append(vms, vendingMachineUseCaseFixtures.VM1, vendingMachineUseCaseFixtures.VM2, vendingMachineUseCaseFixtures.VM3)
	uc := vendingMachineUseCase.NewUseCase(vms)

	_, err := uc.InsertCoin(ctx, &vendingMachineDto.InsertCoinRequestDto{
		MachineID: vendingMachineUseCaseFixtures.VM3.ID,
	})

	assert.True(suite.T(), customError.IsBadRequestError(err))
	assert.Error(suite.T(), err)
}

func (suite *testSuite) TestInsertCoinFailedMachineNotFound() {
	// Create new vending machines
	ctx := context.Background()
	var vms = make([]vendingMachineDomain.VendingMachine, 0)
	vms = append(vms, vendingMachineUseCaseFixtures.VM1, vendingMachineUseCaseFixtures.VM2)
	uc := vendingMachineUseCase.NewUseCase(vms)

	_, err := uc.InsertCoin(ctx, &vendingMachineDto.InsertCoinRequestDto{
		MachineID: vendingMachineUseCaseFixtures.VM3.ID,
	})

	assert.Error(suite.T(), err)
	assert.True(suite.T(), customError.IsNotFoundError(err))
}
