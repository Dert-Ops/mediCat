package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
    // Load environment variables from .env file if it exists
    err := godotenv.Load()
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

	router.GET("/yarra", func(c *gin.Context) {
		c.JSON(http.StatusBadGateway, gin.H{"status": "yarra yedin"})
	})

    // Start the server
    port := os.Getenv("PORT")
    if port == "" {
        port = "8083"
    }
    router.Run(":" + port)
}