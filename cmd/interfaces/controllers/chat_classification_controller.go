package controllers

import (
	"docker-echo-template/cmd/domain"
	"docker-echo-template/cmd/interfaces/database"
	"docker-echo-template/cmd/usecase"
)

type ChatClassificationController struct {
	Interactor usecase.ChatClassificationInteractor
}

func NewChatClassificationController(sqlHandler database.SqlHandler) *ChatClassificationController {
	return &ChatClassificationController{
		Interactor: usecase.ChatClassificationInteractor{
			ChatClassificationRepository: &database.ChatClassificationRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *ChatClassificationController) CreateChatClassificationData(chatClassification domain.ChatClassification) {
	controller.Interactor.Add(chatClassification)
}