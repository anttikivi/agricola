package ui

import (
	"fmt"
	"io"
)

// A BasicUserInterface is a UserInterface that is used for printing without
// colors. It can be used for printing before the color options are parsed.
type BasicUserInterface struct {
	// out is the writer used to write normal messages.
	out io.Writer

	// err is the writer used to write error messages.
	err io.Writer
}

func (ui *BasicUserInterface) Error(msg string) {
	w := ui.out
	if ui.err != nil {
		w = ui.err
	}

	fmt.Fprint(w, msg)
	fmt.Fprint(w, "\n")
}

func (ui *BasicUserInterface) Output(msg string) {
	fmt.Fprint(ui.out, msg)
	fmt.Fprint(ui.out, "\n")
}

func BasicUI(out io.Writer, err io.Writer) *BasicUserInterface {
	return &BasicUserInterface{out: out, err: err}
}
