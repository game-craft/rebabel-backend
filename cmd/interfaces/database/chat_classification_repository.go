package database

import (
	"docker-echo-template/cmd/domain"
)

type ChatClassificationRepository struct {
	SqlHandler
}

func (repo *ChatClassificationRepository) Add(u domain.ChatClassification) (chatClassification domain.ChatClassification, err error) {
	if err = repo.Create(&u).Error; err != nil {
		return
	}
	chatClassification = u

	return
}