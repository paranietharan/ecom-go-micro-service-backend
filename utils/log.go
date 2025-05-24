package utils

import (
	"log"
	"os"
	"time"
)

var (
	fileLogger *log.Logger
)

func InitLogger() error {
	logFile, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return err
	}

	fileLogger = log.New(logFile, "APP_LOG: ", log.Ldate|log.Ltime|log.Lshortfile)
	return nil
}

func LogInfo(message string) {
	fileLogger.Printf("[INFO] %s :: %s", time.Now().Format(time.RFC3339), message)
}

func LogError(err error) {
	fileLogger.Printf("[ERROR] %s :: %v", time.Now().Format(time.RFC3339), err)
}
