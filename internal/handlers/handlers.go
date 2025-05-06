package handlers

import (
	"AI_Interviewer/internal/ai"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

type UserSession struct {
	ChatID      int64
	CurrentStep int
	Answers     []string
}

type MessageHandler struct {
	botAPI *tgbotapi.BotAPI
}

//var sessions = make(map[int64]*UserSession)

func NewMessageHandler(botAPI *tgbotapi.BotAPI) *MessageHandler {
	return &MessageHandler{botAPI: botAPI}
}

func (h *MessageHandler) sendMessage(chatID int64, text string) {
	if text == "" {
		log.Println("Warning: Attempted to send empty message. Skipping.")
		return
	}

	msg := tgbotapi.NewMessage(chatID, text)
	if _, err := h.botAPI.Send(msg); err != nil {
		panic(err)
	}
}

func (h *MessageHandler) HandleMessage(msg *tgbotapi.Message) {
	//session, ok := sessions[msg.Chat.ID]
	//if !ok {
	//	session := &UserSession{
	//		ChatID:      msg.Chat.ID,
	//		CurrentStep: 0,
	//		Answers:     []string{},
	//	}
	//	sessions[msg.Chat.ID] = session
	//}

	switch msg.Text {
	case "/start":
		res := "Добро пожаловать в AI Interviewer! Давайте начнём интервью. Введите /next для следующего вопроса."
		h.sendMessage(msg.Chat.ID, res)

	case "/help":
		res := "Вот доступные команды:\n" +
			"/start — начать интервью\n" +
			"/next — следующий вопрос\n" +
			"/stop — завершить интервью\n" +
			"/restart — начать интервью заново\n" +
			"/score — получить оценку по завершению интервью\n" +
			"/status — узнать текущий статус интервью\n" +
			"Введите команду, чтобы продолжить."
		h.sendMessage(msg.Chat.ID, res)

	case "/stop":
		res := "Интервью завершено. Если хотите начать заново, введите /start."
		h.sendMessage(msg.Chat.ID, res)

	case "/next":
		res := "Вот ваш следующий вопрос: Как вы решаете сложные задачи?"
		h.sendMessage(msg.Chat.ID, res)

	case "/score":
		res := "Вы можете получить свой результат в конце интервью, когда все вопросы будут пройдены. Введите /next для следующего вопроса."
		h.sendMessage(msg.Chat.ID, res)

	case "/restart":
		res := "Интервью начинается заново. Введите /next для первого вопроса."
		h.sendMessage(msg.Chat.ID, res)

	case "/status":
		res := "Вы еще не начали интервью. Введите /start для начала."
		h.sendMessage(msg.Chat.ID, res)

	default:
		res := ai.GetAIReply()
		h.sendMessage(msg.Chat.ID, res)
	}
}
