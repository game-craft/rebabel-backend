package usecase

import (
	"docker-echo-template/cmd/domain"
)

type ChatDictionaryRepository interface {
	FindAll() (domain.ChatDictionaries, error)
}
