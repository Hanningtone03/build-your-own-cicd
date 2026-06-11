package logger

import (
	"fmt"
	"time"
)

const (
	colorReset  = "\033[0m"
	colorGreen  = "\033[32m"
	colorRed    = "\033[31m"
	colorYellow = "\033[33m"
	colorCyan   = "\033[36m"
)

func timestamp() string {
	return time.Now().Format("15:04:05")
}

func Info(msg string) {
	fmt.Printf("%s[%s] [INFO]  %s%s\n", colorCyan, timestamp(), msg, colorReset)
}

func Success(msg string) {
	fmt.Printf("%s[%s] [OK]    %s%s\n", colorGreen, timestamp(), msg, colorReset)
}

func Error(msg string) {
	fmt.Printf("%s[%s] [FAIL]  %s%s\n", colorRed, timestamp(), msg, colorReset)
}

func Warn(msg string) {
	fmt.Printf("%s[%s] [WARN]  %s%s\n", colorYellow, timestamp(), msg, colorReset)
}

func Step(name string) {
	fmt.Printf("\n%s[%s] [STEP] >>> %s%s\n", colorCyan, timestamp(), name, colorReset)
}