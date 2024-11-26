package command

import (
	"flag"
	"fmt"
	"io"
	"runtime"
	"strings"

	"github.com/anttikivi/agricola/version"
)

// A Version is the version subcommand.
type Version struct {
	root *Root
}

func (v *Version) Execute(args []string) int {
	f := flag.NewFlagSet("version", flag.ContinueOnError)
	f.SetOutput(io.Discard)
	f.Bool("v", true, "version")
	f.Bool("version", true, "version")
	f.Usage = func() { v.root.ui.Error(v.Usage()) }

	if err := f.Parse(args); err != nil {
		v.root.ui.Error(fmt.Sprintf("Error parsing command-line flags: %s\n", err.Error()))

		return 1
	}

	v.root.ui.Output(Name + " " + version.Version().String() + " " + runtime.GOOS + "/" + runtime.GOARCH)

	return 0
}

func (v *Version) Summary() string {
	return ""
}

func (v *Version) Usage() string {
	usageText := `
Usage: %s version

  Displays the version of %s.
`

	return strings.TrimSpace(fmt.Sprintf(usageText, CommandName, Name))
}
