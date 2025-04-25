package main

import (
	"AI_Interviewer/internal/handlers"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}

	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_API"))

	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(60)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	msgHandler := handlers.NewMessageHandler(bot)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		msgHandler.HandleMessage(update.Message)
	}
}
