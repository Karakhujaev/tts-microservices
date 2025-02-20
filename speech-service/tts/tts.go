package tts

import (
	"fmt"
	"os"
	"os/exec"
	"speech-service/config"
	"speech-service/models"
	"speech-service/storage"

	"github.com/google/uuid"
)

func ConvertTextToSpeech(text string) (models.Speech, error) {
	id := uuid.New().String()
	audioFile := fmt.Sprintf("%s/%s.wav", config.GetStoragePath(), id)

	cmd := exec.Command("tts", "--text", text, "--out_path", audioFile)
	if err := cmd.Run(); err != nil {
		return models.Speech{}, err
	}

	if _, err := os.Stat(audioFile); os.IsNotExist(err) {
		return models.Speech{}, fmt.Errorf("audio file was not created")
	}

	audioURL := fmt.Sprintf("/audio/%s.wav", id)
	return models.Speech{ID: id, Text: text, AudioURL: audioURL}, nil
}
