package core

import (
	"log"
	"os"
)

var Logger *log.Logger

func InitLogger() {
	f, err := os.OpenFile("log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Println("✗ Failed to create log file:", err)
		return
	}
	Logger = log.New(f, "", log.LstdFlags)
}