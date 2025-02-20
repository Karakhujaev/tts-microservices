package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"text-service/config"
	"text-service/db"
	"text-service/handlers"
)

func main() {
	config.LoadConfig()

	db.InitDB()

	router := gin.Default()
	router.POST("/api/text", handlers.CreateText)
	router.GET("/api/speech/:id", handlers.GetText)

	port := config.GetPort()
	log.Printf("Text Service running on port %s", port)
	router.Run(":" + port)
}
