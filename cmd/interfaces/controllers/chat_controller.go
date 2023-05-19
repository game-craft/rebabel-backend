package controllers

import (
	"os"
	"fmt"
	"bytes"
	"io/ioutil"
	"encoding/json"
	"net/http"

	"docker-echo-template/cmd/interfaces/database"
	"docker-echo-template/cmd/usecase"
)

type ChatController struct {
	Interactor usecase.ChatInteractor
}

type Chat struct {
	ID int `json:"id"`
	Word string `json:"word"`
}

type ChatResponse struct {
	Id int `json:"id"`
	WordList []string `json:"word_list"`
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

func (controller *ChatController) UpdateData() {
	body := controller.callAnalysisApi()
	controller.wordClassification(body)
}

func (controller *ChatController) callAnalysisApi() []byte {
	data, err := controller.Interactor.Chats()
	if err != nil {
		panic(err)
	}

	var chats []Chat

	for i :=0; i < len(data); i++ {
		json := Chat{
			ID: data[i].ID,
			Word: data[i].ChatsContent,
		}
		chats = append(chats, json)
	}

	host := os.Getenv("ANALYSIS_API_HOST")
	url := fmt.Sprintf("%s/analysis", host)
	headers := map[string]string{"Content-Type": "application/json"}
	jsonStr, _ := json.Marshal(chats)
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	
	return body
}

func (controller *ChatController) wordClassification(body []byte) {
	var chatResponse []ChatResponse
	json.Unmarshal(body, &chatResponse)

	fmt.Println(chatResponse)
}