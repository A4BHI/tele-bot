package portscanner

import (
	"bufio"
	"os"
	"strings"
)

//	type ServicesAndProtocols struct {
//		NameOfService string
//		Protocol      string
//	}
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
		return &DB{}, err
	}
	defer file.Close()
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
		ports := strings.Split(field[1], "/")

		if ports[1] != "tcp" {
			continue
		}
		db.Port[ports[0]] = field[0]

	}

	return db, err
}
