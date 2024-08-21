package service

import (
	"fmt"
	"time"

	"github.com/emreaknci/goauthexample/internal/model"
	"github.com/emreaknci/goauthexample/internal/repository"
	"github.com/emreaknci/goauthexample/pkg/util/response"
	"github.com/emreaknci/goauthexample/pkg/util/security/hashing"
	"github.com/emreaknci/goauthexample/pkg/util/security/token"
)

type TokenResponse struct {
	AcessToken   string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type AuthService interface {
	LogIn(email string, password string) response.CustomResponse[TokenResponse]
	Register(email string, password string) response.CustomResponse[any]
	RefreshToken(refreshToken string) response.CustomResponse[TokenResponse]
}

type authService struct {
	repo repository.UserRepository
}

func NewAuthService(repo repository.UserRepository) AuthService {
	return &authService{repo}
}

func (s *authService) LogIn(email string, password string) response.CustomResponse[TokenResponse] {
	user, err := s.repo.FindByFilter(map[string]interface{}{"email": email})
	if err != nil {
		return response.CustomResponse[TokenResponse]{Message: "User not found", Status: false, StatusCode: 404, Error: err.Error()}
	}

	isValid, err := hashing.VerifyPasswordHash(password, user.PasswordHash, user.PasswordSalt)
	if !isValid {
		return response.CustomResponse[TokenResponse]{Status: false, Error: err.Error(), Message: "Your password is incorrect", StatusCode: 401}
	}

	accessToken, err := token.GenerateToken(fmt.Sprintf("%d", user.ID))
	if err != nil {
		return response.CustomResponse[TokenResponse]{Status: false, Error: err.Error(), Message: "Token generation failed", StatusCode: 500}
	}

	refreshToken, expiry, err := token.GenerateRefreshToken()
	if err != nil {
		return response.CustomResponse[TokenResponse]{Status: false, Error: err.Error(), Message: "Refresh token generation failed", StatusCode: 500}
	}

	user.RefreshToken = refreshToken
	user.RefreshTokenExpiry = expiry
	_, err = s.repo.Update(user)
	if err != nil {
		return response.CustomResponse[TokenResponse]{Status: false, Error: err.Error(), Message: "Refresh token update failed", StatusCode: 500}
	}

	return response.CustomResponse[TokenResponse]{Status: true, StatusCode: 200, Data: TokenResponse{AcessToken: accessToken, RefreshToken: refreshToken}, Message: "Login successful"}
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

func (s *authService) RefreshToken(refreshToken string) response.CustomResponse[TokenResponse] {
	user, err := s.repo.FindByFilter(map[string]interface{}{"refresh_token": refreshToken})
	if err != nil {
		return response.CustomResponse[TokenResponse]{Message: "User not found", Status: false, StatusCode: 404, Error: err.Error()}
	}

	if user.RefreshTokenExpiry < time.Now().Unix() {
		return response.CustomResponse[TokenResponse]{Status: false, Message: "Refresh token expired", StatusCode: 401}
	}

	accessToken, err := token.GenerateToken(fmt.Sprintf("%d", user.ID))
	if err != nil {
		return response.CustomResponse[TokenResponse]{Status: false, Error: err.Error(), Message: "Token generation failed", StatusCode: 500}
	}

	refreshToken, expiry, err := token.GenerateRefreshToken()
	if err != nil {
		return response.CustomResponse[TokenResponse]{Status: false, Error: err.Error(), Message: "Refresh token generation failed", StatusCode: 500}
	}

	user.RefreshToken = refreshToken
	user.RefreshTokenExpiry = expiry
	_, err = s.repo.Update(user)
	if err != nil {
		return response.CustomResponse[TokenResponse]{Status: false, Error: err.Error(), Message: "Refresh token update failed", StatusCode: 500}
	}

	return response.CustomResponse[TokenResponse]{Status: true, StatusCode: 200, Data: TokenResponse{AcessToken: accessToken, RefreshToken: refreshToken}, Message: "Token refreshed"}
}
