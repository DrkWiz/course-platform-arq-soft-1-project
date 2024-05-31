package middleware

import (
	"backend/services/users"
	e "backend/utils/errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, e.NewUnauthorizedApiError("Authorization header is required"))
			c.Abort()
			return
		}

		token := strings.Split(authHeader, "Bearer ")[1]
		_, err := users.ValidateToken(token)
		if err != nil {
			c.JSON(err.Status(), err)
			c.Abort()
			return
		}

		c.Next()
	}
}
