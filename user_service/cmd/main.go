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

	router := gin.Default()
	routes.UserRoutes(router)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	config.LogrusLogger.Infof("User Service is running on port %s", port)
	if err := router.Run(":" + port); err != nil {
		config.LogrusLogger.Fatalf("Failed to run server: %v", err)
	}
}
