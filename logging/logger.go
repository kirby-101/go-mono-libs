// Package logging provides a simple thread-safe logging facility.
package logging

import (
	"bufio"
	"os"
	"sync"

	ansicolor "github.com/kirby-101/go-mono-libs/ansi-color"
)

// Level represents the severity of a log message.
type Level int

const (
	defaultLogLevel = LogDebug

	// LogDebug represents debug messages.
	LogDebug Level = iota

	// LogInfo represents informational messages.
	LogInfo

	LogWarn

	// LogError represents errors.
	LogError

	// LogFatal represents fatal errors.
	LogFatal

	unknownFuncName = "??"
)

var (
	// logLevelDisplayNames gives the display name to use for our log levels.
	logLevelDisplayNames = map[Level]string{
		LogDebug: "debug",
		LogInfo:  "info",
		LogWarn:  "warn",
		LogError: "error",
		LogFatal: "fatal",
	}

	logLevelDisplayNamesColored = map[Level]string{
		LogDebug: "debug",
		LogInfo:  ansicolor.ColorizeBlueString("info"),
		LogWarn:  ansicolor.ColorizeYellowString("warn"),
		LogError: ansicolor.ColorizeRedString("error"),
		LogFatal: ansicolor.ColorizeRedString("fatal"),
	}
)

// Logger represents a thread-safe logger with level filtering and buffered output.
type Logger struct {
	Config Config
	Writer *bufio.Writer
	sync.Mutex
}

// NewLogger creates a new Logger instance writing to the given file (or stdout if nil).
func NewLogger(cfg Config) *Logger {
	if cfg.File == nil {
		cfg.File = os.Stdout
	}

	if cfg.Level == 0 {
		cfg.Level = defaultLogLevel
	}

	return &Logger{
		Config: cfg,
		Writer: bufio.NewWriter(cfg.File),
	}
}

// Close flushes buffered log output and closes the underlying file if present.
func (logger *Logger) Close() error {
	if err := logger.Writer.Flush(); err != nil {
		return &FlushError{Err: err}
	}

	if err := logger.Config.File.Close(); err != nil {
		return &CloseError{Err: err}
	}

	return nil
}
