package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"speech-service/config"
	"speech-service/handlers"
	"speech-service/storage"
)

func main() {
	config.LoadConfig()

	storage.InitStorage()

	router := gin.Default()
	router.POST("/process", handlers.ProcessSpeech)

	port := config.GetPort()
	log.Printf("Speech Service running on port %s", port)
	router.Run(":" + port)
}
