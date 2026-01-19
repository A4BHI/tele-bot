package main

import (
	"net"
	"strconv"
)

func ScanPort(target string) {
	for port := 1; port <= 1024; port++ {
		conn, err := net.Dial("tcp", target+":"+strconv.Itoa(port))
	}
}

func main() {
	ScanPort("nil")
}
