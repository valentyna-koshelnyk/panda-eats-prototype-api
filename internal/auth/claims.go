package auth

import (
	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	Role   string `json:"role"`
	UserID string `json:"userid,omitempty"`
	jwt.RegisteredClaims
}
