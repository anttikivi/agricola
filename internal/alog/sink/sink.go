// Package sink has the sinks for logging.
//
// TODO: Implement sink for structured data.
package sink

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/anttikivi/agricola/internal/alog/severity"
)

// This code is derived from code in golang/glog, copyright 2023 Google Inc.
// It is licensed under the Apache License, version 2.0.
// You may obtain a copy of that license at
// https://www.apache.org/licenses/LICENSE-2.0

// MaxLogMessageLen is the limit on length of a formatted log message, including
// the standard line prefix and trailing newline.
const MaxLogMessageLen = 15000

// digits is used in writing digits to buffers.
const digits = "0123456789"

// TextSinks contains the Text sink instances to which the logs are written.
// These are initialized in alog.Init.
var TextSinks []Text //nolint:gochecknoglobals

// buffers is a pool of *bytes.Buffer used in formatting log entries.
var buffers sync.Pool //nolint:gochecknoglobals

// Meta contains metadata about a logging call.
type Meta struct {
	// Time is the time when the logging call was made.
	Time time.Time

	// File is the source file where the logging call was made.
	File string

	// Line is the line offset within the source file.
	Line int

	// Depth is the number of stack frames between the sink and the logging
	// call.
	Depth int

	// Severity is the severity of the logging call.
	Severity severity.Severity

	// Thread is the thread ID.
	// This can be populated with a thread ID from another source, such as a
	// system we are importing logs from.
	// In the normal case, this will be set to the process ID (PID), since Go
	// doesn't have threads.
	Thread int64
}

// Text is a sink that accepts pre-formatted logging lines instead of structured
// data.
type Text interface {
	// Enabled returns whether this sink should output messages for the given
	// Meta.
	// If the sink returns false for a given Meta, the Printf function will not
	// call Emit on it for the corresponding log message.
	Enabled(m *Meta) bool

	// Emit writes a pre-formatted text log entry (including any applicable
	// header) to the log.
	// It returns the number of bytes occupied by the entry (which may differ
	// from the length of the passed-in slice).
	//
	// Emit returns any error encountered *if* it is severe enough that the log
	// package should terminate the process.
	//
	// The sink must not modify the *Meta parameter, nor reference it after
	// Printf has returned: it may be reused in subsequent calls.
	//
	// TODO: Should the implementations escape characters?
	Emit(m *Meta, p []byte) (n int, err error)
}

// Printf writes the logging entry to the registered Text sinks.
//
// The returned int is the maximum across all Emit and Printf calls.
// The returned error is the first non-nil error encountered.
// Sinks that are disabled by configuration should return (0, nil).
func Printf(m *Meta, format string, args ...any) (int, error) {
	m.Depth++
	n, err := printfTextSinks(m, TextSinks, format, args...)

	return n, err
}

// printfTextSinks formats a text log entry and emits it to all specified Text
// sinks.
//
// The returned int is the maximum across all Emit and Printf calls.
// The returned error is the first non-nil error encountered.
// Sinks that are disabled by configuration should return (0, nil).
func printfTextSinks(meta *Meta, textSinks []Text, format string, args ...any) (int, error) { //nolint:funlen
	// We expect at most file, stderr, and perhaps syslog. If there are more,
	// we'll end up allocating - no big deal.
	const maxExpectedTextSinks = 3

	var noAllocSinks [maxExpectedTextSinks]Text

	sinks := noAllocSinks[:0]

	for _, s := range textSinks {
		if s.Enabled(meta) {
			sinks = append(sinks, s)
		}
	}

	if len(sinks) == 0 && meta.Severity != severity.Fatal {
		return 0, nil // No Text sinks specified; don't bother formatting.
	}

	var buf *bytes.Buffer

	bufi := buffers.Get()
	if bufi == nil {
		buf = bytes.NewBuffer(nil)
		bufi = buf
	} else {
		var ok bool

		buf, ok = bufi.(*bytes.Buffer)
		if !ok {
			panic("type assertion while getting *bytes.Buffer from buffers pool failed")
		}

		buf.Reset()
	}

	// Lmmdd hh:mm:ss.uuuuuu PID/GID file:line]
	//
	// The "PID" entry arguably ought to be TID for consistency with other
	// environments, but TID is not meaningful in a Go program due to the
	// multiplexing of goroutines across threads.
	//
	// Avoid Fprintf, for speed. The format is so simple that we can do it
	// quickly by hand. It's worth about 3X. Fprintf is hard.
	const severityChar = "IWEF"

	buf.WriteByte(severityChar[meta.Severity])

	_, month, day := meta.Time.Date()
	hour, minute, second := meta.Time.Clock()

	writeTwoDigits(buf, int(month))
	writeTwoDigits(buf, day)
	buf.WriteByte(' ')
	writeTwoDigits(buf, hour)
	buf.WriteByte(':')
	writeTwoDigits(buf, minute)
	buf.WriteByte(':')
	writeTwoDigits(buf, second)
	buf.WriteByte('.')
	// TODO: Consider finding a (linter-)safe way to do the conversion.
	writeDigits(buf, 6, uint64(meta.Time.Nanosecond()/1000), '0') //nolint:gosec,mnd
	buf.WriteByte(' ')

	// TODO: Consider finding a (linter-)safe way to do the conversion.
	writeDigits(buf, 7, uint64(meta.Thread), ' ') //nolint:gosec,mnd
	buf.WriteByte(' ')

	file := meta.File
	if i := strings.LastIndex(file, "/"); i >= 0 {
		file = file[i+1:]
	}

	buf.WriteString(file)

	buf.WriteByte(':')

	var tmp [19]byte

	buf.Write(strconv.AppendInt(tmp[:0], int64(meta.Line), 10)) //nolint:mnd
	buf.WriteString("] ")

	// msgStart := buf.Len()
	fmt.Fprintf(buf, format, args...)

	if buf.Len() > MaxLogMessageLen-1 {
		buf.Truncate(MaxLogMessageLen - 1)
	}
	// msgEnd := buf.Len()

	if b := buf.Bytes(); b[len(b)-1] != '\n' {
		buf.WriteByte('\n')
	}

	var (
		n   = 0
		err error
	)

	for _, s := range sinks {
		sn, sErr := s.Emit(meta, buf.Bytes())
		if sn > n {
			n = sn
		}

		if sErr != nil && err == nil {
			err = sErr
		}
	}

	// TODO: For example, glog saves to a fatal message store.

	buffers.Put(bufi)

	return n, err
}

// writeTwoDigits formats a zero-prefixed two-digit integer to buf.
func writeTwoDigits(buf *bytes.Buffer, d int) {
	buf.WriteByte(digits[(d/10)%10]) //nolint:mnd
	buf.WriteByte(digits[d%10])
}

// writeDigits formats an n-digit integer to buf, padding with pad on the left.
// It assumes d != 0.
func writeDigits(buf *bytes.Buffer, n int, d uint64, pad byte) {
	var tmp [20]byte

	cutoff := len(tmp) - n
	j := len(tmp) - 1

	for ; d > 0; j-- {
		tmp[j] = digits[d%10]
		d /= 10
	}

	for ; j >= cutoff; j-- {
		tmp[j] = pad
	}

	j++

	buf.Write(tmp[j:])
}
