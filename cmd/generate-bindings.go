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
	"path/filepath"

	"github.com/dotandev/hintents/internal/bindings"
	"github.com/spf13/cobra"
)

var (
	bindingsOutput     string
	bindingsContractID string
	bindingsNetwork    string
	bindingsPackage    string
)

var generateBindingsCmd = &cobra.Command{
	Use:   "generate-bindings <wasm-file>",
	Short: "Generate TypeScript bindings for a Soroban smart contract",
	Long: `Generate strongly-typed TypeScript client bindings from a Soroban smart contract.

This command extracts the contract specification from the WASM file and generates
a TypeScript client that provides type-safe method calls with erst integration
for simulation and debugging.

Example:
  erst generate-bindings contract.wasm
  erst generate-bindings --output ./src/bindings --package my-contract contract.wasm
  erst generate-bindings --contract-id CDLZFC... --network testnet contract.wasm`,
	Args: cobra.ExactArgs(1),
	RunE: runGenerateBindings,
}

func runGenerateBindings(_ *cobra.Command, args []string) error {
	wasmPath := args[0]

	// Read WASM file
	wasmBytes, err := os.ReadFile(wasmPath)
	if err != nil {
		return fmt.Errorf("failed to read WASM file: %w", err)
	}

	// Set default output directory
	if bindingsOutput == "" {
		bindingsOutput = "."
	}

	// Set default package name from filename
	if bindingsPackage == "" {
		base := filepath.Base(wasmPath)
		bindingsPackage = base[:len(base)-len(filepath.Ext(base))]
	}

	// Create generator config
	config := bindings.GeneratorConfig{
		WasmBytes:   wasmBytes,
		OutputDir:   bindingsOutput,
		PackageName: bindingsPackage,
		ContractID:  bindingsContractID,
		Network:     bindingsNetwork,
	}

	// Generate bindings
	generator := bindings.NewGenerator(config)
	files, err := generator.Generate()
	if err != nil {
		return fmt.Errorf("failed to generate bindings: %w", err)
	}

	// Write files
	for _, file := range files {
		fullPath := filepath.Join(bindingsOutput, file.Path)

		// Create directory if needed
		dir := filepath.Dir(fullPath)
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("failed to create directory %s: %w", dir, err)
		}

		// Write file
		if err := os.WriteFile(fullPath, []byte(file.Content), 0644); err != nil {
			return fmt.Errorf("failed to write file %s: %w", fullPath, err)
		}

		fmt.Printf("Generated: %s
", fullPath)
	}

	fmt.Printf("
[OK] TypeScript bindings generated successfully
")
	fmt.Printf("Package: %s
", bindingsPackage)
	fmt.Printf("Output: %s
", bindingsOutput)

	return nil
}

func init() {
	generateBindingsCmd.Flags().StringVarP(&bindingsOutput, "output", "o", "", "Output directory (defaults to current directory)")
	generateBindingsCmd.Flags().StringVarP(&bindingsPackage, "package", "p", "", "Package name (defaults to WASM filename)")
	generateBindingsCmd.Flags().StringVar(&bindingsContractID, "contract-id", "", "Contract ID for network calls")
	generateBindingsCmd.Flags().StringVarP(&bindingsNetwork, "network", "n", "testnet", "Stellar network (testnet, mainnet, futurenet)")

	rootCmd.AddCommand(generateBindingsCmd)
}
