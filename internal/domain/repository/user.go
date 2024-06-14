package repository

import (
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"

	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/domain/entity"
)

// UserRepository interface for interacting with db
type UserRepository interface {
	CreateUser(u *entity.User) error
	GetUserByID(ID string) (entity.User, error)
	GetUserByEmail(email string) (entity.User, error)
}

// userRepository layer for interacting with db
type userRepository struct {
	db *gorm.DB
}

// NewUserRepository creates a new instance of UserRepository
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

// CreateUser creates a new user
func (r *userRepository) CreateUser(u *entity.User) error {
	result := r.db.Create(&u)
	if result.Error != nil {
		log.Error(result.Error)
	}
	return nil
}

// GetUserByID retrieves user based on his id
func (r *userRepository) GetUserByID(ID string) (entity.User, error) {
	var u entity.User
	result := r.db.Where("ID = ?", ID).First(&u)

	if result.Error != nil {
		log.Error("user not found: ", result.Error)
		return entity.User{}, result.Error
	}
	return u, nil
}

// GetUserByEmail retrieves user based on his email
func (r *userRepository) GetUserByEmail(email string) (entity.User, error) {
	var u entity.User
	result := r.db.Where("email = ?", email).First(&u)

	if result.Error != nil {
		log.Error("user not found: ", result.Error)
		return entity.User{}, result.Error
	}
	return u, nil
}
