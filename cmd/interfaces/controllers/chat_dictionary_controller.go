package controllers

import (
	"docker-echo-template/cmd/domain"
	"docker-echo-template/cmd/interfaces/database"
	"docker-echo-template/cmd/usecase"
)

type ChatDictionaryController struct {
	Interactor usecase.ChatDictionaryInteractor
}

func NewChatDictionaryController(sqlHandler database.SqlHandler) *ChatDictionaryController {
	return &ChatDictionaryController{
		Interactor: usecase.ChatDictionaryInteractor{
			ChatDictionaryRepository: &database.ChatDictionaryRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *ChatDictionaryController) GetChatDictionaryData() (chatDictionaries domain.ChatDictionaries) {
	chatDictionaries, err := controller.Interactor.ChatDictionaries()
	if err != nil {
		panic(err)
	}

	return
}