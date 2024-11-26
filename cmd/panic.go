package cmd

import (
	"fmt"
	"os"
	"runtime/debug"
	"sync"
)

// segfault is the exit code for when the panic handler returns.
// It is fine as an exit code for when the program panics.
const segfault = 11

// panicLock is used to ensure that the only the first call to handlePanic prints if multiple goroutines panic.
var panicLock sync.Mutex //nolint:gochecknoglobals

// handlePanic is called to recover from an internal panic.
func handlePanic() {
	panicLock.Lock()
	defer panicLock.Unlock()

	if r := recover(); r != nil {
		fmt.Fprint(os.Stderr, Name, " crashed\n")
		fmt.Fprint(os.Stderr, r, "\n")
		debug.PrintStack()
		os.Exit(segfault) //nolint:gocritic
	}
}
