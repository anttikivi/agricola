package version

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/anttikivi/agricola/internal/command"
	"github.com/anttikivi/agricola/internal/semver"
	"github.com/anttikivi/agricola/internal/ui"
)

func Command(ui ui.UserInterface, ver semver.Version) *command.Command {
	c := &command.Command{
		Run:       func(cmd *command.Command, args []string) int { return runVersion(cmd, args, ver) },
		UsageLine: command.CommandName + " version",
		Short:     "prints " + command.Name + " version",
		Long:      fmt.Sprintf(`Version prints the version information of the %s binary.`, command.CommandName),
		Flag:      command.DefaultFlagSet("version"),
		Commands:  nil,
		UI:        ui,
	}
	c.Flag.Usage = func() { c.Usage() }

	return c
}

func runVersion(cmd *command.Command, _ []string, ver semver.Version) int {
	cmd.UI.Output(strings.ToLower(command.Name) + " version " + ver.String() + " " + runtime.GOOS + "/" + runtime.GOARCH)

	return command.ExitSuccess
}
