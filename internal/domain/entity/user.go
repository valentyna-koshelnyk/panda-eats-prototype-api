package entity

import (
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/errors"
	"gorm.io/gorm"
	"net/mail"
	"time"
)

type User struct {
	gorm.Model
	ID           int64     `json:"id"`
	Email        string    `json:"email"`
	Password     string    `json:"password"`
	Role         string    `json:"role"`
	Token        string    `json:"token"`
	RefreshToken string    `json:"refresh_token"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// NewUser creates a new user entity
func NewUser(email, password, role string) (*User, error) {
	user := &User{
		Email:     email,
		Password:  password,
		Role:      string(role),
		CreatedAt: time.Now(),
		UpdatedAt: time.Time{},
	}

	if err := user.Validate(); err != nil {
		return nil, err
	}

	return user, nil
}

// Validate validates the user entity.
func (user *User) Validate() error {
	if user.Email == "" || user.Password == "" {
		return errors.ErrEmptyUserField
	}

	_, err := mail.ParseAddress(user.Email)
	if err != nil {
		return errors.ErrInvalidEmail
	}

	if len(user.Password) < 8 {
		return errors.ErrShortPassword
	}
	return nil
}

