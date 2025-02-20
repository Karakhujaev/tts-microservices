package handlers

import (
	"net/http"
	"speech-service/tts"

	"github.com/gin-gonic/gin"
)

func ProcessSpeech(c *gin.Context) {
	var request struct {
		ID   string `json:"id" binding:"required"`
		Text string `json:"text" binding:"required"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	speech, err := tts.ConvertTextToSpeech(request.Text)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate speech"})
		return
	}

	c.JSON(http.StatusOK, speech)
}
