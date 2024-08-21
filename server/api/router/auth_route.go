package route

import (
	"github.com/emreaknci/goauthexample/api/handler"
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

func AuthRoutes(container *dig.Container, router *gin.Engine) {
	err := container.Invoke(func(authController handler.AuthHandler) {
		router.POST("/auth/login", authController.LogIn)
		router.POST("/auth/register", authController.Register)
		router.POST("/auth/refresh", authController.RefreshToken)
	})

	if err != nil {
		panic(err)
	}
}
