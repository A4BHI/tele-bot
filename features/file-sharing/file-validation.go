package filesharing

import (
	"fmt"
	"io"
	"net/http"
	"os"

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

		resp, err := http.Get(url)
		if err != nil {
			fmt.Println(err)
		}
		files, err := os.Create(update.Message.Document.FileName)
		io.Copy(files, resp.Body)
		fmt.Println(url)
		return true

	}

	return false

}
