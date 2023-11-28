package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateTransaction(t *testing.T) {
	zelda, _ := NewClient("Zelda", "zelda@hyrule.com")
	contaZelda := NewAccount(zelda)
	link, _ := NewClient("Link", "link@hyrule.com")
	contaLink := NewAccount(link)

	contaZelda.Credit(100000.0)
	contaLink.Credit(10000.0)

	transaction, err := NewTransaction(contaZelda, contaLink, 10000.0)
	assert.Nil(t, err)
	assert.NotNil(t, transaction)
	assert.Equal(t, 90000.0, contaZelda.Balance)
	assert.Equal(t, 20000.0, contaLink.Balance)
}

func TestCreateTransactionWithInsuficientFounds(t *testing.T) {
	zelda, _ := NewClient("Zelda", "zelda@hyrule.com")
	contaZelda := NewAccount(zelda)
	link, _ := NewClient("Link", "link@hyrule.com")
	contaLink := NewAccount(link)

	contaZelda.Credit(100000.0)
	contaLink.Credit(10000.0)

	transaction, err := NewTransaction(contaZelda, contaLink, 150000.0)
	assert.NotNil(t, err)
	assert.Nil(t, transaction)
	assert.Error(t, err, "a conta de origem não tem saldo suficiente")
	assert.Equal(t, 100000.0, contaZelda.Balance)
	assert.Equal(t, 10000.0, contaLink.Balance)
}
