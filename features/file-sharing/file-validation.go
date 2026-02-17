package filesharing

import (
	"fmt"
	"net/http"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func ValidateFile(fileid string, update *tgbotapi.Update, tgbot tgbotapi.BotAPI) bool {

	switch {
	case update.Message.Document != nil:
		mime := update.Message.Document.MimeType
		fmt.Println(mime)
		if mime == "image/jpeg" {
			return false
		}
		fb := tgbotapi.FileConfig{
			FileID: fileid,
		}

		file, _ := tgbot.GetFile(fb)

		url := file.Link(tgbot.Token)
		http.NewRequest()
		fmt.Println(url)
		return true

	}

	return false

}
