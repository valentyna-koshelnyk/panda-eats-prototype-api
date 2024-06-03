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
	//ErrInvalidEmail checks if email consists of valid characters.
	ErrInvalidEmail = errors.New("invalid email address")
)
