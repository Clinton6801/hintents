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

package simulator

import (
	"context"
	"encoding/json"
	"os"
	"path/filepath"
	"runtime"
	"testing"
	"time"
)

// TestGenerateBaseline generates a new perf_baseline.json file
// This is not a regular test, it's a utility to initialize or update the baseline.
// Run it with: go test -v -run TestGenerateBaseline internal/simulator/generate_baseline_test.go internal/simulator/perf_regression_test.go internal/simulator/interface.go internal/simulator/mock_runner.go
func TestGenerateBaseline(t *testing.T) {
	if os.Getenv("GENERATE_BASELINE") != "true" {
		t.Skip("Skipping baseline generation. Set GENERATE_BASELINE=true to run.")
	}

	scenarios := map[string]func(context.Context, RunnerInterface) (*SimulationResponse, error){
		"EndlessLoop":      endlessLoopScenario,
		"MemoryHog":        memoryHogScenario,
		"CombinedWorkload": combinedWorkloadScenario,
	}

	baseline := PerfBaseline{
		Version:    1,
		Updated:    time.Now().Format(time.RFC3339),
		Benchmarks: make(map[string]BenchmarkEntry),
		Threshold:  10.0, // 10% threshold
	}

	for name, scenario := range scenarios {
		t.Logf("Running benchmark for %s...", name)
		result := runBenchmarkScenario(t, name, scenario)
		baseline.Benchmarks[name] = BenchmarkEntry{
			TargetNsPerOp:     result.NsPerOp,
			TargetBytesPerOp:  result.BytesPerOp,
			TargetAllocsPerOp: result.AllocsPerOp,
		}
		t.Logf("Result for %s: %d ns/op, %d B/op, %d allocs/op", name, result.NsPerOp, result.BytesPerOp, result.AllocsPerOp)
	}

	data, err := json.MarshalIndent(baseline, "", "  ")
	if err != nil {
		t.Fatalf("failed to marshal baseline: %v", err)
	}

	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("failed to get current file path")
	}
	dir := filepath.Dir(filename)
	baselinePath := filepath.Join(dir, "perf_baseline.json")

	if err := os.WriteFile(baselinePath, data, 0644); err != nil {
		t.Fatalf("failed to write baseline to %s: %v", baselinePath, err)
	}

	t.Logf("Successfully generated baseline at %s", baselinePath)
}
