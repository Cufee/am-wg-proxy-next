package logs

import (
	"fmt"
	"log"
	"os"
	"strings"
)

var allEnvs []string

func init() {
	var env = os.Environ()
	for _, e := range env {
		split := strings.Split(e, "=")
		if len(split) != 2 {
			continue
		}
		allEnvs = append(allEnvs, split[1])
	}
}

func maskEnv(s string) string {
	for _, e := range allEnvs {
		s = strings.ReplaceAll(s, e, "***")
	}
	return s
}

func Debug(format string, args ...interface{}) {
	if os.Getenv("DEBUG") == "true" {
		log.Println(maskEnv(fmt.Sprintf("DEBUG: "+format, args...)))
	}
}

func Info(format string, args ...interface{}) {
	log.Println(maskEnv(fmt.Sprintf("INFO: "+format, args...)))
}

func Warning(format string, args ...interface{}) {
	log.Println(maskEnv(fmt.Sprintf("WARNING: "+format, args...)))
}

func Error(format string, args ...interface{}) {
	log.Println(maskEnv(fmt.Sprintf("ERROR: "+format, args...)))
}

func Critical(format string, args ...interface{}) {
	log.Println(maskEnv(fmt.Sprintf("CRITICAL: "+format, args...)))
}

func Fatal(format string, args ...interface{}) {
	log.Fatal(maskEnv(fmt.Sprintf("FATAL: "+format, args...)))
}

func Wrap(err error, format string, args ...interface{}) error {
	return fmt.Errorf(format+": %v", append(args, err)...)
}
