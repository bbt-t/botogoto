package botogoto_mainbody

import (
	"BOTOGOTO/pkg/config"
	"BOTOGOTO/pkg/logging"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func (b *Bot) startNotifyAdmin() {
	msg := tgbotapi.NewMessage(config.AdminID, "Бот запущен")
	if _, err := b.botObj.Send(msg); err != nil {
		log.Fatal(err)
	}
	logger := logging.GetLogger()
	logger.Info("Bot started!")
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
				msg := tgbotapi.NewMessage(config.AdminID, "Бот остановлен")
				if _, err := bot.Send(msg); err != nil {
					log.Fatal(err)
				}
				logger := logging.GetLogger()
				logger.Info("Bot Stopped!")
				os.Exit(0)
			}
		}
	}()
}
