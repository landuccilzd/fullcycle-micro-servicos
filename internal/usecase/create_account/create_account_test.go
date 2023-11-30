package createaccount

import (
	"testing"

	"github.com/landucci/ms-wallet/internal/entity"
	"github.com/landucci/ms-wallet/internal/usecase/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateAccountUseCase_Execute(t *testing.T) {
	client, _ := entity.NewClient("Princesa Zelda", "zelda@hyrule.com")
	clientMock := &mocks.ClientGatewayMock{}
	clientMock.On("Get", mock.Anything).Return(client, nil)

	accountMock := &mocks.AccountGatewayMock{}
	accountMock.On("Save", mock.Anything).Return(nil)

	usecase := NewCreateAccountUseCase(accountMock, clientMock)
	inputDto := CreateAccountInputDTO{
		ClientID: client.ID,
	}

	output, err := usecase.Execute(inputDto)

	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.NotEmpty(t, output.ID)

	clientMock.AssertExpectations(t)
	clientMock.AssertNumberOfCalls(t, "Get", 1)

	accountMock.AssertExpectations(t)
	accountMock.AssertNumberOfCalls(t, "Save", 1)

}
