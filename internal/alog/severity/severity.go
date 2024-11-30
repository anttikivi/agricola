// Package severity contains utilities for log severity.
package severity

// This code is derived from code in golang/glog, copyright 2023 Google Inc.
// It is licensed under the Apache License, version 2.0.
// You may obtain a copy of that license at
// https://www.apache.org/licenses/LICENSE-2.0

// A Severity is a severity at which a message can be logged.
type Severity int8

// These constants identify the log levels in order of increasing severity.
const (
	Info Severity = iota
	Warning
	Error
	Fatal
	// TODO: Think about implementing level for panics.
)
