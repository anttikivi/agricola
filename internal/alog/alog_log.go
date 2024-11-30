package alog

import "github.com/anttikivi/agricola/internal/alog/severity"

// This code is derived from code in golang/glog, copyright 2023 Google Inc.
// It is licensed under the Apache License, version 2.0.
// You may obtain a copy of that license at
// https://www.apache.org/licenses/LICENSE-2.0

// This file contains the public logging functions for alog. They are in their
// separate file for clarity as there are a lot of them.

// verbosityLevel is the set verbosity level.
var verbosityLevel Level //nolint:gochecknoglobals

// Level specifies a level of verbosity.
type Level int

// Verbose is a boolean type that implements the logging functions and
// executes them if the logger has the correct verbosity level set.
type Verbose bool

// Info is equivalent to the global Info function, guarded by the value of v.
// See the documentation of V for usage.
func (v Verbose) Info(args ...any) {
	v.InfoDepth(1, args...)
}

// InfoDepth is equivalent to the global InfoDepth function, guarded by the value of v.
// See the documentation of V for usage.
func (v Verbose) InfoDepth(depth int, args ...any) {
	if v {
		logf(depth+1, severity.Info, formatToPrint(args), args...)
	}
}

// InfoDepthf is equivalent to the global InfoDepthf function, guarded by the value of v.
// See the documentation of V for usage.
func (v Verbose) InfoDepthf(depth int, format string, args ...any) {
	if v {
		logf(depth+1, severity.Info, format, args...)
	}
}

// Infoln is equivalent to the global Infoln function, guarded by the value of v.
// See the documentation of V for usage.
func (v Verbose) Infoln(args ...any) {
	if v {
		logf(1, severity.Info, formatToPrintln(args), args...)
	}
}

// Infof is equivalent to the global Infof function, guarded by the value of v.
// See the documentation of V for usage.
func (v Verbose) Infof(format string, args ...any) {
	if v {
		logf(1, severity.Info, format, args...)
	}
}

// V reports whether the verbosity level is set to at least the requested level.
// The returned value is a boolean of type verbose, which implements Info,
// Infoln and Infof.
// These methods will write to the INFO log if called.
// Thus, one may write either
//
//	if alog.V(2) { alog.Info("log this") }
//
// or
//
//	alog.V(2).Info("log this")
//
// The second form is shorter but the first is cheaper if logging is off because
// it does not evaluate its arguments.
//
// Whether an individual call to V generates a log record depends on the
// verbosityLevel variable that is set when alog is initialized.
// If the level in the call to V is at most the value of verbosityLevel, the V
// call will log.
func V(l Level) Verbose {
	return Verbose(verbosityLevel >= l)
}

// Info logs to the INFO log.
// Arguments are handled in the manner of fmt.Print; a newline is appended if missing.
func Info(args ...any) {
	InfoDepth(1, args...)
}

// InfoDepth calls Info from a different depth in the call stack.
// This enables a callee to emit logs that use the callsite information of its caller
// or any other callers in the stack. When depth == 0, the original callee's line
// information is emitted. When depth > 0, depth frames are skipped in the call stack
// and the final frame is treated like the original callee to Info.
func InfoDepth(depth int, args ...any) {
	logf(depth+1, severity.Info, formatToPrint(args), args...)
}

// InfoDepthf acts as InfoDepth but with format string.
func InfoDepthf(depth int, format string, args ...any) {
	logf(depth+1, severity.Info, format, args...)
}

// Infoln logs to the INFO log.
// Arguments are handled in the manner of fmt.Println; a newline is appended if missing.
func Infoln(args ...any) {
	logf(1, severity.Info, formatToPrintln(args), args...)
}

// Infof logs to the INFO log.
// Arguments are handled in the manner of fmt.Printf; a newline is appended if missing.
func Infof(format string, args ...any) {
	logf(1, severity.Info, format, args...)
}

// Warning logs to the WARNING and INFO logs.
// Arguments are handled in the manner of fmt.Print; a newline is appended if missing.
func Warning(args ...any) {
	WarningDepth(1, args...)
}

// WarningDepth acts as Warning but uses depth to determine which call frame to log.
// WarningDepth(0, "msg") is the same as Warning("msg").
func WarningDepth(depth int, args ...any) {
	logf(depth+1, severity.Warning, formatToPrint(args), args...)
}

// WarningDepthf acts as Warningf but uses depth to determine which call frame to log.
// WarningDepthf(0, "msg") is the same as Warningf("msg").
func WarningDepthf(depth int, format string, args ...any) {
	logf(depth+1, severity.Warning, format, args...)
}

// Warningln logs to the WARNING and INFO logs.
// Arguments are handled in the manner of fmt.Println; a newline is appended if missing.
func Warningln(args ...any) {
	logf(1, severity.Warning, formatToPrintln(args), args...)
}

// Warningf logs to the WARNING and INFO logs.
// Arguments are handled in the manner of fmt.Printf; a newline is appended if missing.
func Warningf(format string, args ...any) {
	logf(1, severity.Warning, format, args...)
}

// Error logs to the ERROR, WARNING, and INFO logs.
// Arguments are handled in the manner of fmt.Print; a newline is appended if missing.
func Error(args ...any) {
	ErrorDepth(1, args...)
}

// ErrorDepth acts as Error but uses depth to determine which call frame to log.
// ErrorDepth(0, "msg") is the same as Error("msg").
func ErrorDepth(depth int, args ...any) {
	logf(depth+1, severity.Error, formatToPrint(args), args...)
}

// ErrorDepthf acts as Errorf but uses depth to determine which call frame to log.
// ErrorDepthf(0, "msg") is the same as Errorf("msg").
func ErrorDepthf(depth int, format string, args ...any) {
	logf(depth+1, severity.Error, format, args...)
}

// Errorln logs to the ERROR, WARNING, and INFO logs.
// Arguments are handled in the manner of fmt.Println; a newline is appended if missing.
func Errorln(args ...any) {
	logf(1, severity.Error, formatToPrintln(args), args...)
}

// Errorf logs to the ERROR, WARNING, and INFO logs.
// Arguments are handled in the manner of fmt.Printf; a newline is appended if missing.
func Errorf(format string, args ...any) {
	logf(1, severity.Error, format, args...)
}
