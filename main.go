package main

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

func ScanPort(target string) {
	for port := 1; port <= 1024; port++ {
		address := target + ":" + strconv.Itoa(port)

		conn, err := net.DialTimeout("tcp", address, 500*time.Millisecond)
		if err != nil {
			continue
		}

		fmt.Println("Open port:", port)
		conn.Close()
	}
}

func main() {

}
