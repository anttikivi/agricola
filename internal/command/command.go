package command

import "fmt"

const (
	Name        = "Agricola"
	CommandName = "ager"
)

// A Command represents a command line subcommand.
type Command interface {
	// Execute executes the command with the given arguments and returns
	// the exit code.
	Execute(args []string) int

	// Summary returns a short help string for the command used in the command
	// line help.
	Summary() string

	// Usage returns the usage of the command. It is printed when
	// the subcommand is called incorrectly or with the help flag.
	Usage() string
}

// Commands returns a map the contains the implementations of Command. The keys
// are strings that act as the subcommand.
func Commands(rootCmd RootCommand) map[string]Command {
	return map[string]Command{
		"version": &VersionCommand{RootCommand: rootCmd},
	}
}

// Usage returns the help message for the program when it is run without a
// required subcommand or when the help message for the program is requested.
// The subcommands are passed in as the parameter.
func Usage(_ map[string]Command) string {
	s := `
%s is a tool for managing web application deployments declaratively.

Usage:

	%s <command> [arguments]

The commands are:

%s
`

	return fmt.Sprintf(s, Name, CommandName)
}
