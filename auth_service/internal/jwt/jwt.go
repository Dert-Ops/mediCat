package jwt

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"gitlab.com/dert-ops/mediCat/mediCat-Dev.git/cmd/config"
)

// Token oluşturma fonksiyonu
func GenerateJWT(userID string) (string, error) {
	claims := jwt.MapClaims{}
	claims["user_id"] = userID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix() // Token geçerlilik süresi

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.SecretKey)) // SecretKey kullanımı
}
