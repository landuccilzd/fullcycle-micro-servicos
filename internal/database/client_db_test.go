package database

import (
	"database/sql"
	"testing"

	"github.com/landucci/ms-wallet/internal/entity"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
)

type ClientDBTestSuite struct {
	suite.Suite
	db       *sql.DB
	clientDB *ClientDB
}

func (s *ClientDBTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	s.Nil(err)
	s.db = db

	db.Exec("CREATE TABLE clients (id varchar(255), name varchar(255), email varchar(255), created_at date, updated_at date)")
	s.clientDB = NewClientDB(db)
}

func (s *ClientDBTestSuite) TearDownSuite() {
	defer s.db.Close()
	s.db.Exec("DROP TABLE clients")
}

func TestClientDBTestSuite(t *testing.T) {
	suite.Run(t, new(ClientDBTestSuite))
}

func (s *ClientDBTestSuite) TestSave() {
	client := &entity.Client{
		ID:    "1",
		Name:  "Zelda",
		Email: "zelda@hyrule.com",
	}

	err := s.clientDB.Save(client)
	s.Nil(err)
}

func (s *ClientDBTestSuite) TestGet() {
	client, _ := entity.NewClient("Zelda", "zelda@hyrule.com")
	s.clientDB.Save(client)

	zelda, err := s.clientDB.Get(client.ID)
	s.Nil(err)
	s.NotNil(zelda)
	s.Equal(client.ID, zelda.ID)
	s.Equal(client.Name, zelda.Name)
	s.Equal(client.Email, zelda.Email)
}
