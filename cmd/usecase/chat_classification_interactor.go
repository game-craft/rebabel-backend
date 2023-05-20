package usecase

import (
	"docker-echo-template/cmd/domain"
)

type ChatClassificationInteractor struct {
	ChatClassificationRepository ChatClassificationRepository
}

func (interactor *ChatClassificationInteractor) Add(u domain.ChatClassification) (chatClassification domain.ChatClassification, err error) {
	chatClassification, err = interactor.ChatClassificationRepository.Add(u)

	return
}