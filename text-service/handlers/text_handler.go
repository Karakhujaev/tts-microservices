package handlers

import (
	"net/http"
	"text-service/services"

	"github.com/gin-gonic/gin"
)

func CreateText(c *gin.Context) {
	var request struct {
		Text string `json:"text" binding:"required"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	text, err := services.ProcessText(request.Text)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process text"})
		return
	}

	c.JSON(http.StatusOK, text)
}

func GetText(c *gin.Context) {
	id := c.Param("id")
	text, err := services.GetSpeechResult(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Text not found"})
		return
	}

	c.JSON(http.StatusOK, text)
}
