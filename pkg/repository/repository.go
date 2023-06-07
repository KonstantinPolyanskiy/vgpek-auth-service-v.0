package repository

import (
	"awesomeProject5/types"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	GetAccount(login, password string) (types.Account, error)
	CreateUser(user types.User, accountId int) (int, error)
	CreateAccount(account types.Account) (int, error)
}
type Repository struct {
	Authorization
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
