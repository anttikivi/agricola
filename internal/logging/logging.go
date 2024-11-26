package logging

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"syscall"
)

// Level represents a logging level.
type Level int

const (
	// LevelDefault indicates that no logging level is set and the default is used.
	LevelDefault Level = iota
	LevelTrace
	LevelDebug
	LevelInfo
	LevelWarn
	LevelError
	LevelOff
)

const (
	// defaultLevel is the default logging level to use.
	// TODO: This will be switched to LevelOff when in production.
	defaultLevel       Level       = LevelTrace
	defaultLevelString             = "TRACE"
	defaultLogPerm     os.FileMode = 0o666
	defaultTimeFormat              = "2006-01-02T15:04:05.000Z0700"
)

const (
	// envVariableLog is the name of the environment variable that sets the logging level.
	envVariableLog = "AGER_LOG"

	// envVariableLogPath is the name of the environment variable that sets the output file for the logs.
	envVariableLogPath = "AGER_LOG_PATH"
)

// Init initializes the Go log package for use with the program.
// The functions in the log package are used for actual logging and not the output printed to the user.
func Init() {
	logLevel := LevelTrace
	envLogLevel := strings.ToUpper(strings.TrimSpace(os.Getenv(envVariableLog)))

	if isValidLevelString(envLogLevel) || envLogLevel == "" {
		logLevel = parseLoggingLevel(envLogLevel)
		if logLevel == LevelDefault {
			logLevel = defaultLevel
		}
	} else {
		fmt.Fprintf(
			os.Stderr,
			"[%s] Invalid logging level: %q. Defaulting to %s. Valid logging levels are: %+v",
			getStringForLevel(LevelWarn),
			envLogLevel,
			getStringForLevel(defaultLevel),
			getValidLevels(),
		)
	}

	logOutput := io.Writer(os.Stderr)

	if logPath := os.Getenv(envVariableLogPath); logPath != "" {
		f, err := os.OpenFile(logPath, syscall.O_CREAT|syscall.O_RDWR|syscall.O_APPEND, defaultLogPerm)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error opening the log file: %v\n", err)
		} else {
			logOutput = f
		}
	}

	logger := createLogger(logLevel, logOutput)
	logWriter := logger.createLogWriter()

	log.SetFlags(0)
	log.SetPrefix("")
	log.SetOutput(logWriter)
}

func getStringForLevel(l Level) string {
	switch l {
	case LevelTrace:
		return "TRACE"
	case LevelDebug:
		return "DEBUG"
	case LevelInfo:
		return "INFO"
	case LevelWarn:
		return "WARN"
	case LevelError:
		return "ERROR"
	case LevelOff:
		return "OFF"
	case LevelDefault:
		return defaultLevelString
	default:
		return defaultLevelString
	}
}

func getValidLevels() []string {
	return []string{"trace", "debug", "info", "warn", "error", "off"}
}

func isValidLevelString(level string) bool {
	s := strings.ToLower(level)
	for _, l := range getValidLevels() {
		if s == l {
			return true
		}
	}

	return false
}

func parseLoggingLevel(s string) Level {
	if s == "" {
		return defaultLevel
	}

	switch strings.ToLower(s) {
	case "trace":
		return LevelTrace
	case "debug":
		return LevelDebug
	case "info":
		return LevelInfo
	case "warn":
		return LevelWarn
	case "error":
		return LevelError
	case "off":
		return LevelOff
	default:
		return LevelDefault
	}
}
