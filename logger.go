package logger

import (
	"fmt"
	"os"
	"time"
)

var now = time.Now().Format("_2 Jan 15:04:05.000")

func outputMessage(level string, message string) {
	fmt.Printf("%v\t%s: %s\n", now, level, message)
}

func writeInFile(path string, level, message string) {
	file, err := os.OpenFile(path, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("file create/open error")
		os.Exit(1)
	}

	defer file.Close()
	fullMessage := fmt.Sprintf("%v\t%s: %s\n", now, level, message)
	file.Write([]byte(fullMessage))
}

func Info(message string) {
	outputMessage("Info", message)
}

func FInfo(message string, path string) {
	writeInFile(path, "Info", message)
}

func Error(message string) {
	outputMessage("Error", message)
}

func FError(message string, path string) {
	writeInFile(path, "Error", message)
}

func Warning(message string) {
	outputMessage("Warning", message)
}

func FWarning(message string, path string) {
	writeInFile(path, "Warning", message)
}

func Debug(message string) {
	outputMessage("Debug", message)
}

func FDebug(message string, path string) {
	writeInFile(path, "Debug", message)
}

func Fatal(message string) {
	outputMessage("Fatal", message)
}

func FFatal(message string, path string) {
	writeInFile(path, "Fatal", message)
}
