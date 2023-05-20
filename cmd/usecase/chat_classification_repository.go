package usecase

import (
	"docker-echo-template/cmd/domain"
)

type ChatClassificationRepository interface {
	Add(domain.ChatClassification) (domain.ChatClassification, error)
}
