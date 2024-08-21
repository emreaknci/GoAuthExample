package middleware

import (
	"net/http"

	"github.com/emreaknci/goauthexample/pkg/util/response"
	"github.com/emreaknci/goauthexample/pkg/util/security/jwt"
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {

		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, response.CustomResponse[any]{Message: "Authorization header is required", Error: "Unauthorized", StatusCode: http.StatusUnauthorized})
			c.Abort()
			return
		}

		token = token[7:] // remove "Bearer " from token
		if token == "" {
			c.JSON(http.StatusUnauthorized, response.CustomResponse[any]{Message: "Authorization token is required", Error: "Unauthorized", StatusCode: http.StatusUnauthorized})
			c.Abort()
			return
		}

		claims, err := jwt.ValidateToken(token)

		if err != nil {
			c.JSON(http.StatusUnauthorized, response.CustomResponse[any]{Message: "Invalid token", Error: "Unauthorized", StatusCode: http.StatusUnauthorized})
			c.Abort()
			return
		}

		c.Set("userId", claims.UserID)

		c.Next()
	}
}
