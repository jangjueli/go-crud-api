package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBUrl string
}

func Load() *Config {

	// load .env ถ้ามี
	if err := godotenv.Load("../.env.dev"); err != nil {
		log.Println("No .env file found, using system environment")
	}

	return &Config{
		DBUrl: os.Getenv("DATABASE_URL"), // postgres://user:pass@localhost:5432/mydb
	}
}
