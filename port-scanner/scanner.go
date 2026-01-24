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
	var wg sync.WaitGroup
	var OpenPorts []string
	for port := 1; port <= 1024; port++ {

		go func() {
			wg.Add(1)
			defer wg.Done()
			conn, err := net.DialTimeout("tcp", target+":"+strconv.Itoa(port), 500*time.Millisecond)
			if err != nil {
				fmt.Println(err)
				return
			}
			OpenPorts = append(OpenPorts, strconv.Itoa(port))

			conn.Close()
		}()

	}

	extraPorts := []string{"8080", "3389", "1443", "3306", "3389", "5900"}
	go func() {
		wg.Add(1)
		wg.Done()
		for _, ports := range extraPorts {
			address := target + ":" + ports
			conn, err := net.DialTimeout("tcp", address, 1000*time.Millisecond)
			if err != nil {
				fmt.Println(err)
				continue
			}
			OpenPorts = append(OpenPorts, ports)
			conn.Close()
		}
	}()
	wg.Wait()
	fmt.Println(OpenPorts)
}
