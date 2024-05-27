package service

import (
	log "github.com/sirupsen/logrus"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/auth"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/domain/entity"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/domain/repository"
)

// UserService defines an API for user service to be used by presentation layer
type UserService interface {
	CreateUser(user entity.User) (entity.User, error)
	GetUser(email string) (*entity.User, error)
}

// userService struct as a business layer between controller and repository
type userService struct {
	repository repository.UserRepository
}

// NewUserService a constructor for user service
func NewUserService(repository repository.UserRepository) UserService {
	return &userService{repository: repository}
}

// CreateUser presentation layer for adding user to repository
func (s *userService) CreateUser(u entity.User) (entity.User, error) {
	hashedPassword, err := auth.Hash(u.Password)
	u.Password = hashedPassword

	// For now keeping user as a role and keeping all registered users as "user" by default.
	// TODO: to add admin as a role and all methods related to admin role make accessible just for admins
	u.Role = "user"

	existingUser, err := s.GetUser(u.Email)
	if existingUser != nil {
		log.Error("User already exists")
	}

	err = u.Validate()
	if err != nil {
		return entity.User{}, err
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
