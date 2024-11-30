// Package alog implements leveled execution logs for Agricola.
//
// The package implements an interface similar to the standard Go log that is
// inspired by Google's glog package.
//
// The logging in Agricola is implemented as its own package to later extract it
// as a separate package to use across different Go project.
package alog

import (
	"os"
	"reflect"
	"runtime"
	"sync"
	"time"

	"github.com/anttikivi/agricola/internal/alog/severity"
	"github.com/anttikivi/agricola/internal/alog/sink"
)

// This code is derived from code in golang/glog, copyright 2023 Google Inc.
// It is licensed under the Apache License, version 2.0.
// You may obtain a copy of that license at
// https://www.apache.org/licenses/LICENSE-2.0

// ExitLogError is the exit code when the program exits for error in logging.
const ExitLogError = 5

// TODO: Find a good way to implement controls for this.
const registerStderrSink = true

// TODO: Is there a better way to implement this?
var pid = os.Getpid() //nolint:gochecknoglobals

// metaPool is a pool of *sink.Meta.
var metaPool sync.Pool //nolint:gochecknoglobals

// Init initializes the logging.
// The logger only outputs messages that are greater than or equal to the given
// severity.
// The log messages are written to output which be, for example, os.Stdout or
// a file. By default, the logger writes to os.Stderr.
func Init(verbosity Level) {
	registerSinks()

	verbosityLevel = verbosity
}

// formatToPrint returns a fmt.Printf format specifier that formats its
// arguments as if they were passed to fmt.Print.
func formatToPrint(args []any) string {
	n := len(args)
	switch n {
	case 0:
		return ""
	case 1:
		return "%v"
	}

	b := make([]byte, 0, n*3-1)
	wasString := true // Suppress leading space.

	for _, arg := range args {
		isString := arg != nil && reflect.TypeOf(arg).Kind() == reflect.String
		if wasString || isString {
			b = append(b, "%v"...)
		} else {
			b = append(b, " %v"...)
		}

		wasString = isString
	}

	return string(b)
}

// formatToPrintln returns a fmt.Printf format specifier that formats its
// arguments as if they were passed to fmt.Println.
func formatToPrintln(args []any) string {
	if len(args) == 0 {
		return "\n"
	}

	b := make([]byte, 0, len(args)*3) //nolint:mnd
	for range args {
		b = append(b, "%v "...)
	}

	b[len(b)-1] = '\n' // Replace the last space with a newline.

	return string(b)
}

// logf writes a log message for a log function call.
func logf(depth int, severity severity.Severity, format string, args ...any) {
	// Get the time right in the beginning.
	now := time.Now()

	_, file, line, ok := runtime.Caller(depth + 1)
	if !ok {
		file = "???"
		line = 1
	}

	metai, meta := metaFromPool()
	*meta = sink.Meta{
		Time:     now,
		File:     file,
		Line:     line,
		Depth:    depth + 1,
		Severity: severity,
		Thread:   int64(pid),
	}

	_, err := sink.Printf(meta, format, args...)
	if err != nil {
		sink.Printf(meta, "alog: exiting because of error: %v", err)
		// TODO: Flush files.
		os.Exit(ExitLogError)
	}

	metaPool.Put(metai)
}

// metaPoolGet returns a *sink.Meta from metaPool as both an interface and a
// pointer, allocating a new one if necessary. (Returning the interface value
// directly avoids an allocation if there was an existing pointer in the pool.)
func metaFromPool() (any, *sink.Meta) {
	if metai := metaPool.Get(); metai != nil {
		meta, ok := metai.(*sink.Meta)
		if !ok {
			panic("type assertion while getting *sink.Meta from meta pool failed")
		}

		return metai, meta
	}

	meta := new(sink.Meta)

	return meta, meta
}

func registerSinks() {
	if registerStderrSink {
		sink.TextSinks = append(sink.TextSinks, &sink.Stderr{})
	}
}
