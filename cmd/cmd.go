package cmd

import (
	"log"
	"os"
	"runtime"

	"github.com/anttikivi/agricola/internal/logging"
	"github.com/anttikivi/agricola/version"
)

const (
	Name        = "Agricola"
	CommandName = "ager"
)

// Execute runs the program and returns its exit code.
func Execute() int {
	defer handlePanic()

	logging.Init()

	log.Printf("[INFO] %s version: %s", Name, version.Version())
	log.Printf("[INFO] Go runtime version: %s", runtime.Version())
	log.Printf("[INFO] CLI args: %#v", os.Args)

	return 0
}
