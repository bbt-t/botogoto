package main

import (
	botogoto "BOTOGOTO/pkg/botogoto_mainbody"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	TOKEN := os.Getenv("TOKEN")
	// создаём объект бота и обрабатываем возможную ошибку
	client, err := tgbotapi.NewBotAPI(TOKEN)
	if err != nil {
		log.Panic(err)
	}
	// Для дебага (вывод логов):
	client.Debug = false
	// Отправляем сообщение о завершении работы:
	botogoto.StopNotifyAdmin(client)
	// Передаём в конструктор объект бота:
	dispatcher := botogoto.NewBot(client)
	// Стартуем:
	dispatcher.StartBot()
}
