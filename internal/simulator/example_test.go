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

package simulator_test

import (
	"fmt"

	"github.com/dotandev/hintents/internal/simulator"
)

// ExampleSimulationRequestBuilder demonstrates basic usage of the builder pattern.
func ExampleSimulationRequestBuilder() {
	req, err := simulator.NewSimulationRequestBuilder().
		WithEnvelopeXDR("AAAAAgAAAACE...").
		WithResultMetaXDR("AAAAAQAAAAA...").
		Build()

	if err != nil {
		fmt.Printf("Error: %v
", err)
		return
	}

	fmt.Printf("Request created with envelope XDR length: %d
", len(req.EnvelopeXdr))
	// Output: Request created with envelope XDR length: 15
}

// ExampleSimulationRequestBuilder_withLedgerEntries demonstrates adding ledger entries.
func ExampleSimulationRequestBuilder_withLedgerEntries() {
	req, err := simulator.NewSimulationRequestBuilder().
		WithEnvelopeXDR("AAAAAgAAAACE...").
		WithResultMetaXDR("AAAAAQAAAAA...").
		WithLedgerEntry("a2V5MQ==", "dmFsdWUx").
		WithLedgerEntry("a2V5Mg==", "dmFsdWUy").
		Build()

	if err != nil {
		fmt.Printf("Error: %v
", err)
		return
	}

	fmt.Printf("Request created with %d ledger entries
", len(req.LedgerEntries))
	// Output: Request created with 2 ledger entries
}

// ExampleSimulationRequestBuilder_bulkLedgerEntries demonstrates setting multiple entries at once.
func ExampleSimulationRequestBuilder_bulkLedgerEntries() {
	entries := map[string]string{
		"Y29udHJhY3Rfa2V5XzE=": "Y29udHJhY3RfdmFsdWVfMQ==",
		"Y29udHJhY3Rfa2V5XzI=": "Y29udHJhY3RfdmFsdWVfMg==",
		"Y29udHJhY3Rfa2V5XzM=": "Y29udHJhY3RfdmFsdWVfMw==",
	}

	req, err := simulator.NewSimulationRequestBuilder().
		WithEnvelopeXDR("AAAAAgAAAACE...").
		WithResultMetaXDR("AAAAAQAAAAA...").
		WithLedgerEntries(entries).
		Build()

	if err != nil {
		fmt.Printf("Error: %v
", err)
		return
	}

	fmt.Printf("Request created with %d ledger entries
", len(req.LedgerEntries))
	// Output: Request created with 3 ledger entries
}

// ExampleSimulationRequestBuilder_validation demonstrates validation errors.
func ExampleSimulationRequestBuilder_validation() {
	_, err := simulator.NewSimulationRequestBuilder().
		WithEnvelopeXDR("AAAAAgAAAACE...").
		Build()

	if err != nil {
		fmt.Println("Validation error:", err)
	}
	// Output: Validation error: VALIDATION_FAILED: result meta XDR is required
}

// ExampleSimulationRequestBuilder_reuse demonstrates builder reuse with Reset().
func ExampleSimulationRequestBuilder_reuse() {
	builder := simulator.NewSimulationRequestBuilder()

	req1, _ := builder.
		WithEnvelopeXDR("envelope1").
		WithResultMetaXDR("result1").
		Build()

	fmt.Printf("First request envelope: %s
", req1.EnvelopeXdr)

	req2, _ := builder.
		Reset().
		WithEnvelopeXDR("envelope2").
		WithResultMetaXDR("result2").
		Build()

	fmt.Printf("Second request envelope: %s
", req2.EnvelopeXdr)
	// Output:
	// First request envelope: envelope1
	// Second request envelope: envelope2
}
