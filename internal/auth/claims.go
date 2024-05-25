package auth

import "github.com/dgrijalva/jwt-go"

type Claims struct {
	Role   string `json:"role"`
	UserID string `json:"userid,omitempty"`
	jwt.StandardClaims
}
