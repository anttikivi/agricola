package main

import (
	_ "embed"
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	"slices"
	"strings"

	"github.com/anttikivi/agricola/internal/command"
	"github.com/anttikivi/agricola/internal/command/help"
	"github.com/anttikivi/agricola/internal/command/version"
	"github.com/anttikivi/agricola/internal/crash"
	"github.com/anttikivi/agricola/internal/logging"
	"github.com/anttikivi/agricola/internal/semver"
	"github.com/anttikivi/agricola/internal/ui"
)

const helpCmdName = "help"

// rawVersion is the raw version value read from the VERSION file. It is used
// if buildVersion is not set.
//
//go:embed VERSION
var rawVersion string

// buildVersion is the version set using linker flags build time. It is used to
// over the value embedded from the VERSION file if set.
var buildVersion string //nolint:gochecknoglobals

func main() {
	os.Exit(run())
}

// run initializes the program and runs the specified command.
// The return value is the exit code of the program.
func run() int {
	defer crash.HandlePanic()

	// TODO: Maybe handle the chdir flag "-C" here.
	// The working directory should be set before logging is initialized as the
	// log can be written relative to it.

	logging.Init()
	ui.Init()

	ver := parseVersion()

	log.Printf("[TRACE] Raw version information: %s", rawVersionString())
	log.Printf("[INFO] %s version: %v", command.Name, ver)
	log.Printf("[INFO] Go runtime version: %s", runtime.Version())
	log.Printf("[INFO] CLI args: %#v", os.Args)

	ager := command.BaseCommand()
	ager.Commands = []*command.Command{
		version.Command(ver),
	}

	flag.Usage = func() { help.PrintUsage(ager) }
	flag.Parse()

	args := flag.Args()

	log.Printf("[DEBUG] Arguments after parsing the global flags: %#v", args)

	if len(args) < 1 {
		help.PrintUsage(ager)

		return command.ExitInvalidArgs
	}

	// TODO: Should I also allow using "-h", "-help", and "--help" flags?
	if args[0] == helpCmdName {
		return help.Help(args[1:])
	}

	cmd, used := lookupCmd(ager, args)
	if len(cmd.Commands) > 0 {
		if used >= len(args) {
			help.PrintUsage(cmd)

			return command.ExitInvalidArgs
		}

		if args[used] == helpCmdName {
			// Accept "ager plow help" and "ager plow help foo" for "ager help plow" and "ager help plow foo".
			help.Help(append(slices.Clip(args[:used]), args[used+1:]...))

			return command.ExitSuccess
		}

		helpArg := ""
		if used > 0 {
			helpArg = " " + strings.Join(args[:used], " ")
		}

		cmdName := strings.Join(args[:used], " ")
		if cmdName == "" {
			cmdName = args[0]
		}

		ui.Errorf("%s %s: unknown command\n", command.CommandName, cmdName)
		ui.Errorf("Run '%s help%s' for usage\n", command.CommandName, helpArg)

		return command.ExitInvalidArgs
	}

	exitCode := invoke(cmd, args[used-1:])

	return exitCode
}

func invoke(cmd *command.Command, args []string) int {
	if err := cmd.Flag.Parse(args[1:]); err != nil {
		ui.Errorf("Error parsing command-line flags: %v\n", err)

		return command.ExitInvalidArgs
	}

	args = cmd.Flag.Args()
	log.Printf("[INFO] Running %s command with arguments: %#v", cmd.Name(), args)

	return cmd.Run(cmd, args)
}

// lookupCmd finds the initial command to run from the base command and the
// given args.
// It tries to find the first runnable command or a subcommand group.
// It returns the found command and the number of arguments used in the lookup.
func lookupCmd(baseCmd *command.Command, args []string) (*command.Command, int) {
	cmd, used := baseCmd, 0
	for used < len(args) {
		c := cmd.Lookup(args[used])
		if c == nil {
			break
		}

		if c.Runnable() {
			cmd = c
			used++

			break
		}

		if len(c.Commands) > 0 {
			cmd = c
			used++

			if used >= len(args) || args[0] == helpCmdName {
				break
			}

			continue
		}

		break
	}

	return cmd, used
}

// parseVersion parses the program version from the version data set during
// build.
// It panics if the version cannot be parsed as the version string set during
// builds must not be an invalid version.
func parseVersion() semver.Version {
	v, err := semver.Parse(rawVersionString())
	if err != nil {
		panic(fmt.Sprintf("failed to parse the version: %v", err))
	}

	return v
}

// rawVersionString returns the unparsed version string the will be used to
// parse the program version.
func rawVersionString() string {
	s := buildVersion
	if s == "" {
		s = rawVersion
	}

	s = strings.TrimSpace(s)

	return s
}
