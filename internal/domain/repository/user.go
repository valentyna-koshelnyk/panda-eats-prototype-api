package repository

import (
	log "github.com/sirupsen/logrus"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/domain/entity"
	"gorm.io/gorm"
)

// UserRepository interface for interacting with db
type UserRepository interface {
	GetUser(id int64, username string, email string) (*entity.User, error)
	CreateUser(u *entity.User) error
	UpdateUser(u *entity.User) error
	DeleteUser(id int64) error
}

// userRepository layer for interacting with db
type userRepository struct {
	db *gorm.DB
}

// NewUserRepository creates a new instance of UserRepository
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) GetUser(id int64, username string, email string) (*entity.User, error) {
	var u entity.User
	result := r.db.First(&u, id)
	query := r.db
	if id != 0 {
		query = query.Where("id = ?", id)
	}
	if username != "" {
		query = query.Where("username = ?", username)
	}
	if email != "" {
		query = query.Where("email = ?", email)
	}

	result = query.First(&u)
	if result.Error != nil {
		log.Error("No user found: ", result.Error)
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

func (r *userRepository) UpdateUser(u *entity.User) error {
	result := r.db.Save(&u)
	if result.Error != nil {
		log.Error(result.Error)
	}
	return nil
}

func (r *userRepository) DeleteUser(id int64) error {
	result := r.db.Delete(&entity.User{}, id)
	if result.Error != nil {
		log.Error(result.Error)
	}
	return nil
}
