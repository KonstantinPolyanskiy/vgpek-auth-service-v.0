package repository

import (
	"awesomeProject5/types"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) GetAccount(login, password string) (types.Account, error) {
	var account types.Account

	query := fmt.Sprintf("SELECT id FROM %s WHERE login=$1 AND password_hash=$2", accountsTable)
	err := r.db.Get(&account, query, login, password)

	return account, err
}

func (r *AuthPostgres) CreateUser(user types.User, accountId int) (int, error) {
	var id int

	query := fmt.Sprintf("INSERT INTO %s (name, surname, phone_number, account_id) VALUES ($1, $2, $3, $4) RETURNING id", usersTable)

	row := r.db.QueryRow(query, user.Name, user.Surname, user.PhoneNumber, accountId)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AuthPostgres) CreateAccount(account types.Account) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (login, password_hash) VALUES ($1, $2) RETURNING id", accountsTable)

	row := r.db.QueryRow(query, account.Login, account.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}
