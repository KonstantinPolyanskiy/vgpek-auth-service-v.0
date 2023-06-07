package service

import (
	"awesomeProject5/pkg/repository"
	"awesomeProject5/types"
)

type Authorization interface {
	CreateUser(user types.User, accountId int) (int, error)
	CreateAccount(account types.Account, name string) (int, error)
	GenerateLogin(name string) string
	GeneratePassword() string
	GenerateToken(login string, password string) (string, error)
}
type Service struct {
	Authorization
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
