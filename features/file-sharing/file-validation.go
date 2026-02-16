package filesharing

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func ValidateFile(filetype string, fileid string, update *tgbotapi.Update) bool {

	mime := update.Message.Document.MimeType
	if mime == "exe" {
		return false
	}
	return true

}
