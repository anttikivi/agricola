package logging

import (
	"bytes"
	"strings"
)

type standardWriter struct {
	logger *logger
}

func (w *standardWriter) Write(p []byte) (int, error) {
	s := string(bytes.TrimRight(p, " \t\n"))
	l, s := getLevelFromMessage(s)

	switch l {
	case LevelTrace:
		w.logger.trace(s)
	case LevelDebug:
		w.logger.debug(s)
	case LevelInfo:
		w.logger.info(s)
	case LevelWarn:
		w.logger.warn(s)
	case LevelError:
		w.logger.error(s)
	default:
		w.logger.info(s)
	}
	return len(p), nil
}

func getLevelFromMessage(s string) (Level, string) {
	switch {
	case strings.HasPrefix(s, "[TRACE]"):
		return LevelTrace, strings.TrimSpace(s[len("[TRACE]"):])
	case strings.HasPrefix(s, "[DEBUG]"):
		return LevelDebug, strings.TrimSpace(s[len("[DEBUG]"):])
	case strings.HasPrefix(s, "[INFO]"):
		return LevelInfo, strings.TrimSpace(s[len("[INFO]"):])
	case strings.HasPrefix(s, "[WARN]"):
		return LevelWarn, strings.TrimSpace(s[len("[WARN]"):])
	case strings.HasPrefix(s, "[ERROR]"):
		return LevelError, strings.TrimSpace(s[len("[ERROR]"):])
	default:
		return LevelInfo, s
	}
}
