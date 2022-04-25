package botogoto_mainbody

import (
	"BOTOGOTO/pkg/config"
	"flag"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

type Bot struct {
	botObj *tgbotapi.BotAPI
}

func NewBot(bot *tgbotapi.BotAPI) *Bot {
	return &Bot{botObj: bot}
}

func (b *Bot) StartBot() {
	log.Printf("Authorized on account %s", b.botObj.Self.UserName)
	b.startNotifyAdmin()

	updatesCh := b.buildUpdatesChannel()
	b.allUpdatesRoute(updatesCh)
}

func (b *Bot) buildUpdatesChannel() tgbotapi.UpdatesChannel {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := b.botObj.GetUpdatesChan(u)
	return updates
}

func (b *Bot) allUpdatesRoute(upd tgbotapi.UpdatesChannel) {
	for update := range upd {
		if update.Message.IsCommand() {
			switch update.Message.Command() {
			case config.CommandStart:
				b.startHandler(update.Message)
			default:
				continue
			}
		}
	}
}

func Token() string { // TODO
	token := flag.String(
		"token",
		"",
		"Enter bot token here",
	)
	flag.Parse()
	if *token == "" {
		log.Fatal("Missing token!")
	}
	return *token
}
