package portscanner

import (
	"fmt"
	"net"
	"strconv"
	"sync"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func ScanPort(target string, updates *tgbotapi.Update) {
	var OpenPorts []string
	for port := 1; port <= 1024; port++ {

		go func() {
			conn, err := net.DialTimeout("tcp", target+":"+strconv.Itoa(port), 500*time.Millisecond)
			if err != nil {

				return
			}
			OpenPorts = append(OpenPorts, strconv.Itoa(port))

			conn.Close()
		}()

	}

}

func ScanRestOfThePorts(target string, ports []string, wg *sync.WaitGroup) {
	defer wg.Done()
	for _, ports := range ports {
		address := target + ":" + ports
		conn, err := net.DialTimeout("tcp", address, 1000*time.Millisecond)
		if err != nil {
			continue
		}
		fmt.Println("Open Port: ", ports)
		conn.Close()
	}
}
