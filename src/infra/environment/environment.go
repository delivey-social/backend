package environment

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var PORT string

func Load() {
	err := godotenv.Load()
	if err != nil {
		log.Panicf("Error loading .env file: %v", err)
	}

	PORT = os.Getenv("PORT")
}
