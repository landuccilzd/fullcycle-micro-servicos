package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateNewClient(t *testing.T) {
	client, err := NewClient("Princesa Zelda", "zelda@hyrule.com")
	assert.Nil(t, err)
	assert.Equal(t, "Princesa Zelda", client.Name)
	assert.Equal(t, "zelda@hyrule.com", client.Email)
}

func TestCreateNewClientWhenArgsAreInvalid(t *testing.T) {
	client, err := NewClient("", "")
	assert.NotNil(t, err)
	assert.Nil(t, client)
}

func TestUpdateClient(t *testing.T) {
	client, _ := NewClient("Princesa Zelda", "zelda@hyrule.com")
	err := client.Update("Link", "link@hyrule.com")
	assert.Nil(t, err)
	assert.Equal(t, "Link", client.Name)
	assert.Equal(t, "link@hyrule.com", client.Email)
}

func TestUpdateClientWithInvalidArgs(t *testing.T) {
	client, _ := NewClient("Princesa Zelda", "zelda@hyrule.com")
	err := client.Update("", "")
	assert.Error(t, err, "o nome é obrigatório")
}

func TestAddAccountToClient(t *testing.T) {
	client, _ := NewClient("Princesa Zelda", "zelda@hyrule.com")
	account := NewAccount(client)
	err := client.AddAccount(account)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(client.Accounts))

}
