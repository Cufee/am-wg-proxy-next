package logs

import (
	"fmt"
	"log"
)

func Wrap(err error, message string) error {
	return fmt.Errorf("%s: %v", message, err)
}

func Debug(message string, args ...interface{}) {
	fmt.Printf("[DEBUG] %s\n", fmt.Sprintf(message, args...))
}

func Info(message string, args ...interface{}) {
	fmt.Printf("[INFO] %s\n", fmt.Sprintf(message, args...))
}

func Warning(message string, args ...interface{}) {
	fmt.Printf("[WARNING] %s\n", fmt.Sprintf(message, args...))
}

func Error(message string, args ...interface{}) {
	fmt.Printf("[ERROR] %s\n", fmt.Sprintf(message, args...))
}

func Critical(message string, args ...interface{}) {
	fmt.Printf("[CRITICAL] %s\n", fmt.Sprintf(message, args...))
}

func Fatal(message string, args ...interface{}) {
	log.Fatalf("[FATAL] %s", fmt.Sprintf(message, args...))
}
