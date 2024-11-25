package ui

import (
	"fmt"
	"io"
)

// BasicUi is an implementation of Ui that just outputs to the given writer.
// This UI is not threadsafe by default.
type BasicUi struct {
	Reader      io.Reader
	Writer      io.Writer
	ErrorWriter io.Writer
}

func (u *BasicUi) Error(msg string) {
	w := u.Writer
	if u.ErrorWriter != nil {
		w = u.ErrorWriter
	}

	fmt.Fprint(w, msg)
	fmt.Fprint(w, "\n")
}

func (u *BasicUi) Output(msg string) {
	fmt.Fprint(u.Writer, msg)
	fmt.Fprint(u.Writer, "\n")
}
