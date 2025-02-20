package models

type Speech struct {
	ID       string `json:"id"`
	Text     string `json:"text"`
	AudioURL string `json:"audio_url"`
}
