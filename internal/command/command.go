package command

const (
	Name        = "Agricola"
	CommandName = "ager"
)

// A Command represents a command line subcommand.
type Command interface {
	// Help returns a long help string for the command.
	Help() string

	// Execute executes the command with the given arguments and returns the exit code.
	Execute(args []string) int

	// Summary return a short help string for the command used in the command line help.
	Summary() string
}
