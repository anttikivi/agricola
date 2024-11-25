package logging

import (
	"fmt"
	"io"
	"log"
)

// Debug prints a debug-level message to the log.
func Debug(format string, v ...any) {
	log.Printf(fmt.Sprintf("[DEBUG] %s", format), v...)
}

// Info prints an info-level message to the log.
func Info(format string, v ...any) {
	log.Printf(fmt.Sprintf("[INFO] %s", format), v...)
}

// Init initializes the logger for the given log output.
func Init(w io.Writer) {
	log.SetFlags(0)
	log.SetPrefix("")
	log.SetOutput(w)
}
