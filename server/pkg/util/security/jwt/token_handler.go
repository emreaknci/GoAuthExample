package jwt

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(userID string) (string, error) {
	secretKey := os.Getenv("JWT_SECRET_KEY")
	expirationTime := time.Now().Add(time.Duration(getExpirationTime()) * time.Second)

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

func getExpirationTime() int64 {
	expirationTime := os.Getenv("JWT_EXPIRATION_TIME")
	if expirationTime == "" {
		return 3600
	}

	expTime, err := strconv.ParseInt(expirationTime, 10, 64)
	if err != nil {
		log.Fatalf("Invalid JWT_EXPIRATION_TIME value: %v", err)
		return 3600
	}

	return expTime
}
