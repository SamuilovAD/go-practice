package main

import (
	"log"
	"os"
)

func main() {
	CreateLogger("dev").Log("Test log!")
}

type Logger[T any] interface {
	Log(message T)
}

type ConsoleLogger[T any] struct {
}

func NewConsoleLogger() *ConsoleLogger[string] {
	return &ConsoleLogger[string]{}
}
func (consoleLogger *ConsoleLogger[string]) Log(message string) {
	println(message)
}

type FileLogger struct {
	logger *log.Logger
}

func NewFileLogger() *FileLogger {
	file, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic("cannot open log file: " + err.Error())
	}

	return &FileLogger{
		logger: log.New(file, "[PROD] ", log.Ldate|log.Ltime|log.Lshortfile),
	}
}

func (l *FileLogger) Log(message string) {
	l.logger.Println(message)
}

func CreateLogger(env string) Logger[string] {
	if env == "prod" {
		return NewFileLogger()
	}
	return NewConsoleLogger()
}
