package service

import (
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

//go:generate mockery --name=AuthService

// AuthService is an interface for the authentication service
type AuthService interface {
	Hash(s string) (string, error)
	VerifyPassword(userPassword string, providedPassword string) bool
}

type authService struct{}

// NewAuthService creates a new instance of the AuthService
func NewAuthService() AuthService {
	return &authService{}
}

// Hash is used to encrypt the password before it is stored in the DB
func (a *authService) Hash(s string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(s), 14)
	if err != nil {
		logrus.Panic(err)
	}

	return string(bytes), nil
}

// VerifyPassword checks the input password while verifying it with the password in the DB.
func (a *authService) VerifyPassword(userPassword string, providedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(providedPassword), []byte(userPassword))
	check := true

	if err != nil {
		check = false
		return check
	}
	return check
}
