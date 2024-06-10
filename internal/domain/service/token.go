package service

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"

	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/auth"
)

//go:generate mockery --name=TokenService

// TokenService is an interface for the token service
type TokenService interface {
	GenerateToken(ID int64, email, role string) (string, error)
	VerifyToken(tokenString string) (jwt.MapClaims, error)
}

type tokenService struct{}

// NewTokenService creates a new instance of the TokenService
func NewTokenService() TokenService {
	return &tokenService{}
}

var secretKey = []byte(viper.GetString("secret.key"))

// GenerateToken generates JWT token from email, role, user ID and returns it as string
func (t *tokenService) GenerateToken(ID int64, email, role string) (string, error) {
	userID := strconv.FormatInt(ID, 10)
	claims := &auth.Claims{
		Role:   role,
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   email,
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
func (t *tokenService) VerifyToken(tokenString string) (jwt.MapClaims, error) {
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

func (t *tokenService) ExtractIDFromToken(requestToken string, secret string) (string, error) {
	token, err := jwt.Parse(requestToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Unexpected signing method: %v", token.Header["alg"])
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

	return claims["id"].(string), nil
}

// TokenFromHeader tries to retreive the token string from the
// "Authorization" reqeust header: "Authorization: BEARER T".
func (t *tokenService) TokenFromHeader(r *http.Request) string {
	// Get token from authorization header.
	bearer := r.Header.Get("Authorization")
	if len(bearer) > 7 && strings.ToUpper(bearer[0:6]) == "BEARER" {
		return bearer[7:]
	}
	return ""
}
