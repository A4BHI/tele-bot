package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
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
		conn, err := net.DialTimeout("tcp", address, 500*time.Millisecond)
		if err != nil {
			continue
		}
		fmt.Println("Open Port: ", ports)
		conn.Close()
	}
}

func main() {
	s := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter the target ip or domain: ")
	s.Scan()

	in := s.Text()
	// var ports []string
	ports := []string{"8080", "3389", "1443", "3306", "3389", "5900"}

	var wg sync.WaitGroup
	for port := 1; port <= 1024; port++ {
		wg.Add(1)
		go ScanPort(in, port, &wg)

		continue

	}
	wg.Add(1)
	go ScanRestOfThePorts(in, ports, &wg)
	wg.Wait()

}
