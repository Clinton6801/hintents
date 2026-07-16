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

//go:build !windows

package simulator

import (
	"context"
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
	"time"
)

func TestRunnerRun_ContextCancelTerminatesProcess(t *testing.T) {
	simPath := filepath.Join(t.TempDir(), "fake-erst-sim.sh")
	script := "#!/bin/sh
trap '' TERM
sleep 30 &
child=$!
wait $child
"
	if err := os.WriteFile(simPath, []byte(script), 0755); err != nil {
		t.Fatalf("failed to write script: %v", err)
	}

	runner := &Runner{
		BinaryPath: simPath,
		activeCmds: make(map[*exec.Cmd]struct{}),
	}

	ctx, cancel := context.WithCancel(context.Background())
	done := make(chan error, 1)
	go func() {
		_, err := runner.Run(ctx, &SimulationRequest{
			EnvelopeXdr:   "x",
			ResultMetaXdr: "y",
		})
		done <- err
	}()

	time.Sleep(150 * time.Millisecond)
	cancel()

	select {
	case err := <-done:
		if !errors.Is(err, context.Canceled) {
			t.Fatalf("expected context canceled, got %v", err)
		}
	case <-time.After(5 * time.Second):
		t.Fatal("runner did not stop after context cancel")
	}

	if err := runner.Close(); err != nil {
		t.Fatalf("unexpected close error: %v", err)
	}
}
