package logging

import (
	"fmt"
	"io"
	"os"
	"time"
)

type logger struct {
	level        Level
	levelStrings map[Level]string
	name         string
	timeFormat   string
	timeFunc     func() time.Time
	writer       *bufferWriter
}

func createLogger(level Level, output io.Writer) *logger {
	if output == nil {
		// Let's use stderr as the default output.
		output = io.Writer(os.Stderr)
	}

	return &logger{
		level: level,
		levelStrings: map[Level]string{
			LevelTrace: "[TRACE]",
			LevelDebug: "[DEBUG]",
			LevelInfo:  "[INFO]",
			LevelWarn:  "[WARN]",
			LevelError: "[ERROR]",
		},
		name: "",
		// TODO: Allow customizing this.
		timeFormat: defaultTimeFormat,
		timeFunc:   time.Now,
		writer:     createWriter(output),
	}
}

func (l *logger) createLogWriter() io.Writer {
	return &standardWriter{
		logger: l,
	}
}

// log writes the given message to the log with the given level.
//
// TODO: See if the possible errors while writing should be handled differently.
func (l *logger) log(level Level, msg string) {
	if level < l.level {
		return
	}

	var err error

	t := l.timeFunc()

	// Write the time string at the start.
	// TODO: Maybe add option to disable the timestamps.
	if err = l.writer.writeString(t.Format(l.timeFormat)); err != nil {
		fmt.Printf("failed to write the time format to the log: %v\n", err) //nolint:forbidigo
	}

	if err = l.writer.writeByte(' '); err != nil {
		fmt.Printf("failed to write a space after the time to the log: %v\n", err) //nolint:forbidigo
	}

	if s, ok := l.levelStrings[level]; ok {
		if err = l.writer.writeString(s); err != nil {
			fmt.Printf("failed to write the log level to the log: %v\n", err) //nolint:forbidigo
		}
	} else {
		if err = l.writer.writeString("[UNKNOWN]"); err != nil {
			fmt.Printf("failed to write the unknown log level to the log: %v\n", err) //nolint:forbidigo
		}
	}

	if err = l.writer.writeByte(' '); err != nil {
		fmt.Printf("failed to write a space after the log level to the log: %v\n", err) //nolint:forbidigo
	}

	if l.name != "" {
		if err = l.writer.writeString(l.name); err != nil {
			fmt.Printf("failed to write the log name to the log: %v\n", err) //nolint:forbidigo
		}

		if err = l.writer.writeString(": "); err != nil {
			fmt.Printf("failed to write a colon and a space after the log name to the log: %v\n", err) //nolint:forbidigo
		}
	}

	if msg != "" {
		if err = l.writer.writeString(msg); err != nil {
			fmt.Printf("failed to write the message to the log: %v\n", err) //nolint:forbidigo
		}
	}

	if err = l.writer.writeString("\n"); err != nil {
		fmt.Printf("failed to write a newline to the end of a log message: %v\n", err) //nolint:forbidigo
	}

	l.writer.flush()
}

// trace writes a trace-level message to the log.
func (l *logger) trace(s string) {
	l.log(LevelTrace, s)
}

// debug writes a debug-level message to the log.
func (l *logger) debug(s string) {
	l.log(LevelDebug, s)
}

// info writes a info-level message to the log.
func (l *logger) info(s string) {
	l.log(LevelInfo, s)
}

// warn writes a warning-level message to the log.
func (l *logger) warn(s string) {
	l.log(LevelWarn, s)
}

// error writes a error-level message to the log.
func (l *logger) error(s string) {
	l.log(LevelError, s)
}
