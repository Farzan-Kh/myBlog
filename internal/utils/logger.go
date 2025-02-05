package utils

import (
	"log"
	"os"
	"path/filepath"
)

var (
	InfoLogger  *log.Logger
	ErrorLogger *log.Logger
)

func InitLogger(info *log.Logger, error *log.Logger) {
	logDir := "logs"
	logFile := filepath.Join(logDir, "logs.txt")

	// Create the directory if it doesn't exist
	if err := os.MkdirAll(logDir, 0755); err != nil {
		log.Fatal(err.Error())
	}

	file, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err.Error())
	}

	InfoLogger = log.New(file, "INFO: ", log.LUTC|log.Ltime|log.Ldate|log.Lshortfile)
	ErrorLogger = log.New(file, "ERROR: ", log.LUTC|log.Ltime|log.Ldate|log.Lshortfile)
}
