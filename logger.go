package logger

import (
	"fmt"
	"os"
	"time"
)

var now = time.Now().Format("_2 Jan 15:04:05.000")

const (
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
	levelInfo    = cyan + "Info" + reset
	levelError   = red + "Error" + reset
	levelWarning = yellow + "Warning" + reset
	levelDebug   = purple + "Debug" + reset
	levelFatal   = gray + "Fatal" + reset
)

// Logger - to use logger  you should initialise it and fill the FilePath,
// if you don't want to write in file your logs, just leave the field blank "".
type Logger struct {
	FilePath string
}

func outputMessage(level string, message string) {
	fmt.Printf("%v\t%s: %s\n", now, level, message)
}

func writeInFile(path string, level, message string) {
	file, err := os.OpenFile(path, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		outputMessage(levelError, "create/open file error")
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
func (l *Logger) Info(message string) {
	outputMessage(levelInfo, message)
}

// FInfo appends message to the end of file in path string with level "Info".
func (l *Logger) FInfo(message string) {
	writeInFile(l.FilePath, "Info", message)
}

// Error - any error which is fatal to the operation,
// but not the service or application (can't open a
// required file, missing data, etc.).
// These errors will force user (administrator, or direct user)
// intervention. These are usually reserved (in my apps)
// for incorrect connection strings, missing services, etc.
func (l *Logger) Error(message string) {
	outputMessage(levelError, message)
}

// FError appends message to the end of file in path string with level "Error".
func (l *Logger) FError(message string) {
	writeInFile(l.FilePath, "Error", message)
}

//Warning - Anything that can potentially cause application
// oddities, but for which I am automatically recovering.
// (Such as switching from a primary to backup server,
// retrying an operation, missing secondary data, etc.).
func (l *Logger) Warning(message string) {
	outputMessage(levelWarning, message)
}

// FWarning appends message to the end of file in path string with level "Warning".
func (l *Logger) FWarning(message string) {
	writeInFile(l.FilePath, "Warning", message)
}

// Debug - Information that is diagnostically helpful
// to people more than just developers (IT, sysadmins, etc.).
func (l *Logger) Debug(message string) {
	outputMessage(levelDebug, message)
}

// FDebug appends message to the end of file in path string with level "Debug".
func (l *Logger) FDebug(message string) {
	writeInFile(l.FilePath, "Debug", message)
}

// Fatal - Any error that is forcing a shutdown of
// the service or application to prevent data loss
// (or further data loss). I reserve these only for
// the most heinous errors and situations where
// there is guaranteed to have been data corruption or loss.
func (l *Logger) Fatal(message string) {
	outputMessage(levelFatal, message)
}

// FFatal appends message to the end of file in path string with level "Fatal".
func (l *Logger) FFatal(message string) {
	writeInFile(l.FilePath, "Fatal", message)
}
