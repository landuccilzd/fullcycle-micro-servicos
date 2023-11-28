package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateAccount(t *testing.T) {
	client, _ := NewClient("Princesa Zelda", "zelda@hyrule.com")
	account := NewAccount(client)

	assert.NotNil(t, account)
	assert.Equal(t, client.ID, account.Client.ID)
	assert.Equal(t, client.Name, account.Client.Name)
	assert.Equal(t, client.Email, account.Client.Email)
}

func TestCreateAccountWithNoClient(t *testing.T) {
	account := NewAccount(nil)
	assert.Nil(t, account)
}

func TestCreditAccount(t *testing.T) {
	client, _ := NewClient("Princesa Zelda", "zelda@hyrule.com")
	account := NewAccount(client)

	assert.NotNil(t, account)
	assert.Equal(t, client.ID, account.Client.ID)
	assert.Equal(t, client.Name, account.Client.Name)
	assert.Equal(t, client.Email, account.Client.Email)
	assert.Equal(t, 0.0, account.Balance)

	account.Credit(1000.0)
	assert.Equal(t, 1000.0, account.Balance)
}

func TestDebitAccount(t *testing.T) {
	client, _ := NewClient("Princesa Zelda", "zelda@hyrule.com")
	account := NewAccount(client)

	assert.NotNil(t, account)
	assert.Equal(t, client.ID, account.Client.ID)
	assert.Equal(t, client.Name, account.Client.Name)
	assert.Equal(t, client.Email, account.Client.Email)
	assert.Equal(t, 0.0, account.Balance)

	account.Credit(1000.0)
	assert.Equal(t, 1000.0, account.Balance)

	account.Debit(500.0)
	assert.Equal(t, 500.0, account.Balance)
}
