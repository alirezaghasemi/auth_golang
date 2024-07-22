package utils

import (
	"log"
	"os"
)

var Logger *log.Logger

func InitLogger() {
	file, err := os.OpenFile("project.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Failed to open log file", err)
	}

	Logger = log.New(file, "", log.Ldate|log.Ltime|log.Lshortfile)
}
