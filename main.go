package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"go-crud-api/config"
	"go-crud-api/routes"
)

func main() {
	// โหลดไฟล์ .env
	err := godotenv.Load(".env.dev")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// เชื่อมต่อกับฐานข้อมูล
	db, err := config.ConnectDatabase()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer config.CloseDatabase(db)

	router := gin.Default()
	// log request
	router.Use(config.Middleware())

	// ตั้งค่า route
	routes.UserRoute(router)

	version := os.Getenv("VERSION")
	port := os.Getenv("PORT")
	log.Printf("[GO-CRUD-API:%v] Service start at localhost:%v\n", version, port)

	// เริ่มต้น Gin service
	addr := fmt.Sprintf("0.0.0.0:%s", port)
	if err := router.Run(addr); err != nil {
		log.Fatalf("Failed to start the server: %v", err)
	}
}
