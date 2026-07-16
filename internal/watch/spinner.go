// SPDX-License-Identifier: Apache-2.0
// Copyright (c) 2026 dotandev
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.


// Copyright 2026 Erst Users

package watch

import (
	"fmt"
	"io"
	"os"
	"sync"
	"time"
)

type Spinner struct {
	frames    []string
	current   int
	done      chan struct{}
	mu        sync.Mutex
	isRunning bool
	out       io.Writer
}

func NewSpinner() *Spinner {
	return NewSpinnerWithWriter(os.Stdout)
}

func NewSpinnerWithWriter(w io.Writer) *Spinner {
	return &Spinner{
		frames: []string{"|", "/", "-", "\"},
		done:   make(chan struct{}),
		out:    w,
	}
}

func (s *Spinner) Start(message string) {
	s.mu.Lock()
	if s.isRunning {
		s.mu.Unlock()
		return
	}
	s.isRunning = true
	s.mu.Unlock()

	go func() {
		ticker := time.NewTicker(100 * time.Millisecond)
		defer ticker.Stop()

		for {
			select {
			case <-s.done:
				_, _ = fmt.Fprint(s.out, "[K")
				return
			case <-ticker.C:
				s.mu.Lock()
				_, _ = fmt.Fprintf(s.out, "%s %s", s.frames[s.current], message)
				s.current = (s.current + 1) % len(s.frames)
				s.mu.Unlock()
			}
		}
	}()
}

func (s *Spinner) Stop() {
	s.mu.Lock()
	if !s.isRunning {
		s.mu.Unlock()
		return
	}
	s.isRunning = false
	s.mu.Unlock()

	// Signal the goroutine to exit.
	select {
	case s.done <- struct{}{}:
	default:
		// If the channel is full, someone already signaled it
		// or the goroutine already exited.
	}

	// Give the goroutine a moment to receive and clear the line.
	// In a real implementation, we might want a sync.WaitGroup here.
	time.Sleep(20 * time.Millisecond)
}

func (s *Spinner) StopWithMessage(message string) {
	s.Stop()
	_, _ = fmt.Fprintf(s.out, "[OK] %s
", message)
}

func (s *Spinner) StopWithError(message string) {
	s.Stop()
	_, _ = fmt.Fprintf(s.out, "[ERROR] %s
", message)
}
