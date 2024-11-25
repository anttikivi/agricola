package command

import (
	"flag"
	"io"

	"github.com/anttikivi/agricola/internal/ui"
)

// A Meta is a struct that represents the meta-options that are available on all or most commands.
//
// The Command interface is based on OpenTofu.
// See: https://github.com/opentofu/opentofu/blob/main/internal/command/meta.go
type Meta struct {
	ui.Ui

	// IsColored tells whether or not the printed output should be colored.
	IsColored bool
}

// defaultFlagSet creates the base flag set for the subcommand given with the string n.
func (m *Meta) defaultFlagSet(n string) *flag.FlagSet {
	f := flag.NewFlagSet(n, flag.ContinueOnError)
	f.SetOutput(io.Discard)
	f.Usage = func() {}
	return f
}
