package service

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
	"time"

	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/auth"
)

//go:generate mockery --name=TokenService

// TokenService is an interface for the token service
type TokenService interface {
	GenerateToken(userID string) (string, error)
	ExtractIDFromToken(requestToken string, secret string) (string, error)
}

type tokenService struct{}

// NewTokenService creates a new instance of the TokenService
func NewTokenService() TokenService {
	return &tokenService{}
}

// GenerateToken generates JWT token from user ID and returns it as string
func (t *tokenService) GenerateToken(userID string) (string, error) {
	claims := &auth.Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 1)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secretKey := viper.GetString("secret.key")
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (t *tokenService) ExtractIDFromToken(requestToken string, secret string) (string, error) {
	token, err := jwt.ParseWithClaims(requestToken, &auth.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(*auth.Claims)

	if !ok && !token.Valid {
		return "", fmt.Errorf("invalid Token")
	}

	return claims.UserID, nil
}
