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

// Info - generally useful information to log
// (service start/stop, configuration assumptions, etc).
// Info I want to always have available but usually
// don't care about under normal circumstances.
func Info(message string) {
	outputMessage(levelInfo, message)
}

// FInfo appends message to the end of file in path string with level "Info".
func FInfo(message string, path string) {
	writeInFile(path, "Info", message)
}

// Error - any error which is fatal to the operation,
// but not the service or application (can't open a
// required file, missing data, etc.).
// These errors will force user (administrator, or direct user)
// intervention. These are usually reserved (in my apps)
// for incorrect connection strings, missing services, etc.
func Error(message string) {
	outputMessage(levelError, message)
}

// FError appends message to the end of file in path string with level "Error".
func FError(message string, path string) {
	writeInFile(path, "Error", message)
}

//Warning - Anything that can potentially cause application
// oddities, but for which I am automatically recovering.
// (Such as switching from a primary to backup server,
// retrying an operation, missing secondary data, etc.).
func Warning(message string) {
	outputMessage(levelWarning, message)
}
// FWarning appends message to the end of file in path string with level "Warning".
func FWarning(message string, path string) {
	writeInFile(path, "Warning", message)
}

// Debug - Information that is diagnostically helpful
// to people more than just developers (IT, sysadmins, etc.).
func Debug(message string) {
	outputMessage(levelDebug, message)
}
// FDebug appends message to the end of file in path string with level "Debug".
func FDebug(message string, path string) {
	writeInFile(path, "Debug", message)
}
// Fatal - Any error that is forcing a shutdown of
// the service or application to prevent data loss
// (or further data loss). I reserve these only for
// the most heinous errors and situations where
// there is guaranteed to have been data corruption or loss.
func Fatal(message string) {
	outputMessage(levelFatal, message)
}

// FFatal appends message to the end of file in path string with level "Fatal".
func FFatal(message string, path string) {
	writeInFile(path, "Fatal", message)
}
