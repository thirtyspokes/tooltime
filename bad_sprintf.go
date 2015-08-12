package main

import (
	"fmt"
	"time"
)

func buildLogEntry(level int, message string) string {
	return fmt.Sprintf("[%s - %d] %d", time.Now().UTC(), level, message)
}
