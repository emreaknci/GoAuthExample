package middleware

import (
	"net/http"

	"github.com/emreaknci/goauthexample/pkg/util/response"
	"github.com/emreaknci/goauthexample/pkg/util/security/token"
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {

		accessToken := c.GetHeader("Authorization")
		if accessToken == "" {
			c.JSON(http.StatusUnauthorized, response.CustomResponse[any]{Message: "Authorization header is required", Error: "Unauthorized", StatusCode: http.StatusUnauthorized})
			c.Abort()
			return
		}

		accessToken = accessToken[7:] // remove "Bearer " from accessToken
		if accessToken == "" {
			c.JSON(http.StatusUnauthorized, response.CustomResponse[any]{Message: "Authorization accessToken is required", Error: "Unauthorized", StatusCode: http.StatusUnauthorized})
			c.Abort()
			return
		}

		claims, err := token.ValidateToken(accessToken)

		if err != nil {
			c.JSON(http.StatusUnauthorized, response.CustomResponse[any]{Message: "Invalid accessToken", Error: "Unauthorized", StatusCode: http.StatusUnauthorized})
			c.Abort()
			return
		}

		c.Set("userId", claims.UserID)

		c.Next()
	}
}
