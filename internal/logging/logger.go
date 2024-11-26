package logging

import (
	"io"
	"os"
	"time"
)

type logger struct {
	level         Level
	levelBrackets map[Level]string
	name          string
	timeFormat    string
	timeFunc      func() time.Time
	writer        *writer
}

func createLogger(level Level, output io.Writer) *logger {
	if output == nil {
		// Let's use stderr as the default output.
		output = io.Writer(os.Stderr)
	}

	return &logger{
		level: level,
		levelBrackets: map[Level]string{
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

func (l *logger) log(level Level, msg string) {
	if level < l.level {
		return
	}

	t := l.timeFunc()

	// Write the time string at the start.
	// TODO: Maybe add option to disable the timestamps.
	l.writer.writeString(t.Format(l.timeFormat))
	l.writer.writeByte(' ')

	if s, ok := l.levelBrackets[level]; ok {
		l.writer.writeString(s)
	} else {
		l.writer.writeString("[UNKNOWN]")
	}

	l.writer.writeByte(' ')

	if l.name != "" {
		l.writer.writeString(l.name)
		l.writer.writeString(": ")
	}

	if msg != "" {
		l.writer.writeString(msg)
	}

	l.writer.writeString("\n")

	l.writer.flush()
}

func (l *logger) trace(s string) {
	l.log(LevelTrace, s)
}

func (l *logger) debug(s string) {
	l.log(LevelDebug, s)
}

func (l *logger) info(s string) {
	l.log(LevelInfo, s)
}

func (l *logger) warn(s string) {
	l.log(LevelWarn, s)
}

func (l *logger) error(s string) {
	l.log(LevelError, s)
}
