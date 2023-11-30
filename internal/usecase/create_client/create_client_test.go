package createclient

import (
	"testing"

	"github.com/landucci/ms-wallet/internal/usecase/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateClientUseCase_Execute(t *testing.T) {
	m := &mocks.ClientGatewayMock{}
	m.On("Save", mock.Anything).Return(nil)
	usecase := NewCreateClientUseCase(m)

	output, err := usecase.Execute(CreateClientInputDTO{
		Name:  "Princesa Zelda",
		Email: "zelda@hyrule.com",
	})

	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.NotEmpty(t, output.ID)
	assert.Equal(t, "Princesa Zelda", output.Name)
	assert.Equal(t, "zelda@hyrule.com", output.Email)

	m.AssertExpectations(t)
	m.AssertNumberOfCalls(t, "Save", 1)

}
