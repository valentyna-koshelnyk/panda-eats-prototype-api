package repository

import (
	log "github.com/sirupsen/logrus"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/domain/entity"
	"gorm.io/gorm"
)

// UserRepository interface for interacting with db
type UserRepository interface {
	GetUser(email string) (*entity.User, error)
	CreateUser(u *entity.User) error
}

// userRepository layer for interacting with db
type userRepository struct {
	db *gorm.DB
}

// NewUserRepository creates a new instance of UserRepository
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

// GetUser checks if user with the input email exists
func (r *userRepository) GetUser(email string) (*entity.User, error) {
	var u entity.User
	result := r.db.Where("email = ?", email).First(&u)

	if result.Error != nil {
		log.Error("user not found: ", result.Error)
		return nil, result.Error
	}
	return &u, nil
}

func (r *userRepository) CreateUser(u *entity.User) error {
	result := r.db.Create(&u)
	if result.Error != nil {
		log.Error(result.Error)
	}
	return nil
}
