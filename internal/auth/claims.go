package auth

import (
	"github.com/golang-jwt/jwt/v4"
)

// Claims represents the JWT claims
type Claims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}
