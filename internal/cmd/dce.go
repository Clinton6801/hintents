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

	"github.com/dotandev/hintents/internal/dce"
	"github.com/spf13/cobra"
)

var dceOutput string

var dceCmd = &cobra.Command{
	Use:   "dce <wasm-file>",
	Short: "Eliminate dead code from a WASM binary",
	Long: `Analyze a compiled WASM binary, build a call graph from exported functions,
and strip unreachable functions to reduce contract size.

Without -o, performs a dry run and prints statistics only.

Examples:
  erst dce ./contract.wasm -o ./contract-optimized.wasm
  erst dce ./contract.wasm`,
	Args: cobra.ExactArgs(1),
	RunE: dceExec,
}

func dceExec(cmd *cobra.Command, args []string) error {
	wasmBytes, err := os.ReadFile(args[0])
	if err != nil {
		return fmt.Errorf("reading WASM file: %w", err)
	}

	out, stats, err := dce.Eliminate(wasmBytes)
	if err != nil {
		return err
	}

	fmt.Printf("Total functions:    %d
", stats.TotalFunctions)
	fmt.Printf("Removed functions:  %d
", stats.RemovedFunctions)
	fmt.Printf("Original size:      %d bytes
", stats.OriginalSize)
	fmt.Printf("Optimized size:     %d bytes
", stats.OptimizedSize)

	if stats.OriginalSize > 0 {
		saved := stats.OriginalSize - stats.OptimizedSize
		pct := float64(saved) / float64(stats.OriginalSize) * 100
		fmt.Printf("Saved:              %d bytes (%.1f%%)
", saved, pct)
	}

	if dceOutput == "" {
		return nil
	}

	if err := os.WriteFile(dceOutput, out, 0644); err != nil {
		return fmt.Errorf("writing output: %w", err)
	}
	fmt.Printf("Written to:         %s
", dceOutput)

	return nil
}

func init() {
	dceCmd.Flags().StringVarP(&dceOutput, "output", "o", "", "Output file path (omit for dry run)")
	rootCmd.AddCommand(dceCmd)
}
