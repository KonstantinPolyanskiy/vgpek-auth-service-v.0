package service

import (
	"awesomeProject5/pkg/repository"
	"awesomeProject5/types"
	"crypto/sha1"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/mdigger/translit"
	"github.com/sethvargo/go-password/password"
	"math/rand"
	"time"
)

const (
	salt       = "fgnjgdfgdfgdfdfdsaaa"
	signingKey = "435h734fggfdg4%$$@fd2buksq"
)

type AuthService struct {
	repo repository.Authorization
}
type tokenClaims struct {
	jwt.StandardClaims
	AccountID int `json:"account_id"`
}

func (s *AuthService) GenerateToken(login string, password string) (string, error) {
	account, err := s.repo.GetAccount(login, s.GeneratePasswordHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(12 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		account.Id,
	})

	return token.SignedString([]byte(signingKey))
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

	return s.repo.CreateAccount(account)
}

func (s *AuthService) GenerateLogin(name string) string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%s%d", translit.Ru(name), (rand.Intn(1000 - 1)))
}

func (s *AuthService) GeneratePassword() string {
	res, _ := password.Generate(5, 0, 5, false, false)
	return res
}
func (s *AuthService) GeneratePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
