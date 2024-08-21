package config

import (
	"fmt"
	"log"
	"os"

	"github.com/emreaknci/goauthexample/internal/model"
	"github.com/emreaknci/goauthexample/pkg/util/security/hashing"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDatabase() (*gorm.DB, error) {
	db, err := newDatabase()
	if err != nil {
		return nil, err
	}

	err = migrateModels(db)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func newDatabase() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))
	if dsn == "" {
		log.Fatal("DATABASE_DSN environment variable is not set")
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func migrateModels(db *gorm.DB) error {
	err := db.AutoMigrate(&model.User{})
	if err != nil {
		panic("Failed to auto-migrate database: " + err.Error())
	}

	err = userSeeds(db)
	if err != nil {
		panic("Failed to seed users: " + err.Error())
	}

	return nil
}

func userSeeds(db *gorm.DB) error {
	var count int64
	db.Model(&model.User{}).Count(&count)
	if count > 0 {
		return nil
	}

	hash, salt, err := hashing.CreatePasswordHash("123456")
	if err != nil {
		return err
	}

	users := []model.User{
		{
			Email:        "emreaknci@github.com",
			PasswordHash: hash,
			PasswordSalt: salt,
		},
	}

	for _, user := range users {
		err := db.Create(&user).Error
		if err != nil {
			return err
		}
	}

	return nil
}
