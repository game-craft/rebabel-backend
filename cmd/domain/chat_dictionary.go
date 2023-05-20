package domain

import (
    "time"
)

type ChatDictionaries []ChatDictionary

type ChatDictionary struct {
	ID int `json:"id"`
	ChatDictionariesContent string `json:"chat_dictionaries_content"`
	ChatDictionariesStatus string `json:"chat_dictionaries_status"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
    UpdatedAt time.Time `gorm:"autoUpdateTime"`
}