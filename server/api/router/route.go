package route

import (
	"github.com/emreaknci/goauthexample/api/middleware"
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

func RegisterRoutes(container *dig.Container, router *gin.Engine) {
	AuthRoutes(container, router)
	router.GET("/test", middleware.Auth(), func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello, World!"})
	})
}
