package main

import (
	"fmt"
	"net"
	"strconv"
)

func ScanPort(target string) {
	for port := 1; port <= 1024; port++ {
		conn, err := net.Dial("tcp", target+":"+strconv.Itoa(port))
		if err != nil {
			fmt.Println("Closed Port: ", port)
			continue
		}

		defer conn.Close()
		fmt.Println("Port is open:", port)
	}
}

func main() {
	ScanPort("nil")
}
