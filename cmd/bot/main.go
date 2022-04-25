package main

import (
	bot "BOTOGOTO/pkg/botogoto_mainbody"
	"BOTOGOTO/pkg/config"
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

	db, _ := config.DBConnect()
	bot.CreateSchema(db)

	client, err := tgbotapi.NewBotAPI(TOKEN)
	if err != nil {
		log.Panic(err)
	}
	client.Debug = false
	bot.StopNotifyAdmin(client)
	dp := bot.NewBot(client)
	dp.StartBot()
}
