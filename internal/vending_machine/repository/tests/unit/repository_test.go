package unit

import (
	vendingMachineDomain "github.com/samannsr/vending-machine-control-system/internal/vending_machine/domain"
	vendingMachineRepository "github.com/samannsr/vending-machine-control-system/internal/vending_machine/repository"
	vendingMachineRepositoryFixtures "github.com/samannsr/vending-machine-control-system/internal/vending_machine/repository/tests/fixtures"
	customError "github.com/samannsr/vending-machine-control-system/pkg/error/custom_error"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type testSuite struct {
	suite.Suite
	// repo vendingMachineDomain.Repository
}

func (suite *testSuite) SetupSuite() {
	// TODO: place the repository creation
	// suite.repo = vendingMachineRepo.NewVendingMachineRepository(&vendingMachineDomain.Inventory{Cola: 10, Coffee: 10})
}

func (suite *testSuite) TestFindVendingMachineByIDSuccess() {
	// Create new vending machines
	var vms = make([]vendingMachineDomain.VendingMachine, 0)
	vms = append(vms, vendingMachineRepositoryFixtures.VM1, vendingMachineRepositoryFixtures.VM2)
	repo := vendingMachineRepository.NewRepository(vms)

	result, err := repo.FindVendingMachineByID(1)
	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), result)
	assert.Equal(suite.T(), vendingMachineRepositoryFixtures.VM1, result)
}

func (suite *testSuite) TestFindVendingMachineByIDNotFound() {
	// Create new vending machines
	var vms = make([]vendingMachineDomain.VendingMachine, 0)
	vms = append(vms, vendingMachineRepositoryFixtures.VM1, vendingMachineRepositoryFixtures.VM2)
	repo := vendingMachineRepository.NewRepository(vms)

	result, err := repo.FindVendingMachineByID(3)
	assert.Equal(suite.T(), nil, result)
	assert.Error(suite.T(), err)
	assert.True(suite.T(), customError.IsNotFoundError(err))
}
