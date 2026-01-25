package portscanner

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

type service struct {
	service string
	port    string
}

var slicesOfservice []service

func Services() {

	file, err := os.Open("/etc/services")
	if err != nil {
		fmt.Println(err)
	}
	var serviceName string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}
		if line[0] == '#' {
			continue
		}

		for _, ch := range line {

			if !unicode.IsSpace(ch) && !unicode.IsNumber(ch) {
				serviceName += string(ch)
			}

			if unicode.IsSpace(ch) {
				continue
			}

		}

		fmt.Println(serviceName)
		break

	}

	//fuvk
}
