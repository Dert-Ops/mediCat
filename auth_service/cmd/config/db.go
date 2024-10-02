package config

import (
	"gitlab.com/dert-ops/mediCat/mediCat-Dev.git/cmd/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	db, err := gorm.Open(postgres.Open("postgresql://admin1:admin123@postgres-master:5432/medicat_db?sslmode=disable"), &gorm.Config{})
	if err != nil {
		LogrusLogger.Fatalf("Failed to connect to database: %v", err)
	}
	if db == nil {
		LogrusLogger.Fatalf("Failed to connect to database: db is nil")
	}
	DB = db
	err = DB.AutoMigrate(&models.User{})
	if err != nil {
		LogrusLogger.Fatalf("Failed to migrate database: %v", err)
	}
	LogrusLogger.Debugf("Database connection successful.")
}
