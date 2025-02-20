package storage

import (
	"log"
	"os"
	"speech-service/config"
)

func InitStorage() {
	path := config.GetStoragePath()
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			log.Fatal("Failed to create audio storage directory:", err)
		}
	}
}
