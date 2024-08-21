package token

import (
	"crypto/rand"
	"encoding/base64"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateRefreshToken() (string, int64, error) {
	token := make([]byte, 32)
	_, err := rand.Read(token)
	if err != nil {
		return "", 0, err
	}
	refreshToken := base64.URLEncoding.EncodeToString(token)

	expiry := time.Now().Add(7 * 24 * time.Hour).Unix()

	return refreshToken, expiry, nil
}

func GenerateToken(userID string) (string, error) {
	secretKey := os.Getenv("JWT_SECRET_KEY")
	expirationTime := time.Now().Add(time.Minute * 15)

	claims := &TokenClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateToken(tokenStr string) (*TokenClaims, error) {
	secretKey := os.Getenv("JWT_SECRET_KEY")

	token, err := jwt.ParseWithClaims(tokenStr, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*TokenClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, err
}
