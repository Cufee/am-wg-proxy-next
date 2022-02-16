package logs

import (
	"fmt"
	"log"
	"os"
)

func Debug(format string, args ...interface{}) {
	if os.Getenv("DEBUG") == "true" {
		log.Printf("DEBUG: "+format, args...)
	}
}

func Info(format string, args ...interface{}) {
	log.Printf("INFO: "+format, args...)
}

func Warning(format string, args ...interface{}) {
	log.Printf("WARNING: "+format, args...)
}

func Error(format string, args ...interface{}) {
	log.Printf("ERROR: "+format, args...)
}

func Critical(format string, args ...interface{}) {
	log.Printf("CRITICAL: "+format, args...)
}

func Fatal(format string, args ...interface{}) {
	log.Fatalf("FATAL: "+format, args...)
}

func Wrap(err error, format string, args ...interface{}) error {
	return fmt.Errorf(format+": %v", append(args, err)...)
}
