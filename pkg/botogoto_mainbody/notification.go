package botogoto_mainbody

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func (b *Bot) startNotifyAdmin() {
	msg := tgbotapi.NewMessage(adminID, "Бот запущен")
	if _, err := b.botObj.Send(msg); err != nil {
		log.Fatal(err)
	}
}

func StopNotifyAdmin(bot *tgbotapi.BotAPI) {
	signalCancel := make(chan os.Signal, 1)
	signal.Notify(signalCancel, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		for {
			s := <-signalCancel
			switch s {
			case os.Interrupt:
				fallthrough
			case syscall.SIGINT:
				fallthrough
			case syscall.SIGTERM:
				msg := tgbotapi.NewMessage(adminID, "Бот остановлен")
				if _, err := bot.Send(msg); err != nil {
					log.Fatal(err)
				}
				log.Fatal("Bot Stopped!")
			}
		}
	}()
}
