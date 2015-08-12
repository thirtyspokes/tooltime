package main

import (
	"fmt"
)

func saveResult(result string) (err error) {
	conn, err := getDatabaseConnection()

	if err == nil {
		_, err := conn.SaveResult(result) // HL

		if err != nil {
			fmt.Printf("Error!")
		}
	}

	return
}

func getDatabaseConnection() (*Conn, error) {
	return &Conn{}, nil
}

type Conn struct {
}

func (c Conn) SaveResult(s string) (string, error) {
	return s, nil
}
