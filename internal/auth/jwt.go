package auth

import (
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

// HashPassword is used to encrypt the password before it is stored in the DB
func Hash(s string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(s), 14)
	if err != nil {
		log.Panic(err)
	}

	return string(bytes), nil
}
