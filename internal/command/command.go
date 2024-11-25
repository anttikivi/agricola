package command

// A Command is a runnable sub-command of the CLI.
//
// The Command interface is based on `mitchellh/cli` by Mitchell Hashimoto.
// See: https://github.com/mitchellh/cli/blob/main/command.go
type Command interface {
	// Help returns the long-form help message for the command.
	Help() string

	// Run should run the actual command.
	// It should return the exit code of the command.
	Run(args []string) int

	// Summary returns the short, one-line help synopsis for the command.
	Summary() string
}
