## info
Простенький Телеграм бот на библиотеке [telegram-bot-api](http://github.com/go-telegram-bot-api/telegram-bot-api)

***

**Может:**

- `поприветствовать при старте бота`
- `выслать сообщение администратору при старте и завершении работы бота`
  

_Функционал будет расширяться._

## usage
- Заменить [константу](https://github.com/bbt-t/botogoto/blob/fa4b7ebeefda89dc2fc971035e7d633766316e8a/pkg/botogoto_mainbody/bot.go#L10) с `telegram_id` админа
- Создать файл в корне с именем `.env` и прописать в нём токен бота: 

>TOKEN=ЗДЕСЬ_ТВОЙ_ТОКЕН

- Стартуем через команды:

>make build && make run
