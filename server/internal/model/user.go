package model

type User struct {
	ID           uint   `json:"id" gorm:"primarykey"`
	Email        string `json:"email" gorm:"unique"`
	PasswordHash string `json:"password_hash"`
	PasswordSalt string `json:"password_salt"`
}
