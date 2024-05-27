package auth

import "github.com/golang-jwt/jwt"

type Claims struct {
	Role   string `json:"role"`
	UserID string `json:"userid,omitempty"`
	jwt.StandardClaims
}
