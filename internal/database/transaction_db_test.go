package database

import (
	"database/sql"
	"testing"

	"github.com/landucci/ms-wallet/internal/entity"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
)

type TransactionDBTestSuite struct {
	suite.Suite
	db            *sql.DB
	transactionDB *TransactionDB
	zelda         *entity.Client
	link          *entity.Client
	accountZelda  *entity.Account
	accountLink   *entity.Account
}

func (s *TransactionDBTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	s.Nil(err)
	s.db = db

	db.Exec("CREATE TABLE clients (id varchar(255), name varchar(255), email varchar(255), created_at date, updated_at date)")
	db.Exec("CREATE TABLE accounts (id varchar(255), client_id varchar(255), balance int, created_at date, updated_at date)")
	db.Exec("CREATE TABLE transactions (id varchar(255), account_from_id varchar(255), account_to_id varchar(255), amount float, created_at date)")

	s.transactionDB = NewTransactionDB(db)
	s.zelda, _ = entity.NewClient("Zelda", "zelda@hyrule.com")
	s.link, _ = entity.NewClient("Link", "link@hyrule.com")
	s.accountZelda = entity.NewAccount(s.zelda)
	s.accountZelda.Credit(100000.0)
	s.accountLink = entity.NewAccount(s.link)
	s.accountLink.Credit(10000.0)
}

func (s *TransactionDBTestSuite) TearDownSuite() {
	defer s.db.Close()
	s.db.Exec("DROP TABLE transactions")
	s.db.Exec("DROP TABLE accounts")
	s.db.Exec("DROP TABLE clients")
}

func TestTransactionDBTestSuite(t *testing.T) {
	suite.Run(t, new(TransactionDBTestSuite))
}

func (s *TransactionDBTestSuite) TestCreate() {
	s.db.Exec("INSERT INTO clients (id, name, email, created_at, updated_at) VALUES (?, ?, ?, ?, ?)",
		s.zelda.ID, s.zelda.Name, s.zelda.Email, s.zelda.CreatedAt, s.zelda.UpdatedAt)
	s.db.Exec("INSERT INTO clients (id, name, email, created_at, updated_at) VALUES (?, ?, ?, ?, ?)",
		s.link.ID, s.link.Name, s.link.Email, s.link.CreatedAt, s.link.UpdatedAt)
	s.db.Exec("INSERT INTO accounts (id, client_id, balance, created_at, updated_at) VALUES (?, ?, ?, ?, ?)",
		s.accountZelda.ID, s.accountZelda.Client.ID, s.accountZelda.Balance, s.accountZelda.CreatedAt, s.accountZelda.UpdatedAt)
	s.db.Exec("INSERT INTO accounts (id, client_id, balance, created_at, updated_at) VALUES (?, ?, ?, ?, ?)",
		s.accountLink.ID, s.accountLink.Client.ID, s.accountLink.Balance, s.accountLink.CreatedAt, s.accountLink.UpdatedAt)

	transaction, err := entity.NewTransaction(s.accountZelda, s.accountLink, 10000.0)
	s.Nil(err)

	err = s.transactionDB.Create(transaction)
	s.Nil(err)
	s.NotNil(transaction)
	s.Equal(s.accountZelda.ID, transaction.AccountFrom.ID)
	s.Equal(90000.0, s.accountZelda.Balance)

	s.Equal(s.zelda.ID, transaction.AccountFrom.Client.ID)
	s.Equal(s.zelda.Name, transaction.AccountFrom.Client.Name)
	s.Equal(s.zelda.Email, transaction.AccountFrom.Client.Email)

	s.Equal(s.link.ID, transaction.AccountTo.Client.ID)
	s.Equal(s.link.Name, transaction.AccountTo.Client.Name)
	s.Equal(s.link.Email, transaction.AccountTo.Client.Email)
}
