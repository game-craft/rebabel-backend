package database

import (
	"docker-echo-template/cmd/domain"
)

type ChatDictionaryRepository struct {
	SqlHandler
}

func (repo *ChatDictionaryRepository) FindAll() (chatDictionaries domain.ChatDictionaries, err error) {
	if err = repo.Find(&chatDictionaries).Error; err != nil {
		return
	}

	return
}
