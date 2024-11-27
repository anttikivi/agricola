package cmd

import (
	"fmt"
	"log"
	"os"
	"runtime"

	"github.com/anttikivi/agricola/internal/command"
	"github.com/anttikivi/agricola/internal/crash"
	"github.com/anttikivi/agricola/internal/logging"
	"github.com/anttikivi/agricola/version"
)

const commandNotFound = 4

// Execute runs the program and returns its exit code.
func Execute() int {
	defer crash.HandlePanic()

	logging.Init()

	log.Printf("[INFO] %s version: %s", command.Name, version.Version())
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

	log.Printf("[INFO] Arguments for the command: %#v", args)

	commands := command.GetCommands()

	log.Printf("[DEBUG] Using %q as the subcommand", args[0])

	cmd, ok := commands[args[0]]
	if !ok {
		log.Printf("[WARN] The subcommand %q was not found", args[0])
		// TODO: Print help.
		fmt.Printf("%s has no subcommand named %q.\n", command.Name, args[0]) //nolint:forbidigo

		return commandNotFound
	}

	log.Printf("[DEBUG] The arguments passed to the command: %#v", args[1:])

	exitCode := cmd.Execute(args[1:])

	return exitCode
}
