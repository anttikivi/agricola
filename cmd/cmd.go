package cmd

import (
	"log"
	"os"
	"runtime"

	"github.com/anttikivi/agricola/internal/command"
	"github.com/anttikivi/agricola/internal/crash"
	"github.com/anttikivi/agricola/internal/logging"
	"github.com/anttikivi/agricola/internal/ui"
	"github.com/anttikivi/agricola/version"
)

const commandNotFound = 4

// Execute runs the program and returns its exit code.
func Execute() int {
	defer crash.HandlePanic()

	logging.Init()

	ui := ui.CreateBasicUI(os.Stdout, os.Stderr) //nolint:varnamelen
	rootCmd := command.CreateRootCommand(ui, version.GetVersion())

	log.Printf("[TRACE] Raw version information: %s", version.GetRawVersion())
	log.Printf("[INFO] %s version: %v", command.Name, version.GetVersion())
	log.Printf("[INFO] Go runtime version: %s", runtime.Version())
	log.Printf("[INFO] CLI args: %#v", os.Args)

	// TODO: Initialize output.
	// TODO: Initialze configuration.

	args := os.Args[1:]

	// Allow using '-v', '-version', or '--version' as a shortcut for
	// version command.
	for _, arg := range args {
		if arg == "-v" || arg == "-version" || arg == "--version" {
			newArgs := make([]string, len(args)+1)
			newArgs[0] = "version"
			copy(newArgs[1:], args)
			args = newArgs

			break
		}
	}

	noColor := false
	i := 0

	// Allow using '-no-color' or '--no-color' to disable coloring the output.
	for _, arg := range args {
		if arg == "-no-color" || arg == "--no-color" {
			noColor = true
		} else {
			args[i] = arg
			i++
		}
	}

	// Remove the duplicated last argument.
	args = args[:i]

	if noColor {
		log.Print("[INFO] Output coloring disabled")
	}

	log.Printf("[INFO] Arguments for the command: %#v", args)

	subcommands := command.Commands(rootCmd)

	cmdStr := ""
	if len(args) > 0 {
		cmdStr = args[0]
	}

	log.Printf("[DEBUG] Using %q as the subcommand", cmdStr)

	cmd, ok := subcommands[cmdStr]
	if !ok {
		log.Printf("[INFO] The subcommand %q was not found", cmdStr)
		// TODO: Print help.
		ui.Output(command.Usage(subcommands))

		return commandNotFound
	}

	log.Printf("[DEBUG] The arguments passed to the command: %#v", args[1:])

	exitCode := cmd.Execute(args[1:])

	return exitCode
}
