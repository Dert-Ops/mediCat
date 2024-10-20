package middleware

import (
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, "Authorization header missing")
			c.Abort()
			return
		}

		bearerToken := strings.TrimPrefix(tokenString, "Bearer ")

		if bearerToken == tokenString { // Eğer değişmedi ise, Bearer kelimesi yoktur
			c.JSON(http.StatusUnauthorized, "Invalid token format")
			c.Abort()
			return
		}

		token, err := jwt.Parse(bearerToken, func(token *jwt.Token) (interface{}, error) {
			return []byte("selamdostumyagmurvarmiorda"), nil
		})
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, "Invalid token")
			c.Abort()
			return
		}

		c.Next() // Token geçerli, işlem devam ediyor.
	}
}
