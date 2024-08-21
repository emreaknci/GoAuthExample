package handler

import (
	"github.com/emreaknci/goauthexample/internal/api/request"
	"github.com/emreaknci/goauthexample/internal/service"
	"github.com/gin-gonic/gin"
)

type AuthHandler interface {
	LogIn(c *gin.Context)
	Register(c *gin.Context)
}

type authHandler struct {
	service service.AuthService
}

func NewAuthHandler(service service.AuthService) AuthHandler {
	return &authHandler{service}
}


func (a *authHandler) LogIn(c *gin.Context) {
	var req request.LogInRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	res := a.service.LogIn(req.Email, req.Password)

	c.JSON(res.StatusCode, res)
}

func (a *authHandler) Register(c *gin.Context) {
	var req request.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	res := a.service.Register(req.Email, req.Password)

	c.JSON(res.StatusCode, res)
}


