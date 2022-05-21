package main

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	token, _ := os.LookupEnv("BOT_API_TOKEN")
	bot, err := tgbotapi.NewBotAPI(token)

	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)
	for update := range updates {
		if update.Message != nil { // If we got a message
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
			msg := update.Message.Text
			switch msg {
			case "/start":
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Welcome")
				bot.Send(msg)
			default:
				msg := tgbotapi.NewDeleteMessage(update.Message.Chat.ID, update.Message.MessageID)
				bot.Send(msg)

			}
		}
	}
}
