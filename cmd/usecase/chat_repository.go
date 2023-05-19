package usecase

import (
	"docker-echo-template/cmd/domain"
)

type ChatRepository interface {
	FindAll() (domain.Chats, error)
}
