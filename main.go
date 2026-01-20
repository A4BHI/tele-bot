package main

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

func main() {
	var wg sync.WaitGroup
	for port := 1; port <= 1024; port++ {
		wg.Add(1)
		go ScanPort("", port, &wg)
		continue

	}
	wg.Wait()
}
