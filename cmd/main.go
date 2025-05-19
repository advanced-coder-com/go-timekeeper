package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initConfig() {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Printf("No .env file found or error reading it: %v", err)
	}
}

func initDB() *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		viper.GetString("DB_HOST"),
		viper.GetString("DB_USER"),
		viper.GetString("DB_PASSWORD"),
		viper.GetString("DB_NAME"),
		viper.GetString("DB_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Cannot connect to database: %v", err)
	}

	log.Println("Database connection established")
	return db
}

func main() {
	initConfig()

	port := viper.GetString("APP_PORT")
	if port == "" {
		port = "8080"
	}

	db := initDB()
	_ = db

	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		msg := "pong"
		c.JSON(200, gin.H{"message": msg})
	})

	log.Printf("Starting server on port %s...", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
