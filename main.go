package main

import (
	file_sharing "bot/features/file-sharing"
	portscanner "bot/features/port-scanner"
	"fmt"
	"log"
	"os"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

type PendingUpload struct {
	Fileid string
	State  string
}

var userstate = make(map[int64]PendingUpload)

func main() {
	// var db *portscanner.DB
	// ch := make(chan *portscanner.DB)
	// go func() {
	// 	db, err := portscanner.LoadService("/etc/services")
	// 	if err != nil {
	// 		fmt.Println("Error loadingservices:", err)
	// 		return
	// 	}

	// 	ch <- db
	// }() NOT NECESSARY IN HERE
	db, err := portscanner.LoadService("/etc/services")
	if err != nil {
		fmt.Println("Error loadingservices:", err)
		return
	}

	godotenv.Load()

	tgbot, err := tgbotapi.NewBotAPI(os.Getenv("BOT_API"))
	if err != nil {
		log.Fatal("error: ", err)
	}

	cmds := []tgbotapi.BotCommand{}
	cmds = append(cmds,
		tgbotapi.BotCommand{
			Command:     "port_scanner",
			Description: "Enter Domain or IP adress to scan ports"},
		tgbotapi.BotCommand{
			Command:     "ping",
			Description: "You can check if the bot is alive"})

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

		if pendingUpload, exist := userstate[updates.Message.Chat.ID]; exist {
			if pendingUpload.State == "Awaiting_Response" {
				if strings.ToLower(updates.Message.Text) == "yes" {
					userstate[updates.Message.Chat.ID] = PendingUpload{
						Fileid: pendingUpload.Fileid,
						State:  "Waiting_For_Password",
					}
					reply := tgbotapi.NewMessage(updates.Message.Chat.ID, "Enter your password.")
					reply.ReplyToMessageID = updates.Message.MessageID
					tgbot.Send(reply)

				} else if strings.ToLower(updates.Message.Text) == "no" {

					reply := tgbotapi.NewMessage(updates.Message.Chat.ID, "File Stored Without Password")
					reply.ReplyToMessageID = updates.Message.MessageID
					tgbot.Send(reply)
				}
			}
		}

		if !updates.Message.IsCommand() {
			switch {
			case updates.Message.Document != nil:
				fileid := updates.Message.Document.FileID
				if !file_sharing.ValidateFile(fileid, &updates, *tgbot) {
					reply := tgbotapi.NewMessage(updates.Message.Chat.ID, "exe file extensions are not supported.")
					reply.ReplyToMessageID = updates.Message.MessageID
					tgbot.Send(reply)
					continue

				}

				reply := tgbotapi.NewMessage(updates.Message.Chat.ID, "With or Without Password reply yes or no")
				reply.ReplyToMessageID = updates.Message.MessageID
				tgbot.Send(reply)

				userstate[updates.Message.Chat.ID] = PendingUpload{
					Fileid: updates.Message.Document.FileID,
					State:  "Awaiting_Response",
				}

			}
			continue
		}

		//Commands Area
		switch updates.Message.Command() {
		case "port_scanner":
			// if db == nil {
			// 	fmt.Println("db is nil now ")
			// 	db = <-ch
			// }
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

			portscanner.ScanPort(arg, &updates, *tgbot, db)

		}

	}

}
