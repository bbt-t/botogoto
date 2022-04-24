package botogoto_mainbody

import (
	"flag"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

const commandStart = "start"
const adminID = 2018211211

type Bot struct {
	botObj *tgbotapi.BotAPI
}

// Конструктор

func NewBot(bot *tgbotapi.BotAPI) *Bot {
	// Присваивает атрибуту "botObj" значение принятой ссылки (указателя) на объект бота "*tgbotapi.BotAPI"
	return &Bot{botObj: bot}
}

func (b *Bot) StartBot() {
	// Вывод лога с именем бота:
	log.Printf("Authorized on account %s", b.botObj.Self.UserName)
	// Отправка сообщения о старте бота:
	b.startNotifyAdmin()

	updatesCh := b.buildUpdatesChannel()
	b.allUpdatesRoute(updatesCh)
}

func (b *Bot) buildUpdatesChannel() tgbotapi.UpdatesChannel {
	// Указываем таймаут:
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	// Создаём канал (горутину) с таймаутом, в который будем получать значения от телеграмма:
	updates := b.botObj.GetUpdatesChan(u)
	return updates
}

func (b *Bot) allUpdatesRoute(upd tgbotapi.UpdatesChannel) {
	// Читаем значения из созданного канала в цикле (блокируется дальнейший код):
	for update := range upd {
		if update.Message.IsCommand() { // Если сообщение получено от пользователя (игнор остальных).
			switch update.Message.Command() {
			case commandStart:
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
