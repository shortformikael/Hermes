// Package logger: Handles logging
package logger

import (
	"fmt"
	"time"
)

func getTime() string {
	return time.Now().Format("02-01-06 15:04")
}

func Info(message string) {
	logTime := getTime()
	fmt.Printf("[%v] [INFO] %v \n", logTime, message)
}
