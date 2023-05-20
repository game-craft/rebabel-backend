package main

import (
	"os"
	"fmt"
	"bytes"
	"io/ioutil"
	"encoding/json"
	"net/http"

	"docker-echo-template/cmd/domain"
	"docker-echo-template/cmd/infrastructure"
	"docker-echo-template/cmd/interfaces/controllers"
)

type Chat struct {
	ID int `json:"id"`
	Word string `json:"word"`
}

type ChatResponse struct {
	Id int `json:"id"`
	WordList []string `json:"word_list"`
}

func main() {
	chatController := controllers.NewChatController(infrastructure.NewSqlHandler())
	chatDictionaryController := controllers.NewChatDictionaryController(infrastructure.NewSqlHandler())

	chats := chatController.GetChatData()
	chatResponse := callAnalysisApi(chats)
	fmt.Println(chatResponse)

	chatDictionarys := chatDictionaryController.GetChatDictionaryData()
	fmt.Println(chatDictionarys)
}

func callAnalysisApi(chats domain.Chats) []ChatResponse {
	var chatBodys []Chat

	for i :=0; i < len(chats); i++ {
		json := Chat{
			ID: chats[i].ID,
			Word: chats[i].ChatsContent,
		}
		chatBodys = append(chatBodys, json)
	}

	host := os.Getenv("ANALYSIS_API_HOST")
	url := fmt.Sprintf("%s/analysis", host)
	headers := map[string]string{"Content-Type": "application/json"}
	jsonStr, _ := json.Marshal(chatBodys)
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

	var chatResponse []ChatResponse
	json.Unmarshal(body, &chatResponse)
	
	return chatResponse
}