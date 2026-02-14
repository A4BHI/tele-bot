package portscanner

import (
	"bufio"
	"log"
	"os"
	"strings"
)

type ServicesAndProtocols struct {
	NameOfService string
	Protocol      string
}
type DB struct {
	Port map[string]ServicesAndProtocols
}

func (db *DB) LookUP(portno string) (serivcename string, protocol string) {
	if name, ok := db.Port[portno]; ok {
		return name.NameOfService, name.Protocol
	}

	return "unknown", "unknown"
}

func LoadService(path string) (*DB, error) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	db := &DB{
		Port: make(map[string]ServicesAndProtocols),
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		lines := scanner.Text()
		if len(lines) == 0 {
			continue
		}
		if lines[0] == '#' {
			continue
		}

		field := strings.Fields(lines)
		ports := strings.Split(field[1], "/")

		// fmt.Println(field[0], field[1])
		if ports[1] != "tcp" {
			continue
		}
		db.Port[ports[0]] = ServicesAndProtocols{
			NameOfService: field[0],
			Protocol:      ports[1],
		}

	}

	// fmt.Println(db)

	return db, err
}

// var SlicesOfservice []Service

// func Services() {

// 	file, err := os.Open("/etc/services")
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	defer file.Close()

// 	scanner := bufio.NewScanner(file)

// 	for scanner.Scan() {

// 		var serviceName, portno string

// 		line := scanner.Text()
// 		if len(line) == 0 {
// 			continue
// 		}
// 		if line[0] == '#' {
// 			continue
// 		}

// 		newline := strings.Split(line, "/")
// 		for _, ch := range newline[0] {

// 			if !unicode.IsDigit(ch) && !unicode.IsSpace(ch) {
// 				serviceName += string(ch)
// 				continue
// 			}
// 			if unicode.IsDigit(ch) {
// 				portno += string(ch)
// 			}

// 		}

// 		SlicesOfservice = append(SlicesOfservice, Service{service: serviceName, port: portno})

// 	}

// 	//changing the design

// }

// func GetServiceName(portno string) string {
// 	for _, servicestruct := range SlicesOfservice {
// 		if portno == servicestruct.port {
// 			return servicestruct.service

// 		}
// 	}
// 	return "unknown"
// }
