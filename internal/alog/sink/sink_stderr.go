package sink

import (
	"fmt"
	"os"
	"sync"

	"github.com/anttikivi/agricola/internal/alog/severity"
)

// This code is derived from code in golang/glog, copyright 2023 Google Inc.
// It is licensed under the Apache License, version 2.0.
// You may obtain a copy of that license at
// https://www.apache.org/licenses/LICENSE-2.0

// Stderr is a Text sink that writes log entries to stderr.
type Stderr struct {
	lock sync.Mutex
}

func (s *Stderr) Enabled(m *Meta) bool {
	// TODO: Find a good way to implement a way to control the severity
	// threshold.
	return m.Severity >= severity.Info
}

func (s *Stderr) Emit(_ *Meta, p []byte) (int, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	n, err := os.Stderr.Write(p)
	if err != nil {
		return n, fmt.Errorf("failed to write to os.Stderr: %w", err)
	}

	return n, nil
}
