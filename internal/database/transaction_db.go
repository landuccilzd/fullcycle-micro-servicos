package database

import (
	"database/sql"

	"github.com/landucci/ms-wallet/internal/entity"
)

type TransactionDB struct {
	DB *sql.DB
}

func NewTransactionDB(db *sql.DB) *TransactionDB {
	return &TransactionDB{DB: db}
}

func (db *TransactionDB) Create(transaction *entity.Transaction) error {
	stmt, err := db.DB.Prepare("INSERT INTO transactions (id, account_from_id, account_to_id, amount, created_at) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(transaction.ID, transaction.AccountFrom.ID, transaction.AccountTo.ID, transaction.Amount, transaction.CreatedAt); err != nil {
		return err
	}

	return nil
}
