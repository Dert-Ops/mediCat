package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gitlab.com/dert-ops/mediCat/mediCat-Dev.git/cmd/config"
	"gitlab.com/dert-ops/mediCat/mediCat-Dev.git/cmd/middleware"
)

var serviceMap = map[string]string{
	"users":         "http://user_service:8080",
	"media-service": "http://media-service-container:8081",
	"auth":          "http://auth_service:8080",
	"post-service":  "http://post-service-container:8080",
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

	req, err := http.NewRequest(c.Request.Method, proxyURL, c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Yeni istek oluşturulamadı"})
		return
	}

	// Orijinal isteğin header'larını yeni isteğe kopyala
	for key, values := range c.Request.Header {
		for _, value := range values {
			req.Header.Add(key, value)
		}
	}

	if token := c.Request.Header.Get("Authorization"); token != "" {
		req.Header.Set("Authorization", token)
	}

	// İsteği ilgili servise yönlendir
	fmt.Print(req)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": serviceName + " servisine bağlanılamadı"})
		return
	}
	defer resp.Body.Close()

	// Yanıtı geri döndür
	c.Status(resp.StatusCode)
	for key, values := range resp.Header {
		for _, value := range values {
			c.Writer.Header().Add(key, value)
		}
	}
	io.Copy(c.Writer, resp.Body)
}

func main() {
	// Load environment variables from .env file if it exists
	config.LoadEnv()

	// Set up Gin router with CORS
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://45.9.30.65"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Example route
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "API Gateway is up and running"})
	})

	router.POST("/auth/signup", proxyHandler)
	router.POST("/auth/signin", proxyHandler)

	JWTProtected := router.Group("/")
	JWTProtected.Use(middleware.AuthMiddleware())
	{
		log.Println("[!!!]  middleware gecildi")
		JWTProtected.Any("/:service/*any", proxyHandler)
	}

	// Start the server
	port := "8080"
	router.Run(":" + port)
}
