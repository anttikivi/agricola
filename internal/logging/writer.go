package logging

import (
	"bytes"
	"io"
)

type writer struct {
	b bytes.Buffer
	w io.Writer
}

func (w *writer) Write(p []byte) (int, error) {
	return w.b.Write(p)
}

func (w *writer) flush() error {
	toWrite := w.b.Bytes()
	_, err := w.w.Write(toWrite)
	w.b.Reset()

	return err
}

func (w *writer) writeByte(c byte) error {
	return w.b.WriteByte(c)
}

func (w *writer) writeString(s string) (int, error) {
	return w.b.WriteString(s)
}

func createWriter(w io.Writer) *writer {
	return &writer{w: w}
}
