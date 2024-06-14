package service

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/auth"
)

//go:generate mockery --name=TokenService

// TokenService is an interface for the token service
type TokenService interface {
	GenerateToken(userID string) (string, error)
	VerifyToken(tokenString string) error
	ExtractIDFromToken(requestToken string, secret string) (string, error)
}

type tokenService struct{}

// NewTokenService creates a new instance of the TokenService
func NewTokenService() TokenService {
	return &tokenService{}
}

var secretKey = []byte(viper.GetString("secret.key"))

// GenerateToken generates JWT token from email, role, user ID and returns it as string
func (t *tokenService) GenerateToken(userID string) (string, error) {
	claims := &auth.Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   userID,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 1)),
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
func (t *tokenService) VerifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return secretKey, nil
	})

	if err != nil {
		return err
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}

	return errors.New("invalid token")
}

func (t *tokenService) ExtractIDFromToken(requestToken string, secret string) (string, error) {
	token, err := jwt.Parse(requestToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Unexpected signing method")
		}
		return []byte(secret), nil
	})

	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok && !token.Valid {
		return "", fmt.Errorf("Invalid Token")
	}

	return claims["user_id"].(string), nil
}
