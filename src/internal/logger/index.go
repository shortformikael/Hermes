// Package logger: Handles logging
package logger

import (
	"fmt"
	"time"
)

// TODO: Add colors

func getTime() string {
	return time.Now().Format("02-01-06 15:04")
}

func Info(message string) {
	logTime := getTime()
	fmt.Printf("[%v] [INFO] %v \n", logTime, message)
}

func Error(message string) {
	fmt.Printf("[%v] [ERROR] %v \n", getTime(), message)
}

func Warning(message string) {
	fmt.Printf("[%v] [WARNING] %v \n", getTime(), message)
}

func Debug(message string) {
	fmt.Printf("[%v] [DEBUG] %v \n", getTime(), message)
}

func Fatal(message string) {
	fmt.Printf("[%v] [FATAL] %v \n", getTime(), message)
}
