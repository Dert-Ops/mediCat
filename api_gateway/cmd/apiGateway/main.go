package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var serviceMap = map[string]string{
	"user":  "http://user_service:8080",
	"media-service": "http://media-service-container:8081",
	"auth-service":  "http://auth-service-container:8082",
	"post-service":  "http://post-service-container:8083",
}

func proxyHandler(c *gin.Context) {
	// Rota veya path'e göre hangi servise yönlendirme yapılacağını belirle
	path := c.Request.URL.Path
	serviceName := strings.Split(path, "/")[1]

	// Servisin URL'sini map'ten bul
	serviceURL, exists := serviceMap[serviceName]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Servis bulunamadı"})
		return
	}

	// Servis URL'sini tam path ile birleştir
	proxyURL := serviceURL + c.Request.URL.Path[len("/"+serviceName):]
	if c.Request.URL.RawQuery != "" {
		proxyURL += "?" + c.Request.URL.RawQuery
	}

	// İsteği ilgili servise yönlendir
	resp, err := http.Get(proxyURL)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": serviceName + " servisine bağlanılamadı"})
		return
	}
	defer resp.Body.Close()

	// Yanıtı geri döndür
	c.Status(resp.StatusCode)
	io.Copy(c.Writer, resp.Body)
}

func main() {
    // Load environment variables from .env file if it exists
    err := godotenv.Load(".env")
    if err != nil {
        log.Println("No .env file found, continuing without loading environment variables from file.")
    }

    // Set up Gin router with CORS
    router := gin.Default()
    router.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"*"},
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
    }))

    // Example route
    router.GET("/health", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"status": "API Gateway is up and running"})
    })

	router.Any("/:service/*any", proxyHandler)

    // Start the server
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }
    router.Run(":" + port)
}