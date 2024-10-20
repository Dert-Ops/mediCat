package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"gitlab.com/dert-ops/mediCat/mediCat-Dev.git/cmd/config"
	"gitlab.com/dert-ops/mediCat/mediCat-Dev.git/cmd/routes"
)

func logOriginMiddleware(c *gin.Context) {
	origin := c.Request.Header.Get("Origin")
	log.Printf("Request from Origin: %s", origin)
	c.Next()

	// Yanıt başlıklarını kontrol et
	for name, values := range c.Writer.Header() {
		if name == "Access-Control-Allow-Origin" {
			log.Printf("Access-Control-Allow-Origin: %v", values)
		}
	}
}

func main() {

	config.LoadEnv()

	// Logrus logger'ı başlat
	config.InitLogrusLogger()

	// Veritabanına bağlan
	config.ConnectDB()

	// Gin router'ı oluştur
	router := gin.Default()

	// CORS ayarları
	router.Use(logOriginMiddleware)

	// Kullanıcı rotalarını ekle
	routes.UserRoutes(router)

	// Sunucu portunu ortam değişkenlerinden al, yoksa varsayılan portu kullan
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Varsayılan port
	}

	// Sunucuyu başlat
	config.LogrusLogger.Infof("Server is running on port %s", port)
	if err := router.Run(":" + port); err != nil {
		config.LogrusLogger.Fatalf("Failed to run server: %v", err)
	}
}
