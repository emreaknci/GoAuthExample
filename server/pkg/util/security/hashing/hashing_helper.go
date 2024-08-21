package hashing

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"errors"
)

func generateRandomKey(length int) ([]byte, error) {
	if length <= 0 {
		return nil, errors.New("length must be greater than 0")
	}
	key := make([]byte, length)
	_, err := rand.Read(key)
	if err != nil {
		return nil, err
	}
	return key, nil
}

func CreatePasswordHash(password string) (passwordHash string, passwordSalt string, err error) {
	salt, err := generateRandomKey(64) 
	if err != nil {
		return "", "", err
	}

	hmac := hmac.New(sha512.New, salt)
	hmac.Write([]byte(password))
	passwordHashBytes := hmac.Sum(nil)

	passwordSalt = base64.StdEncoding.EncodeToString(salt)
	passwordHash = hex.EncodeToString(passwordHashBytes)
	return
}

func VerifyPasswordHash(password string, passwordHash string, passwordSalt string) (bool, error) {
	salt, err := base64.StdEncoding.DecodeString(passwordSalt)
	if err != nil {
		return false, err
	}
	
	hmac := hmac.New(sha512.New, salt)
	hmac.Write([]byte(password))
	computedHash := hmac.Sum(nil)
	expectedHash, err := hex.DecodeString(passwordHash)
	if err != nil {
		return false, err
	}

	if len(computedHash) != len(expectedHash) {
		return false, nil
	}

	for i, b := range computedHash {
		if b != expectedHash[i] {
			return false, nil
		}
	}
	return true, nil
}
