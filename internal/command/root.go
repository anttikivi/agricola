package command

import "github.com/anttikivi/agricola/internal/ui"

// A Root is the root command of the program. It is included into the other
// commands and it holds common information, for example the user interface.
type Root struct {
	// ui is the user interface used by the commands.
	ui ui.UserInterface
}

func CreateRootCommand(ui ui.UserInterface) *Root {
	return &Root{ui: ui}
}
