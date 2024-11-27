package command

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

// GetCommands returns a map the contains the implementations of Command.
// The keys are strings that act as the subcommand.
func GetCommands(rootCmd *Root) map[string]Command {
	return map[string]Command{
		"version": &Version{root: rootCmd},
	}
}
