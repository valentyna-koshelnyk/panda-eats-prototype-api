package service

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/auth"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/domain/entity"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/domain/repository"
)

// UserService defines an API for user service to be used by presentation layer
//
//go:generate mockery --name=UserService
type UserService interface {
	CreateUser(user entity.User) (entity.User, error)
	GetUser(email string) (*entity.User, error)
	VerifyUser(user entity.User) (bool, error)
	GenerateTokenResponse(u entity.User) (*entity.Response, error)
}

var (
	token string
	items []entity.Item
)

// userService struct as a business layer between controller and repository
type userService struct {
	repository repository.UserRepository
	auth       auth.AuthService
	token      auth.TokenService
}

// NewUserService a constructor for user service
func NewUserService(repository repository.UserRepository, auth auth.AuthService, token auth.TokenService) UserService {
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

	existingUser, err := s.GetUser(u.Email)
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

// GetUser retrieves the user based on his id, username and/or email
func (s *userService) GetUser(email string) (*entity.User, error) {
	user, err := s.repository.GetUser(email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// VerifyUser verifies if login and password are
func (s *userService) VerifyUser(u entity.User) (bool, error) {
	existingUser, err := s.GetUser(u.Email)
	if err != nil {
		return false, errors.New("invalid user")
	}
	if s.auth.VerifyPassword(u.Password, existingUser.Password) {
		return true, nil
	}
	return false, errors.New("invalid password")
}

func (s *userService) GenerateTokenResponse(u entity.User) (*entity.Response, error) {
	verified, err := s.VerifyUser(u)

	if verified {
		token, err = s.token.GenerateToken(u.ID, u.Email, u.Role)
		if err != nil {
			return nil, errors.New("invalid user")
		}
		items = append(items, token)
		response := entity.NewResponse(items)
		return response, nil
	}

	return nil, errors.New("invalid user")
}
