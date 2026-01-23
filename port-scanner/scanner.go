package portscanner

import (
	"fmt"
	"net"
	"strconv"
	"sync"
	"time"
)

func ScanPort(target string, port int, wg *sync.WaitGroup) {
	defer wg.Done()
	address := target + ":" + strconv.Itoa(port)

	conn, err := net.DialTimeout("tcp", address, 500*time.Millisecond)
	if err != nil {

		return
	}

	fmt.Println("Open port:", port)
	conn.Close()

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
