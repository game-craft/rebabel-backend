package controllers

import (
	"docker-echo-template/cmd/domain"
	"docker-echo-template/cmd/interfaces/database"
	"docker-echo-template/cmd/usecase"
)

type ChatController struct {
	Interactor usecase.ChatInteractor
}

func NewChatController(sqlHandler database.SqlHandler) *ChatController {
	return &ChatController{
		Interactor: usecase.ChatInteractor{
			ChatRepository: &database.ChatRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *ChatController) GetChatData() (chats domain.Chats) {
	chats, err := controller.Interactor.Chats()
	if err != nil {
		panic(err)
	}

	return
}