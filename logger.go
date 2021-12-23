package logger

import (
	"fmt"
	"os"
	"time"
)

var (
	now = time.Now().Format("_2 Jan 15:04:05.000")

	// Colors
	reset  = "\033[0m"
	red    = "\033[31m"
	green  = "\033[32m"
	yellow = "\033[33m"
	blue   = "\033[34m"
	purple = "\033[35m"
	cyan   = "\033[36m"
	gray   = "\033[37m"
	white  = "\033[97m"

	// Levels
	levelInfo = cyan + "Info" + reset
	levelError = red + "Error" + reset
	levelWarning = yellow + "Warning" + reset
	levelDebug = purple + "Debug" + reset
	levelFatal = gray + "Fatal" + reset

)

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
	outputMessage(levelInfo, message)
}

func FInfo(message string, path string) {
	writeInFile(path, "Info", message)
}

func Error(message string) {
	outputMessage(levelError, message)
}

func FError(message string, path string) {
	writeInFile(path, "Error", message)
}

func Warning(message string) {
	outputMessage(levelWarning, message)
}

func FWarning(message string, path string) {
	writeInFile(path, "Warning", message)
}

func Debug(message string) {
	outputMessage(levelDebug, message)
}

func FDebug(message string, path string) {
	writeInFile(path, "Debug", message)
}

func Fatal(message string) {
	outputMessage(levelFatal, message)
}

func FFatal(message string, path string) {
	writeInFile(path, "Fatal", message)
}
