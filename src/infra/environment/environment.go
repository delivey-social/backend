package environment

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	PORT,
	DB_HOST,
	DB_PORT,
	DB_USER,
	DB_PASS,
	DB_NAME string
)

func Load() {
	err := godotenv.Load()
	if err != nil {
		log.Panicf("Error loading .env file: %v", err)
	}

	PORT = os.Getenv("PORT")

	DB_HOST = os.Getenv("DB_HOST")
	DB_PORT = os.Getenv("DB_PORT")
	DB_USER = os.Getenv("DB_USER")
	DB_PASS = os.Getenv("DB_PASS")
	DB_NAME = os.Getenv("DB_NAME")
}
