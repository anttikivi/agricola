package help

import (
	"github.com/anttikivi/agricola/internal/command"
	"github.com/anttikivi/agricola/internal/ui"
)

// Help implements the 'help' command.
func Help(_ ui.UserInterface, _ []string) int {
	return 0
}

// PrintUsage prints the usage for the given command to the user interface.
func PrintUsage(ui ui.UserInterface, _ *command.Command) {
	ui.Error("TODO: Implement.")
}
