package repository

import (
	"github.com/emreaknci/goauthexample/internal/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *model.User) (*model.User, error)
	Update(user *model.User) (*model.User, error)
	Delete(user *model.User) error

	FindByFilter(filter map[string]interface{}) (*model.User, error)
	FindAll() ([]*model.User, error)
	FindAllWithFilter(filter map[string]interface{}) ([]*model.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) Create(user *model.User) (*model.User, error) {
	if err := r.db.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) Update(user *model.User) (*model.User, error) {
	if err := r.db.Save(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) Delete(user *model.User) error {
	return r.db.Delete(&user).Error
}

func (r *userRepository) FindByFilter(filter map[string]interface{}) (*model.User, error) {
	var user model.User
	if err := r.db.Where(filter).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) FindAll() ([]*model.User, error) {
	var users []*model.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (r *userRepository) FindAllWithFilter(filter map[string]interface{}) ([]*model.User, error) {
	var users []*model.User
	if err := r.db.Where(filter).Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}
