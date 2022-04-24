package botogoto_mainbody

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func (b *Bot) echoHandler(message *tgbotapi.Message) {
	log.Printf("[%s] %s", message.From.UserName, message.Text)
	// Создаём объект сообщения (msg = Message(chat_id, message.text)):
	msg := tgbotapi.NewMessage(message.Chat.ID, message.Text)
	// Задаём ID сообщения, которое цитируем:
	msg.ReplyToMessageID = message.MessageID
	// Отправляем сообщение с проверкой на ошибку:
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
}
