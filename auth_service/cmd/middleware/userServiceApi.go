package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// UserServiceAuthMiddleware user_service yetkilendirmesi için middleware
func UserServiceAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKey := c.GetHeader("User-Service-Authorization")
		expectedApiKey := "USER_SERVICE_API_KEY"

		if apiKey != expectedApiKey {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Yetkisiz erişim"})
			c.Abort()
			return
		}

		c.Next()
	}
}
