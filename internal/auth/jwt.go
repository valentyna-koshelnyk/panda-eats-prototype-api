package auth

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
	"strconv"
	"time"
)

//go:generate mockery --name=AuthService
type AuthService interface {
	Hash(s string) (string, error)
	VerifyPassword(userPassword string, providedPassword string) bool
}

type authService struct{}

func NewAuthService() AuthService {
	return &authService{}
}

//go:generate mockery --name=TokenService
type TokenService interface {
	GenerateToken(ID int64, email, role string) (string, error)
	VerifyToken(tokenString string) (jwt.MapClaims, error)
}

type tokenService struct{}

func NewTokenService() TokenService {
	return &tokenService{}
}

var secretKey = []byte(viper.GetString("secret.key"))

// HashPassword is used to encrypt the password before it is stored in the DB
func (a *authService) Hash(s string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(s), 14)
	if err != nil {
		log.Panic(err)
	}

	return string(bytes), nil
}

// VerifyPassword checks the input password while verifying it with the password in the DB.
func (a *authService) VerifyPassword(userPassword string, providedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(providedPassword), []byte(userPassword))
	check := true

	if err != nil {
		check = false
		return check
	}
	return check
}

// GenerateToken generates JWT token from email, role, user ID and returns it as string
func (t *tokenService) GenerateToken(ID int64, email, role string) (string, error) {
	userID := strconv.FormatInt(ID, 10)
	claims := &Claims{
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
