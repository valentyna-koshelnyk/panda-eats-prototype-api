package service

import (
	"errors"

	log "github.com/sirupsen/logrus"

	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/domain/entity"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/domain/repository"
)

//go:generate mockery --name=UserService

// UserService defines an API for user service to be used by presentation layer
type UserService interface {
	CreateUser(user entity.User) (entity.User, error)
	VerifyUser(user entity.User) (bool, error)
	GenerateTokenResponse(email, password string) (string, error)
	GetUserById(id string) (*entity.User, error)
	GetUserByEmail(email string) (*entity.User, error)
}

var (
	token string
)

// userService struct as a business layer between controller and repository
type userService struct {
	repository repository.UserRepository
	auth       AuthService
	token      TokenService
}

// NewUserService a constructor for user service
func NewUserService(repository repository.UserRepository, auth AuthService, token TokenService) UserService {
	return &userService{
		repository: repository,
		auth:       auth,
		token:      token,
	}
}

// CreateUser presentation layer for adding user to repository
func (s *userService) CreateUser(u entity.User) (entity.User, error) {
	hashedPassword, err := s.auth.Hash(u.Password)
	u.Password = hashedPassword

	// For now keeping user as a role and keeping all registered users as "user" by default.
	// TODO: to add admin as a role and all methods related to admin role make accessible just for admins
	u.Role = "user"

	existingUser, err := s.GetUserByEmail(u.Email)
	if existingUser != nil {
		log.Error("User already exists")
		return u, errors.New("user already exists")
	}

	err = u.Validate()
	if err != nil {
		return entity.User{}, errors.New("invalid user")
	}

	err = s.repository.CreateUser(&u)
	if err != nil {
		return entity.User{}, err
	}
	return u, nil
}

// GetUserByid retrieves user based on his id
func (s *userService) GetUserById(id string) (*entity.User, error) {
	user, err := s.repository.GetUserByID(id)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// VerifyUser verifies if login and password are
func (s *userService) VerifyUser(u entity.User) (bool, error) {
	existingUser, err := s.GetUserByEmail(u.Email)
	if err != nil {
		return false, errors.New("invalid user")
	}
	if s.auth.VerifyPassword(u.Password, existingUser.Password) {
		return true, nil
	}
	return false, errors.New("invalid password")
}

func (s *userService) GenerateTokenResponse(email, password string) (string, error) {
	existingUser, err := s.GetUserByEmail(email)
	if err != nil {
		return "", errors.New("invalid user")
	}
	if !s.auth.VerifyPassword(password, existingUser.Password) {
		return "", errors.New("invalid password")
	}

	token, err = s.token.GenerateToken(existingUser.ID)
	if err != nil {
		return "", errors.New("invalid user")
	}
	return token, nil
}

// GetUserByEmail retrieves user based on his email
func (s *userService) GetUserByEmail(email string) (*entity.User, error) {
	user, err := s.repository.GetUserByEmail(email)
	if err != nil {
		return nil, err
	} else {
		return &user, nil
	}

}
