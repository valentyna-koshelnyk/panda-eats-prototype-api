package entity

import (
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/errors"
	"gorm.io/gorm"
	"net/mail"
	"strings"
	"time"
)

type User struct {
	gorm.Model
	ID        int64     `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// NewUser creates a new user entity
func NewUser(username, password, email, role string) (*User, error) {
	user := &User{
		Username:  username,
		Password:  password,
		Email:     email,
		Role:      role,
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
	if user.Username == "" || user.Password == "" || user.Email == "" {
		return errors.ErrEmptyUserField
	}

	if strings.ContainsAny(user.Username, " \t\r\n") || strings.ContainsAny(user.Password, " \t\r\n") {
		return errors.ErrFieldWithSpaces
	}

	if len(user.Password) < 8 {
		return errors.ErrShortPassword
	}

	_, err := mail.ParseAddress(user.Email)
	if err != nil {
		return errors.ErrInvalidEmail
	}

	return nil
}
