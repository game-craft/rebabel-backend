package domain

import (
    "time"
)

type ChatClassifications []ChatClassification

type ChatClassification struct {
	ID int `json:"id"`
	WorldsId int `json:"world_id"`
	ChatClassificationsContent string `json:"chat_classifications_content"`
	ChatClassificationsStatus string `json:"chat_classifications_status"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
    UpdatedAt time.Time `gorm:"autoUpdateTime"`
}