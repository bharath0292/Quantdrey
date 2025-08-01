package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {

	var validAPIKey = os.Getenv("API_AUTH_KEY")

	return func(c *gin.Context) {
		if c.Request.Method == http.MethodOptions {
			c.Next()
			return
		}

		authHeader := c.GetHeader("Authorization")

		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": "Unauthorized",
			})
			c.Abort()
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")

		if token != validAPIKey {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": "Unauthorized",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
