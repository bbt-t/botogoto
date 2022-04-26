package botogoto_mainbody

import (
	"BOTOGOTO/pkg/config"
	"BOTOGOTO/pkg/logging"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"time"
)

func (b *Bot) echoHandler(message *tgbotapi.Message) {
	logger := logging.GetLogger()
	logger.Infof("[%s] %s", message.From.UserName, message.Text)

	msg := tgbotapi.NewMessage(message.Chat.ID, message.Text)
	msg.ReplyToMessageID = message.MessageID
	if _, err := b.botObj.Send(msg); err != nil {
		log.Fatal(err)
	}
}

func (b *Bot) startHandler(message *tgbotapi.Message) {
	text := fmt.Sprintf("Привет %s !", message.From.FirstName)

	msg := tgbotapi.NewMessage(message.Chat.ID, text)
	if _, err := b.botObj.Send(msg); err != nil {
		log.Fatal(err)
	}
	db, _ := config.DBConnect()
	defer Add(db, UserSchema{time.Now(), message.From.FirstName, message.From.ID})
}
