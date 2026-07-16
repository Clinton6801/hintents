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

package main

import (
	"bytes"
	"errors"
	"testing"

	"github.com/dotandev/hintents/internal/cmd"
)

func TestRun_Interrupted(t *testing.T) {
	var stderr bytes.Buffer
	code := run(func() error { return cmd.ErrInterrupted }, &stderr)
	if code != cmd.InterruptExitCode {
		t.Fatalf("expected %d, got %d", cmd.InterruptExitCode, code)
	}
	if got := stderr.String(); got != "Interrupted. Shutting down...
" {
		t.Fatalf("unexpected stderr: %q", got)
	}
}

func TestRun_GenericError(t *testing.T) {
	var stderr bytes.Buffer
	code := run(func() error { return errors.New("boom") }, &stderr)
	if code != 1 {
		t.Fatalf("expected 1, got %d", code)
	}
	if got := stderr.String(); got != "Error: boom
" {
		t.Fatalf("unexpected stderr: %q", got)
	}
}

func TestRun_Success(t *testing.T) {
	var stderr bytes.Buffer
	code := run(func() error { return nil }, &stderr)
	if code != 0 {
		t.Fatalf("expected 0, got %d", code)
	}
	if stderr.Len() != 0 {
		t.Fatalf("expected empty stderr, got %q", stderr.String())
	}
}
