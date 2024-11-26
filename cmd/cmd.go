package cmd

import (
	"log"

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

	return 0
}
