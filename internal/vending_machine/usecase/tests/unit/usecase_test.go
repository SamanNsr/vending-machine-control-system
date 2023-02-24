package usecaseUnitTest

import (
	"context"
	vendingMachineDomain "github.com/samannsr/vending-machine-control-system/internal/vending_machine/domain"
	vendingMachineDto "github.com/samannsr/vending-machine-control-system/internal/vending_machine/dto"
	vendingMachineUseCase "github.com/samannsr/vending-machine-control-system/internal/vending_machine/usecase"
	"github.com/samannsr/vending-machine-control-system/internal/vending_machine/usecase/tests/fixtures"
	customError "github.com/samannsr/vending-machine-control-system/pkg/error/custom_error"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type testSuite struct {
	suite.Suite
}

func (suite *testSuite) TestInsertCoinSuccess() {
	// Create new vending machines
	ctx := context.Background()
	var vms = make([]vendingMachineDomain.VendingMachine, 0)
	vms = append(vms, vendingMachineRepositoryFixtures.VM1, vendingMachineRepositoryFixtures.VM2)
	uc := vendingMachineUseCase.NewUseCase(vms)

	result, err := uc.InsertCoin(ctx, &vendingMachineDto.InsertCoinRequestDto{
		MachineID: vendingMachineRepositoryFixtures.VM1.ID,
	})
	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), result.Message)
	assert.Equal(suite.T(), result.MachineID, vendingMachineRepositoryFixtures.VM1.ID)
	assert.Equal(suite.T(), result.Coffee, vendingMachineRepositoryFixtures.VM1.Inventory.Coffee)
	assert.Equal(suite.T(), result.MachineID, vendingMachineRepositoryFixtures.VM1.Inventory.Coffee)
}

func (suite *testSuite) TestInsertCoinFailedStatusNotIdle() {
	// Create new vending machines
	ctx := context.Background()
	var vms = make([]vendingMachineDomain.VendingMachine, 0)
	vms = append(vms, vendingMachineRepositoryFixtures.VM1, vendingMachineRepositoryFixtures.VM2, vendingMachineRepositoryFixtures.VM3)
	uc := vendingMachineUseCase.NewUseCase(vms)

	result, err := uc.InsertCoin(ctx, &vendingMachineDto.InsertCoinRequestDto{
		MachineID: vendingMachineRepositoryFixtures.VM3.ID,
	})

	assert.Equal(suite.T(), nil, result)
	assert.Error(suite.T(), err)
	assert.True(suite.T(), customError.IsBadRequestError(err))
}

func (suite *testSuite) TestInsertCoinFailedMachineNotFound() {
	// Create new vending machines
	ctx := context.Background()
	var vms = make([]vendingMachineDomain.VendingMachine, 0)
	vms = append(vms, vendingMachineRepositoryFixtures.VM1, vendingMachineRepositoryFixtures.VM2)
	uc := vendingMachineUseCase.NewUseCase(vms)

	result, err := uc.InsertCoin(ctx, &vendingMachineDto.InsertCoinRequestDto{
		MachineID: vendingMachineRepositoryFixtures.VM3.ID,
	})
	assert.Equal(suite.T(), nil, result)
	assert.Error(suite.T(), err)
	assert.True(suite.T(), customError.IsNotFoundError(err))
}

func TestRunSuite(t *testing.T) {
	tSuite := new(testSuite)
	suite.Run(t, tSuite)
}
