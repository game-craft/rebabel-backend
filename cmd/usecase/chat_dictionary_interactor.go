package usecase

import (
	"docker-echo-template/cmd/domain"
)

type ChatDictionaryInteractor struct {
	ChatDictionaryRepository ChatDictionaryRepository
}

func (interactor *ChatDictionaryInteractor) ChatDictionaries() (chatDictionaries domain.ChatDictionaries, err error) {
	chatDictionaries, err = interactor.ChatDictionaryRepository.FindAll()

	return
}