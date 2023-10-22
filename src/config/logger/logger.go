package logger

import (
	"fmt"
	"log"
	"os"
	"runtime"
)

const logFile = "logs.log"

func logCaller() string {
	_, file, line, _ := runtime.Caller(3)
	return fmt.Sprintf("%s:%d", file, line)
}

func newLog(logInfo *TLogger) {
	logFile, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		log.Panic(err)
	}

	defer logFile.Close()

	log.SetOutput(logFile)

	log.SetPrefix(prefix[logInfo.Prefix])

	info := fmt.Sprintf("%s - %s", logInfo.Journey, logInfo.Message)

	if env := os.Getenv("ENV"); env == "production" {
		fmt.Println(info)
	}

	log.Println(info)
}

func NewInfoLog(journey, message string) {
	newLog(&TLogger{
		Prefix:  INFO,
		Journey: journey,
		Message: message,
	})
}

func NewErrorLog(journey, message string) {
	message = fmt.Sprintf("%s - %s", message, logCaller())

	newLog(&TLogger{
		Prefix:  ERROR,
		Journey: journey,
		Message: message,
	})
}

func NewWarningLog(journey, message string) {
	newLog(&TLogger{
		Prefix:  WARN,
		Journey: journey,
		Message: message,
	})
}
