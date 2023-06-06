package service

import (
	"awesomeProject5/pkg/repository"
	"awesomeProject5/types"
	"crypto/sha1"
	"fmt"
	"github.com/mdigger/translit"
	"github.com/sethvargo/go-password/password"
)

const salt = "fgnjgdfgdfgdfdfdsaaa"

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user types.User, accountId int) (int, error) {
	return s.repo.CreateUser(user, accountId)
}

func (s *AuthService) CreateAccount(account types.Account, name string) (int, error) {
	psw := s.GeneratePassword()

	account.Login = s.GenerateLogin(name)
	account.Password = s.GeneratePasswordHash(psw)

	return s.repo.CreateAccount(account, name)
}

func (s *AuthService) GenerateLogin(name string) string {
	return translit.Ru(name)
}

func (s *AuthService) GeneratePassword() string {
	res, _ := password.Generate(10, 5, 5, false, false)
	return res
}
func (s *AuthService) GeneratePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
