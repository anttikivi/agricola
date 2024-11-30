package help

import (
	"text/template"

	"github.com/anttikivi/agricola/internal/command"
	"github.com/anttikivi/agricola/internal/ui"
)

// Help implements the 'help' command.
func Help(_ []string) int {
	return 0
}

// PrintUsage prints the usage for the given command to the user interface.
func PrintUsage(_ *command.Command) {
	ui.Error("TODO: Implement.\n")
}

// outputTemplate writes the given template text with the data from command cmd
// to the command's output.
func outputTemplate(text string, _ *command.Command) {
	t := template.New("top")
	template.Must(t.Parse(text))
}
