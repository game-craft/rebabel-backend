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
	WorldsId int `json:"worlds_id"`
	Word string `json:"word"`
}

type ChatResponse struct {
	WorldsId int `json:"worlds_id"`
	WordList []string `json:"word_list"`
}

func main() {
	chatController := controllers.NewChatController(infrastructure.NewSqlHandler())
	chatDictionaryController := controllers.NewChatDictionaryController(infrastructure.NewSqlHandler())

	chats := chatController.GetChatData()
	chatResponse := callAnalysisApi(chats)
	chatDictionarys := chatDictionaryController.GetChatDictionaryData()
	chatDictionariesGood, chatDictionariesBad := createChatChatClassificationList(chatDictionarys)
	chatClassifications := createChatChatClassificationQuery(chatResponse, chatDictionariesGood, chatDictionariesBad)
	fmt.Println(chatClassifications)
}

func callAnalysisApi(chats domain.Chats) []ChatResponse {
	var chatBodys []Chat

	for i :=0; i < len(chats); i++ {
		json := Chat{
			WorldsId: chats[i].WorldsId,
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

func createChatChatClassificationList(chatDictionarys domain.ChatDictionaries) (chatDictionariesGood []string, chatDictionariesBad []string) {
	for i :=0; i < len(chatDictionarys); i++ {
		if chatDictionarys[i].ChatDictionariesStatus == "status:Good" {
			chatDictionariesGood = append(chatDictionariesGood, chatDictionarys[i].ChatDictionariesContent)
		} else if chatDictionarys[i].ChatDictionariesStatus == "status:Bad" {
			chatDictionariesBad = append(chatDictionariesBad, chatDictionarys[i].ChatDictionariesContent)
		} else {
			fmt.Println("Could not classify")
		}
	}

	return chatDictionariesGood, chatDictionariesBad
}

func createChatChatClassificationQuery(chatResponse []ChatResponse, chatDictionariesGood []string, chatDictionariesBad []string) (chatClassifications domain.ChatClassifications) {
	for i :=0; i < len(chatResponse); i++ {
		for j :=0; j < len(chatResponse[i].WordList); j++ {
			if searchContains(chatDictionariesGood, chatResponse[i].WordList[j]) {
				data := domain.ChatClassification{
					WorldsId: chatResponse[i].WorldsId,
					ChatClassificationsContent: chatResponse[i].WordList[j],
					ChatClassificationsStatus: "status:Good",
				}
				chatClassifications = append(chatClassifications, data)
			}
			if searchContains(chatDictionariesBad, chatResponse[i].WordList[j]) {
				data := domain.ChatClassification{
					WorldsId: chatResponse[i].WorldsId,
					ChatClassificationsContent: chatResponse[i].WordList[j],
					ChatClassificationsStatus: "status:Bad",
				}
				chatClassifications = append(chatClassifications, data)
			}
		}
	}

	return chatClassifications
}

func searchContains(arr []string, target string) bool {
    for _, element := range arr {
        if element == target {
            return true
        }
    }

    return false
}