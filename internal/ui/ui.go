// Package ui implements utilities for interacting with the user. It should be
// used for interacting with the user in contrast to the logger that is used to
// record messages related to the execution of the program. For example, the
// output printed by the utilities in the ui package should be printed visibly
// to the command line by default.
package ui

import (
	"fmt"
	"io"
	"os"
)

// DefaultWriter is the writer that the ui package writes the messages to when
// no specific writer is specified.
// It is used in particular by the default writing function that takes no writer
// as its input.
// The DefaultWriter is specified as a global variable because alternative
// implementations unnecessarily complicate the code.
var DefaultWriter io.Writer //nolint:gochecknoglobals

// DefaultErrorWriter is the writer that the ui package writes error messages to
// when no specific writer is specified.
// It is used in particular by the default error writing function that takes no
// writer as its input.
// The DefaultErrorWriter is specified as a global variable because alternative
// implementations unnecessarily complicate the code.
var DefaultErrorWriter io.Writer //nolint:gochecknoglobals

// Init initializes the global variables in the ui package.
// It should be called at the start of the program.
func Init() {
	DefaultWriter = os.Stdout
	DefaultErrorWriter = os.Stderr
}

// Write prints an error message to the default error output.
func Error(msg string) {
	fmt.Fprint(DefaultErrorWriter, msg)
}

// Write prints a message to the default output.
func Write(msg string) {
	fmt.Fprint(DefaultWriter, msg)
}

// Errorf formats according to a format specifier and prints the formatted
// message to the default error output.
func Errorf(format string, v ...any) {
	fmt.Fprintf(DefaultErrorWriter, format, v...)
}

// Printf formats according to a format specifier and prints the formatted
// message to the default output.
func Printf(format string, v ...any) {
	fmt.Fprintf(DefaultErrorWriter, format, v...)
}
