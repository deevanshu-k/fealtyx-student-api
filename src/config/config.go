package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var PORT int

func init() {
	// Load environment variables from .env file
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		log.Fatal("Error parsing PORT environment variable")
	}
	PORT = port
}
