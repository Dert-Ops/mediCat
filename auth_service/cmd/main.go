package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gitlab.com/dert-ops/mediCat/mediCat-Dev.git/cmd/config"
	"gitlab.com/dert-ops/mediCat/mediCat-Dev.git/cmd/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		config.InitLogrusLogger()
		config.LogrusLogger.Warn("No .env file found, continuing without loading environment variables.")
	}
	config.LoadEnv()

	config.InitLogrusLogger()
	// rabbitmq.InitRabbitMQ()
	// rabbitmq.CreateQueue("email_queue")
	// Veritabanına bağlan
	config.ConnectDB()

	// Gin router'ı oluştur
	router := gin.Default()

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
