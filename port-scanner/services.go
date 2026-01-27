package portscanner

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

type Service struct {
	service string
	port    string
}

var SlicesOfservice []Service

func Services() {

	file, err := os.Open("/etc/services")
	if err != nil {
		fmt.Println(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		var serviceName, portno string

		line := scanner.Text()
		if len(line) == 0 {
			continue
		}
		if line[0] == '#' {
			continue
		}

		newline := strings.Split(line, "/")
		for _, ch := range newline[0] {

			if !unicode.IsDigit(ch) {
				serviceName += string(ch)
				continue
			}

			portno += string(ch)

		}

		SlicesOfservice = append(SlicesOfservice, Service{service: serviceName, port: portno})

	}

	//fuvk

}

func GetServiceName(portno string) string {
	for _, servicestruct := range SlicesOfservice {
		if portno == servicestruct.port {
			return servicestruct.service

		}
	}
	return ""
}
