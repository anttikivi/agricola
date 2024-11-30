package command

import (
	"flag"
	"io"
	"strings"

	"github.com/anttikivi/agricola/internal/ui"
)

const (
	Name        = "Agricola"
	CommandName = "ager"
)

const (
	ExitSuccess         = 0
	ExitInvalidArgs     = 2
	ExitCommandNotFound = 4
)

// A Command is an implementation of an Agricola command.
type Command struct {
	// Run runs the command.
	// The args are the arguments passed in after the command name.
	// The function returns the exit code of the command.
	Run func(cmd *Command, args []string) int

	// UsageLine is the one-line usage message.
	// The words between "ager" and the first flag or argument in the line are
	// taken to be the command name.
	UsageLine string

	// Short is the short description shown in the "ager help" output.
	Short string

	// Long is the long message shown in the "ager help <this-command>" output.
	Long string

	// Flag is a set of flags specific to this command.
	Flag *flag.FlagSet

	// Commands is a list of the available commands (so-called subcommands) and
	// other help topics for this command.
	// The order here is the order in which they are printed when running the
	// 'help' command.
	Commands []*Command
}

// LongName returns the command's long name.
// Command's long name is the string between "ager" and the first arguments.
func (c *Command) LongName() string {
	name := c.UsageLine
	if i := strings.Index(name, " ["); i >= 0 {
		name = name[:i]
	}

	if name == CommandName {
		return ""
	}

	return strings.TrimPrefix(name, CommandName+" ")
}

// Lookup returns the subcommand with the given name, if any.
// Otherwise it returns nil.
// Lookup ignores subcommands that have len(c.Commands) == 0 and c.Run == nil.
// Such subcommands are only meant to be used as arguments to "help".
func (c *Command) Lookup(n string) *Command {
	for _, cmd := range c.Commands {
		if cmd.Name() == n && (len(c.Commands) > 0 || c.Runnable()) {
			return cmd
		}
	}

	return nil
}

// Name return the command's short name.
// The short name of a command is the last word in the usage line before a flag
// or an argument.
func (c *Command) Name() string {
	n := c.LongName()
	if i := strings.LastIndex(n, " "); i >= 0 {
		n = n[i+1:]
	}

	return n
}

// Runnable reports whether the command can be run.
// If not, it is a pseudo-command used for printing help.
func (c *Command) Runnable() bool {
	return c.Run != nil
}

func (c *Command) Usage() {
	if c.UsageLine == CommandName {
		panic("(*Command).Usage() should not be called for the base command")
	}

	ui.Error("usage: " + c.UsageLine + "\n")
	ui.Error("Run '" + CommandName + " help " + c.LongName() + "' for details.\n")
}

func BaseCommand() *Command {
	return &Command{
		// TODO: I should implement something.
		Run:       nil,
		UsageLine: CommandName,
		Short:     "",
		Long:      Name + " is a tool for managing web application deployments declaratively.",
		Flag:      nil, // initialized in the main package
		Commands:  nil, // initialized in the main package
	}
}

func DefaultFlagSet(n string) *flag.FlagSet {
	f := flag.NewFlagSet(n, flag.ContinueOnError)
	f.SetOutput(io.Discard)
	f.Usage = func() {} // Usage must be set separately.

	return f
}
