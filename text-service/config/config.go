package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadConfig() {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: No .env file found, using default values")
	}
}

func GetPort() string {
	port := os.Getenv("TEXT_SERVICE_PORT")
	if port == "" {
		port = "8080"
	}
	return port
}

func GetDBURL() string {
	return os.Getenv("DATABASE_URL")
}

func GetSpeechServiceURL() string {
	return os.Getenv("SPEECH_SERVICE_URL")
}
