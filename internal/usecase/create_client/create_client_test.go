package createclient

import (
	"testing"

	"github.com/landucci/ms-wallet/internal/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type ClientGatewayMock struct {
	mock.Mock
}

func (m *ClientGatewayMock) Get(id string) (*entity.Client, error) {
	args := m.Called(id)
	return args.Get(0).(*entity.Client), args.Error(1)
}

func (m *ClientGatewayMock) Save(client *entity.Client) error {
	args := m.Called(client)
	return args.Error(0)
}

func TestCreateClientUseCase_Execute(t *testing.T) {
	m := &ClientGatewayMock{}
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
