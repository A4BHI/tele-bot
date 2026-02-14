package portscanner

import (
	"bufio"
	"log"
	"os"
	"strings"
)


type  ServicesAndProtocols struct{
	Name string
	Protocol string
}
type DB struct {
	Port map[string]string
}

func (db *DB) LookUP(portno string) (serivcename string) {
	if name, ok := db.Port[portno]; ok {
		return name
	}

	return "unknown"
}

func LoadService(path string) (*DB, error) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	db := &DB{
		Port: make(map[string]string),
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
		ports:= 

		// fmt.Println(field[0], field[1])

		db.Port[field[0]] = field[1]

	}

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
