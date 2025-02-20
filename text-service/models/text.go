package models

import "time"

type Text struct {
	ID        string    `json:"id" gorm:"primaryKey"`
	Text      string    `json:"text"`
	Status    string    `json:"status"`
	AudioURL  string    `json:"audio_url"`
	CreatedAt time.Time `json:"created_at"`
}
