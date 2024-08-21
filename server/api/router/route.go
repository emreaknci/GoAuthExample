package route

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

func RegisterRoutes(container *dig.Container, router *gin.Engine) {
	AuthRoutes(container, router)
}
