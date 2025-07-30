package config

import (
	"fmt"
	"go-crud-api/models"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() (*gorm.DB, error) {

	// ใช้ environment variable
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	sslMode := os.Getenv("DB_SSLMODE") // เช่น disable, require

	// connection string ของ PostgreSQL
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", dbHost, dbUser, dbPass, dbName, dbPort, sslMode)

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database!:%v", err)
		return nil, err
	}
	sqlDB, err := database.DB()
	if err != nil {
		log.Fatalf("failed to get generic database object:%v", err)
	}
	// Connection pooling
	sqlDB.SetMaxOpenConns(20)
	sqlDB.SetMaxIdleConns(20)
	sqlDB.SetConnMaxLifetime(time.Hour)

	// AutoMigrate model
	err = database.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("Failed to migrate:%v", err)
	}
	DB = database

	return DB, nil
}

func CloseDatabase(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get database connection:%v", err)
	}
	sqlDB.Close()
}
