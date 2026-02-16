package portscanner

import (
	"fmt"
	"net"
	"strconv"
	"sync"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func ScanPort(target string, updates *tgbotapi.Update, tgbot tgbotapi.BotAPI, db *DB) {
	var wg sync.WaitGroup
	var OpenPorts []string
	for port := 1; port <= 1024; port++ {
		p := port
		wg.Add(1)
		go func() {

			defer wg.Done()
			conn, err := net.DialTimeout("tcp", target+":"+strconv.Itoa(p), 1000*time.Millisecond)
			if err != nil {
				fmt.Println(err)
				return
			}
			OpenPorts = append(OpenPorts, strconv.Itoa(port))

			conn.Close()
		}()

	}

	extraPorts := []string{"8080", "3389", "1443", "3306", "3389", "5900"}

	for _, ports := range extraPorts {
		wg.Add(1)
		address := target + ":" + ports
		go func() {
			defer wg.Done()
			conn, err := net.DialTimeout("tcp", address, 1000*time.Millisecond)
			if err != nil {
				fmt.Println(err)
				return
			}
			OpenPorts = append(OpenPorts, ports)
			conn.Close()
		}()

	}

	wg.Wait()
	stringBuilder := "\nOpen Port: "
	var temp string
	for _, port := range OpenPorts {
		name := db.LookUP(port)
		temp += stringBuilder + name + " " + port
	}

	reply := tgbotapi.NewMessage(updates.Message.Chat.ID, temp)
	reply.ReplyToMessageID = updates.Message.MessageID
	tgbot.Send(reply)
	fmt.Println(">>")
	fmt.Println(temp)
	fmt.Println(">>")
}
