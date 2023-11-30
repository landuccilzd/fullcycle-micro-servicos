package createtransaction

import (
	"context"
	"testing"

	"github.com/landucci/ms-wallet/internal/entity"
	"github.com/landucci/ms-wallet/internal/event"
	"github.com/landucci/ms-wallet/internal/usecase/mocks"
	"github.com/landucci/ms-wallet/pkg/events"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateTransactionUseCase_Execute(t *testing.T) {
	zelda, _ := entity.NewClient("Zelda", "zelda@hyrule.com")
	accountZelda := entity.NewAccount(zelda)
	accountZelda.Credit(100000)

	link, _ := entity.NewClient("Link", "link@hyrule.com")
	accountLink := entity.NewAccount(link)
	accountLink.Credit(10000)

	mockUow := &mocks.UowMock{}
	mockUow.On("Do", mock.Anything, mock.Anything).Return(nil)

	input := CreateTransactionInputDTO{
		AccountIdFrom: accountZelda.ID,
		AccountIdTo:   accountLink.ID,
		Amount:        1000.0,
	}

	dispatcher := events.NewEventDispatcher()
	event := event.NewTransactionCreated()
	ctx := context.Background()

	usecase := NewCreateTransactionUseCase(mockUow, dispatcher, event)
	output, err := usecase.Execute(ctx, input)

	assert.Nil(t, err)
	assert.NotNil(t, output)
	mockUow.AssertExpectations(t)
	mockUow.AssertNotCalled(t, "Do", 1)
}
