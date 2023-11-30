package database

import (
	"database/sql"

	"github.com/landucci/ms-wallet/internal/entity"
)

type AccountDB struct {
	DB *sql.DB
}

func NewAccountDB(db *sql.DB) *AccountDB {
	return &AccountDB{DB: db}
}

func (db *AccountDB) FindByID(id string) (*entity.Account, error) {
	var account entity.Account
	var client entity.Client
	account.Client = &client

	var query = "SELECT a.id, a.client_id, a.balance, a.created_at, a.updated_at, c.id, c.name, c.email, c.created_at, c.updated_at"
	query += "  FROM accounts a JOIN clients c ON a.client_id = c.id"
	query += " WHERE a.id = ?"

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(id)
	err = row.Scan(&account.ID, &account.Client.ID, &account.Balance, &account.CreatedAt, &account.UpdatedAt,
		&client.ID, &client.Name, &client.Email, &client.CreatedAt, &client.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return &account, nil
}

func (db *AccountDB) Save(account *entity.Account) error {
	stmt, err := db.DB.Prepare("INSERT INTO accounts (id, client_id, balance, created_at, updated_at) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(account.ID, account.Client.ID, account.Balance, account.CreatedAt, account.UpdatedAt); err != nil {
		return err
	}

	return nil
}

func (db *AccountDB) UpdateBalance(account *entity.Account) error {
	stmt, err := db.DB.Prepare("UPDATE accounts SET balance = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(account.Balance, account.ID)
	if err != nil {
		return err
	}

	return nil
}
