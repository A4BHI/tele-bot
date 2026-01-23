package main

import (
	"fmt"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	tgbot, err := tgbotapi.NewBotAPI(os.Getenv("BOT_API"))
	if err != nil {
		log.Fatal("error: ", err)
	}

	cmds := []tgbotapi.BotCommand{}
	cmds = append(cmds, tgbotapi.BotCommand{Command: "port_scanner", Description: "Enter Domain or IP adress to scan ports"},
		tgbotapi.BotCommand{Command: "ping", Description: "You can check if the bot is alive"})

	config := tgbotapi.NewSetMyCommands(cmds...)

	_, err = tgbot.Request(config)
	if err != nil {
		log.Fatal(err)
	}

	up := tgbotapi.NewUpdate(0)

	updates := tgbot.GetUpdatesChan(up)
	for updates := range updates {
		if updates.Message == nil {
			continue
		}

		if !updates.Message.IsCommand() {
			fmt.Println(updates.Message.Chat.UserName + ":" + updates.Message.Text)
			continue
		}

		switch updates.Message.Command() {
		case "port_scanner":
			arg := updates.Message.CommandArguments()

			if len(arg) < 1 {
				reply := tgbotapi.NewMessage(updates.Message.Chat.ID, "Please provide domain or ip adress of the target")
				reply.ReplyToMessageID = updates.Message.MessageID
				tgbot.Send(reply)
				break
			}

			reply := tgbotapi.NewMessage(updates.Message.Chat.ID, "Scanning please wait for a moment.")
			reply.ReplyToMessageID = updates.Message.MessageID
			tgbot.Send(reply)

		}

	}

}
