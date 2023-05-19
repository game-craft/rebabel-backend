package main

import (
	"docker-echo-template/cmd/infrastructure"
	"docker-echo-template/cmd/interfaces/controllers"
)

func main() {
	chatController := controllers.NewChatController(infrastructure.NewSqlHandler())
	chatController.UpdateData()
}