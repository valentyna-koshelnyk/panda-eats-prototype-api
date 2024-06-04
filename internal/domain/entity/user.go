package entity

import (
	"net/mail"
	"time"

	"gorm.io/gorm"

	ce "github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/custom-errors"
)

// User represents a user entity/
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

// Validate validates the user entity.
func (user *User) Validate() error {
	if user.Email == "" || user.Password == "" {
		return ce.ErrEmptyUserField
	}

	_, err := mail.ParseAddress(user.Email)
	if err != nil {
		return ce.ErrInvalidEmail
	}

	if len(user.Password) < 8 {
		return ce.ErrShortPassword
	}
	return nil
}
