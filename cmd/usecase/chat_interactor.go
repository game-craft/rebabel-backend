package usecase

import (
	"docker-echo-template/cmd/domain"
)

type ChatInteractor struct {
	ChatRepository ChatRepository
}

func (interactor *ChatInteractor) Chats() (chats domain.Chats, err error) {
	chats, err = interactor.ChatRepository.FindAll()

	return
}