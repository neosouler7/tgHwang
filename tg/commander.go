package tg

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func whoami(update tgbotapi.Update) string {
	messageFromID := update.Message.From.ID
	firstName := update.Message.From.FirstName
	lastName := update.Message.From.LastName
	return fmt.Sprintf("Hi %s %s !\nYou're ID is %d", firstName, lastName, messageFromID)
}

func whereami(update tgbotapi.Update) string {
	messageFromID := update.Message.From.ID
	chatId := update.Message.Chat.ID
	return fmt.Sprintf("Hi %d !\nYou're now in chat %d", messageFromID, chatId)
}

func hi() string {
	return "hi"
}
