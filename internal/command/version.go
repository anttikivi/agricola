package command

import (
	"flag"
	"fmt"
	"io"
	"runtime"
	"strings"
)

// A VersionCommand is the version subcommand.
type VersionCommand struct {
	RootCommand
}

func (v *VersionCommand) Execute(args []string) int {
	f := flag.NewFlagSet("version", flag.ContinueOnError)
	f.SetOutput(io.Discard)
	f.Bool("v", true, "version")
	f.Bool("version", true, "version")
	f.Usage = func() { v.RootCommand.ui.Error(v.Usage()) }

	if err := f.Parse(args); err != nil {
		v.RootCommand.ui.Error(fmt.Sprintf("Error parsing command-line flags: %s\n", err.Error()))

		return 1
	}

	v.RootCommand.ui.Output(Name + " " + v.RootCommand.version.String() + " " + runtime.GOOS + "/" + runtime.GOARCH)

	return 0
}

func (v *VersionCommand) Summary() string {
	return "prints " + Name + " version"
}

func (v *VersionCommand) Usage() string {
	usageText := `
Usage: %s version

  Displays the version of %s.
`

	return strings.TrimSpace(fmt.Sprintf(usageText, CommandName, Name))
}
