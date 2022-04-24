package tg

import (
	"errors"
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/neosouler7/tgHwang/utils"
)

type botConfig struct {
	token        string
	chatId       int64
	commanderIds []int64
}

var (
	bc  *botConfig
	bot *tgbotapi.BotAPI
)

func initBot() (*botConfig, *tgbotapi.BotAPI) {
	c := utils.TgConfig()
	bc = &botConfig{
		token:        c.Token,
		chatId:       c.ChatId,
		commanderIds: c.CommanderIds,
	}

	botPointer, err := tgbotapi.NewBotAPI(bc.token)
	if err != nil {
		log.Panic(errors.New("Failed initializing bot"))
	}
	// botPointer.Debug = true

	bot = botPointer
	log.Printf("Authorized on account %s", bot.Self.UserName)
	return bc, bot
}

func getNonCommanderMsg(update tgbotapi.Update) string {
	messageFromID := update.Message.From.ID
	userName := update.Message.From.UserName
	firstName := update.Message.From.FirstName
	lastName := update.Message.From.LastName
	return fmt.Sprintf("Trying to approach!\n\nID: %d\nUser: %s\nFirst: %s\nLast: %s\n", messageFromID, userName, firstName, lastName)
}

func listenMsg(bc *botConfig, bot *tgbotapi.BotAPI) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)
	for update := range updates {
		if update.Message == nil {
			continue
		}

		if !update.Message.IsCommand() { // server ignores non-command messages
			continue
		}

		commander := utils.Contains(bc.commanderIds, update.Message.From.ID)

		var msgText string
		switch update.Message.Command() {
		case "whoami":
			msgText = whoami(update)
		case "whereami":
			msgText = whereami(update)
		case "hi":
			if commander {
				msgText = hi()
			} else { // someone is trying to approach!
				msgText = getNonCommanderMsg(update)
			}
		default:
			msgText = "Wrong command :("
		}

		SendMsg(bc, bot, msgText)
	}
}

func SendMsg(bc *botConfig, bot *tgbotapi.BotAPI, msgText string) {
	msg := tgbotapi.NewMessage(bc.chatId, msgText)
	if _, err := bot.Send(msg); err != nil {
		log.Panic(err)
	}
}

func Start() {
	bc, bot := initBot()
	listenMsg(bc, bot)
}
