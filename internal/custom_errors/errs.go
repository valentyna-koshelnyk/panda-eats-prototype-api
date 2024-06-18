package custom_errors

import (
	"errors"
)

var (
	// ErrEmptyUserField checks if all necessary fields of user struct are filled out.
	ErrEmptyUserField = errors.New("username, password and email can't be empty")
	// ErrFieldWithSpaces checks if no spaces were used since spaces can introduce ambiguity and potential security risks.
	ErrFieldWithSpaces = errors.New("username and password can't have spaces")
	// ErrShortPassword checks the length of password to avoid potential security risks.
	ErrShortPassword = errors.New("password shorter than 8 characters")
	// ErrInvalidEmail checks if email consists of valid characters.
	ErrInvalidEmail = errors.New("invalid email address")
	// ErrNotFound is used when the requested item is not found.
	ErrNotFound = errors.New("items not found")
	// ErrUserNotFound is used when the user doesn't have account in the system.
	ErrUserNotFound = errors.New("user not found")
	// ErrInvalidToken is used for invalid tokens.
	ErrInvalidToken = errors.New("Unexpected signing method")
)
