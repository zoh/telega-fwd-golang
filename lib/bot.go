package lib

import (
	"log"
	"os"

	"gopkg.in/telegram-bot-api.v4"
)

func CreateBot() *tgbotapi.BotAPI {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_BOT_TOKEN"))
	if err != nil {
		log.Panic(err)
	}
	log.Printf("Authorized on account %s", bot.Self.UserName)

	return bot
}
