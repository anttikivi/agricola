package command

import (
	"github.com/anttikivi/agricola/internal/ui"
	"github.com/anttikivi/agricola/version"
)

// A RootCommand is the root command of the program. It is included into the other
// commands and it holds common information, for example the user interface.
type RootCommand struct {
	// ui is the user interface used by the commands.
	ui ui.UserInterface

	// version is the first of the program.
	version version.Version
}

func CreateRootCommand(ui ui.UserInterface, version version.Version) RootCommand {
	return RootCommand{ui: ui, version: version}
}
