package ownJwt

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Token oluşturma fonksiyonu
func GenerateJWT(userID string) (string, error) {
	claims := jwt.MapClaims{}
	claims["user_id"] = userID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix() // Token geçerlilik süresi

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte("selamdostumyagmurvarmiorda")) // SecretKey kullanımı
}
