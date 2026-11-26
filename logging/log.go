package logging

import (
	"bytes"
	"fmt"
	"os"
	"runtime"
	"time"
)

// Log writes a message at the given log level with the given type and message parts.
func (logger *Logger) Log(level Level, logType string, messageParts ...string) {
	// check if we log on the right level
	if level < logger.Config.Level {
		return
	}

	var (
		logLevelStr string
		rawBuf      bytes.Buffer
	)

	if logger.Config.Colorize {
		logLevelStr = logLevelDisplayNamesColored[level]
	} else {
		logLevelStr = logLevelDisplayNames[level]
	}

	// assemble our log line
	fmt.Fprintf(
		&rawBuf,
		"[%s]-[%s]-[%s] ", time.Now().Format("02.01.06_15:04:05"), logLevelStr, logType,
	)

	for i, p := range messageParts {
		rawBuf.WriteString(p)

		if i != len(messageParts)-1 {
			rawBuf.WriteString(" : ")
		}
	}

	rawBuf.WriteRune('\n')
	logger.Lock()

	// output line
	if _, err := logger.Writer.Write(rawBuf.Bytes()); err != nil {
		panic(err)
	}

	if err := logger.Writer.Flush(); err != nil {
		panic(err)
	}

	logger.Unlock()
}

// Debug logs a debug-level message including the caller function name.
func (logger *Logger) Debug(messageParts ...string) {
	pc, _, _, ok := runtime.Caller(1)
	funcname := runtime.FuncForPC(pc).Name()

	if !ok {
		funcname = unknownFuncName
	}

	logger.Log(LogDebug, funcname, messageParts...)
}

// Info logs an info-level message including the caller function name.
func (logger *Logger) Info(messageParts ...string) {
	pc, _, _, ok := runtime.Caller(1)
	funcname := runtime.FuncForPC(pc).Name()

	if !ok {
		funcname = unknownFuncName
	}

	logger.Log(LogInfo, funcname, messageParts...)
}

// Warn logs an warn-level message including the caller function name.
func (logger *Logger) Warn(messageParts ...string) {
	pc, _, _, ok := runtime.Caller(1)
	funcname := runtime.FuncForPC(pc).Name()

	if !ok {
		funcname = unknownFuncName
	}

	logger.Log(LogWarn, funcname, messageParts...)
}

// Error logs an error-level message including the caller function name.
func (logger *Logger) Error(messageParts ...string) {
	pc, _, _, ok := runtime.Caller(1)
	funcname := runtime.FuncForPC(pc).Name()

	if !ok {
		funcname = unknownFuncName
	}

	logger.Log(LogError, funcname, messageParts...)
}

// Fatal logs a fatal error message including the caller function name and exits the program.
func (logger *Logger) Fatal(messageParts ...string) {
	pc, _, _, ok := runtime.Caller(1)
	funcname := runtime.FuncForPC(pc).Name()

	if !ok {
		funcname = unknownFuncName
	}

	logger.Log(LogFatal, funcname, messageParts...)
	os.Exit(1)
}
