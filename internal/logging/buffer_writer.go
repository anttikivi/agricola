package logging

import (
	"bytes"
	"fmt"
	"io"
)

// bufferWriter is the writer used by the logger to write the messages.
// After writing a message, the logger flushes the bufferWriter.
// flush writes the buffer in the bufferWriter to its writer that outputs the message.
type bufferWriter struct {
	buffer bytes.Buffer
	writer io.Writer
}

func (w *bufferWriter) Write(p []byte) (int, error) {
	n, err := w.buffer.Write(p)
	if err != nil {
		return n, fmt.Errorf("failed to write to the buffer writer's byte buffer: %w", err)
	}

	return n, nil
}

func (w *bufferWriter) flush() error {
	toWrite := w.buffer.Bytes()
	if _, err := w.writer.Write(toWrite); err != nil {
		return fmt.Errorf("failed to flush the buffer writer: %w", err)
	}

	w.buffer.Reset()

	return nil
}

func (w *bufferWriter) writeByte(c byte) error {
	if err := w.buffer.WriteByte(c); err != nil {
		return fmt.Errorf("failed to write a byte to the buffer writer's byte buffer: %w", err)
	}

	return nil
}

func (w *bufferWriter) writeString(s string) error {
	// TODO: There was another return value, n, but it was removed as it was not used.
	_, err := w.buffer.WriteString(s)
	if err != nil {
		return fmt.Errorf("failed to write a string to the buffer writer's byte buffer: %w", err)
	}

	return nil
}

func createWriter(w io.Writer) *bufferWriter {
	return &bufferWriter{
		buffer: bytes.Buffer{},
		writer: w,
	}
}
