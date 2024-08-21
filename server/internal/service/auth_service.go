package service

import (
	"fmt"

	"github.com/emreaknci/goauthexample/internal/model"
	"github.com/emreaknci/goauthexample/internal/repository"
	"github.com/emreaknci/goauthexample/pkg/util/response"
	"github.com/emreaknci/goauthexample/pkg/util/security/hashing"
	"github.com/emreaknci/goauthexample/pkg/util/security/jwt"
)

type AuthService interface {
	LogIn(email string, password string) response.CustomResponse[string] //jwt token
	Register(email string, password string) response.CustomResponse[any]
}

type authService struct {
	repo repository.UserRepository
}

func NewAuthService(repo repository.UserRepository) AuthService {
	return &authService{repo}
}

func (s *authService) LogIn(email string, password string) response.CustomResponse[string] {
	user, err := s.repo.FindByFilter(map[string]interface{}{"email": email})
	if err != nil {
		return response.CustomResponse[string]{Message: "User not found", Status: false, StatusCode: 404, Error: err.Error()}
	}

	isValid, err := hashing.VerifyPasswordHash(password, user.PasswordHash, user.PasswordSalt)
	if !isValid {
		return response.CustomResponse[string]{Status: false, Error: err.Error(), Message: "Your password is incorrect", StatusCode: 401}
	}

	token, err := jwt.GenerateToken(fmt.Sprintf("%d", user.ID))
	if err != nil {
		return response.CustomResponse[string]{Status: false, Error: err.Error(), Message: "Token generation failed", StatusCode: 500}
	}

	return response.CustomResponse[string]{Status: true, StatusCode: 200, Data: token, Message: "Login successful"}

}

func (s *authService) Register(email string, password string) response.CustomResponse[any] {

	_, err := s.repo.FindByFilter(map[string]interface{}{"email": email})
	if err == nil {
		return response.CustomResponse[any]{Status: false, Message: "This email is already used", StatusCode: 409}
	}

	passwordHash, passwordSalt, err := hashing.CreatePasswordHash(password)
	if err != nil {
		return response.CustomResponse[any]{Status: false, Message: "Password hash generation failed", StatusCode: 500, Error: err.Error()}
	}

	user := model.User{
		Email:        email,
		PasswordHash: passwordHash,
		PasswordSalt: passwordSalt,
	}

	_, err = s.repo.Create(&user)
	if err != nil {
		return response.CustomResponse[any]{Status: false, Message: "User creation failed", StatusCode: 500, Error: err.Error()}
	}

	return response.CustomResponse[any]{Status: true, Message: "Registration successful", StatusCode: 201}
}
