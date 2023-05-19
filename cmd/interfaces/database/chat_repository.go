package database

import (
	"docker-echo-template/cmd/domain"
)

type ChatRepository struct {
	SqlHandler
}

func (repo *ChatRepository) FindAll() (chats domain.Chats, err error) {
	if err = repo.Find(&chats).Error; err != nil {
		return
	}

	return
}
