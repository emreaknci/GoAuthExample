package jwt

import "github.com/golang-jwt/jwt/v5"

type TokenClaims struct {
	UserID string `json:"user"`
	jwt.RegisteredClaims
}
