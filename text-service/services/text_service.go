package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"text-service/db"
	"text-service/models"

	"github.com/google/uuid"
	"github.com/spf13/viper"
)

func ProcessText(text string) (models.Text, error) {
	textID := uuid.New().String()
	newText := models.Text{
		ID:     textID,
		Text:   text,
		Status: "processing",
	}

	if err := db.DB.Create(&newText).Error; err != nil {
		return models.Text{}, err
	}

	payload := map[string]string{"id": textID, "text": text}
	jsonData, _ := json.Marshal(payload)
	speechServiceURL := viper.GetString("SPEECH_SERVICE_URL") + "/process"

	resp, err := http.Post(speechServiceURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil || resp.StatusCode != http.StatusOK {
		return models.Text{}, errors.New("failed to forward text to speech service")
	}

	return newText, nil
}

func GetSpeechResult(id string) (models.Text, error) {
	var text models.Text
	if err := db.DB.First(&text, "id = ?", id).Error; err != nil {
		return models.Text{}, errors.New("text not found")
	}
	return text, nil
}
