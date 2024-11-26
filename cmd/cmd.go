package cmd

import "github.com/anttikivi/agricola/internal/logging"

const (
	Name        = "Agricola"
	CommandName = "ager"
)

// Execute runs the program and returns its exit code.
func Execute() int {
	defer handlePanic()

	logging.Init()

	return 0
}
