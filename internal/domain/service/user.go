package service

import (
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/domain/entity"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/domain/repository"
)

// UserService defines an API for user service to be used by presentation layer
type UserService interface {
	CreateUser(user entity.User) (entity.User, error)
	GetUser(id int64, username string, email string) (*entity.User, error)
	UpdateUser(user entity.User) (entity.User, error)
	DeleteUser(id int64) error
}

// userService struct as a business layer between controller and repository
type userService struct {
	repository repository.UserRepository
}

// NewUserService a constructor for user service
func NewUserService(repository repository.UserRepository) UserService {
	return &userService{repository: repository}
}

func (s *userService) CreateUser(u entity.User) (entity.User, error) {
	err := s.repository.CreateUser(&u)
	if err != nil {
		return entity.User{}, err
	}
	return u, nil
}

// UpdateUser takes user information, updates his information and returns user struct
func (s *userService) UpdateUser(user entity.User) (entity.User, error) {
	err := s.repository.UpdateUser(&user)
	if err != nil {
		return entity.User{}, err
	}
	return user, nil
}

// DeleteUser takes user id and deletes users record
func (s *userService) DeleteUser(id int64) error {
	err := s.repository.DeleteUser(id)
	if err != nil {
		return err
	}
	return nil
}

// GetUser retrieves the user based on his id, username and/or email
func (s *userService) GetUser(id int64, username string, email string) (*entity.User, error) {
	user, err := s.repository.GetUser(id, username, email)
	if err != nil {
		return nil, err
	}
	return user, nil
}
