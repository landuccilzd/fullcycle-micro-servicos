package createtransaction

import (
	"testing"

	"github.com/landucci/ms-wallet/internal/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type TransactionGatewayMock struct {
	mock.Mock
}

func (m *TransactionGatewayMock) Create(transaction *entity.Transaction) error {
	args := m.Called(transaction)
	return args.Error(0)
}

type AccountGatewayMock struct {
	mock.Mock
}

func (m *AccountGatewayMock) FindByID(id string) (*entity.Account, error) {
	args := m.Called(id)
	return args.Get(0).(*entity.Account), args.Error(1)
}

func (m *AccountGatewayMock) Save(account *entity.Account) error {
	args := m.Called(account)
	return args.Error(0)
}
func TestCreateTransactionUseCase_Execute(t *testing.T) {
	zelda, _ := entity.NewClient("Zelda", "zelda@hyrule.com")
	accountZelda := entity.NewAccount(zelda)
	accountZelda.Credit(100000)

	link, _ := entity.NewClient("Link", "link@hyrule.com")
	accountLink := entity.NewAccount(link)
	accountLink.Credit(10000)

	accountGatewayMock := &AccountGatewayMock{}
	accountGatewayMock.On("FindByID", accountZelda.ID).Return(accountZelda, nil)
	accountGatewayMock.On("FindByID", accountLink.ID).Return(accountLink, nil)

	transactionGatewayMock := &TransactionGatewayMock{}
	transactionGatewayMock.On("Create", mock.Anything).Return(nil)

	input := CreateTransactionInputDTO{
		AccountIdFrom: accountZelda.ID,
		AccountIdTo:   accountLink.ID,
		Amount:        1000.0,
	}

	usecase := NewCreateTransactionUseCase(transactionGatewayMock, accountGatewayMock)
	output, err := usecase.Execute(input)

	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.NotEmpty(t, output.ID)

	accountGatewayMock.AssertExpectations(t)
	accountGatewayMock.AssertNumberOfCalls(t, "FindByID", 2)
	transactionGatewayMock.AssertExpectations(t)
	transactionGatewayMock.AssertNumberOfCalls(t, "Create", 1)

}
