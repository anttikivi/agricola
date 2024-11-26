package cmd

import (
	"log"
	"os"
	"runtime"

	"github.com/anttikivi/agricola/internal/command"
	"github.com/anttikivi/agricola/internal/crash"
	"github.com/anttikivi/agricola/internal/logging"
	"github.com/anttikivi/agricola/version"
)

// Execute runs the program and returns its exit code.
func Execute() int {
	defer crash.HandlePanic()

	logging.Init()

	log.Printf("[INFO] %s version: %s", command.Name, version.Version())
	log.Printf("[INFO] Go runtime version: %s", runtime.Version())
	log.Printf("[INFO] CLI args: %#v", os.Args)

	return 0
}
