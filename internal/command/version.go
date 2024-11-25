package command

import (
	"bytes"
	"fmt"
	"runtime"
	"strings"
)

type VersionCommand struct {
	Meta

	Version    string
	Prerelease string
}

func (c *VersionCommand) Help() string {
	helpText := `
Usage: %s version

  Displays the version of %s.
`
	return strings.TrimSpace(fmt.Sprintf(helpText, CommandName, Name))
}

func (c *VersionCommand) Run(args []string) int {
	flags := c.defaultFlagSet("version")
	flags.Bool("v", true, "version")
	flags.Bool("version", true, "version")
	flags.Usage = func() { c.Ui.Error(c.Help()) }
	if err := flags.Parse(args); err != nil {
		c.Ui.Error(fmt.Sprintf("Error parsing command-line flags: %s\n", err.Error()))
		return 1
	}

	var versionString bytes.Buffer
	fmt.Fprintf(&versionString, "%s v%s", Name, c.Version)
	if c.Prerelease != "" {
		fmt.Fprintf(&versionString, "-%s", c.Prerelease)
	}
	fmt.Fprintf(&versionString, " %s/%s", runtime.GOOS, runtime.GOARCH)

	c.Ui.Output(versionString.String())

	return 0
}

func (c *VersionCommand) Summary() string {
	return fmt.Sprintf("Show the current %s version", Name)
}
