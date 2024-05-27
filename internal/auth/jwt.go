package auth

import (
	"errors"
	"github.com/golang-jwt/jwt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
	"time"
)

var secretKey = []byte(viper.GetString("secret.key"))

// HashPassword is used to encrypt the password before it is stored in the DB
func Hash(s string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(s), 14)
	if err != nil {
		log.Panic(err)
	}

	return string(bytes), nil
}

// VerifyPassword checks the input password while verifying it with the password in the DB.
func VerifyPassword(userPassword string, providedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(providedPassword), []byte(userPassword))
	check := true

	if err != nil {
		check = false
		return check
	}
	return check
}

// GenerateToken generates JWT token from email, role, user ID and returns it as string
func GenerateToken(email, role, userID string) (string, error) {
	claims := &Claims{
		Role:   role,
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			Subject:   email,
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// VerifyToken verifies a JWT token, a middleware for services which require authentication
func VerifyToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
