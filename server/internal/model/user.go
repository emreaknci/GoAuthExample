package model

type User struct {
	ID           uint   `json:"id" gorm:"primarykey"`
	Email        string `json:"email" gorm:"unique"`
	PasswordHash string `json:"password_hash"`
	PasswordSalt string `json:"password_salt"`
	RefreshToken string `json:"refresh_token"`
	RefreshTokenExpiry int64 `json:"refresh_token_expiry"`
}
