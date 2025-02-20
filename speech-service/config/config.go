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
	port := os.Getenv("SPEECH_SERVICE_PORT")
	if port == "" {
		port = "8081"
	}
	return port
}

func GetStoragePath() string {
	path := os.Getenv("AUDIO_STORAGE_PATH")
	if path == "" {
		path = "./audio"
	}
	return path
}
