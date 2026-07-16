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

package cmd

import (
	"fmt"
	"os"

	"github.com/dotandev/hintents/internal/wasmopt"
)

func optimizeWasmBytesIfRequested(input []byte, enabled bool) ([]byte, *wasmopt.Report, error) {
	if !enabled {
		return input, nil, nil
	}
	optimized, report, err := wasmopt.EliminateDeadCode(input)
	if err != nil {
		return nil, nil, err
	}
	return optimized, &report, nil
}

func optimizeWasmFileIfRequested(path string, enabled bool) (string, *wasmopt.Report, func(), error) {
	cleanup := func() {}
	if !enabled {
		return path, nil, cleanup, nil
	}

	raw, err := os.ReadFile(path)
	if err != nil {
		return "", nil, cleanup, err
	}
	optimized, report, err := wasmopt.EliminateDeadCode(raw)
	if err != nil {
		return "", nil, cleanup, err
	}

	tmp, err := os.CreateTemp("", "erst-opt-*.wasm")
	if err != nil {
		return "", nil, cleanup, err
	}
	tmpPath := tmp.Name()
	if _, err := tmp.Write(optimized); err != nil {
		_ = tmp.Close()
		_ = os.Remove(tmpPath)
		return "", nil, cleanup, err
	}
	if err := tmp.Close(); err != nil {
		_ = os.Remove(tmpPath)
		return "", nil, cleanup, err
	}

	cleanup = func() { _ = os.Remove(tmpPath) }
	return tmpPath, &report, cleanup, nil
}

func printOptimizationReport(report *wasmopt.Report) {
	if report == nil {
		return
	}
	fmt.Printf(
		"Optimization (DCE): removed %d/%d functions, kept %d
",
		report.RemovedDefinedFunctions,
		report.OriginalDefinedFunctions,
		report.KeptDefinedFunctions,
	)
}
