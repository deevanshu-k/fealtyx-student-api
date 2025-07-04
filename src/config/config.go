package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var PORT int
var OLLAMA_MODEL string
var SUMMARIZE_STUDENT_PROMPT string
var OLLAMA_GENERATE_URL string

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

	OLLAMA_MODEL = os.Getenv("OLLAMA_MODEL")
	if OLLAMA_MODEL == "" {
		log.Fatal("OLLAMA_MODEL environment variable is not set")
	}

	SUMMARIZE_STUDENT_PROMPT = "Summarize this student as if you're writing a brief for a teacher, Summary should be in plain text without formating and should contain only the most important information."

	OLLAMA_GENERATE_URL = os.Getenv("OLLAMA_GENERATE_URL")
	if OLLAMA_GENERATE_URL == "" {
		log.Fatal("OLLAMA_GENERATE_URL environment variable is not set")
	}
}
