package handlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type MessageHandler struct {
	botAPI *tgbotapi.BotAPI
}

func NewMessageHandler(botAPI *tgbotapi.BotAPI) *MessageHandler {
	return &MessageHandler{botAPI: botAPI}
}

func (h *MessageHandler) sendMessage(chatID int64, text string) {
	msg := tgbotapi.NewMessage(chatID, text)
	if _, err := h.botAPI.Send(msg); err != nil {
		panic(err)
	}
}
func (h *MessageHandler) HandleMessage(msg *tgbotapi.Message) {
	switch msg.Text {
	case "/start":
		res := "Добро пожаловать в AI Interviewer! Давайте начнём интервью. Введите /next для следующего вопроса."
		h.sendMessage(msg.Chat.ID, res)
	}
}
