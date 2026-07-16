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

package rpc_test

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/dotandev/hintents/internal/rpc"
)

// ExampleClient_GetLedgerHeader demonstrates how to fetch ledger header information
func ExampleClient_GetLedgerHeader() {
	// Create a client for testnet
	client, _ := rpc.NewClient(rpc.WithNetwork(rpc.Testnet))

	// Create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Fetch ledger header for a specific sequence
	header, err := client.GetLedgerHeader(ctx, 12345678)
	if err != nil {
		// Handle different error types
		if rpc.IsLedgerNotFound(err) {
			log.Printf("Ledger not found: %v", err)
			return
		}
		if rpc.IsLedgerArchived(err) {
			log.Printf("Ledger archived: %v", err)
			return
		}
		if rpc.IsRateLimitError(err) {
			log.Printf("Rate limited, retry later: %v", err)
			return
		}
		log.Fatalf("Failed to fetch ledger: %v", err)
	}

	// Use the ledger header information
	fmt.Printf("Ledger %d:
", header.Sequence)
	fmt.Printf("  Hash: %s
", header.Hash)
	fmt.Printf("  Protocol Version: %d
", header.ProtocolVersion)
	fmt.Printf("  Close Time: %s
", header.CloseTime)
	fmt.Printf("  Base Fee: %d stroops
", header.BaseFee)
	fmt.Printf("  Transactions: %d successful, %d failed
",
		header.SuccessfulTxCount, header.FailedTxCount)
}

// ExampleClient_GetLedgerHeader_errorHandling demonstrates error handling patterns
func ExampleClient_GetLedgerHeader_errorHandling() {
	client, _ := rpc.NewClient(rpc.WithNetwork(rpc.Testnet))
	ctx := context.Background()

	header, err := client.GetLedgerHeader(ctx, 999999999)
	if err != nil {
		switch {
		case rpc.IsLedgerNotFound(err):
			// Ledger doesn't exist yet or is invalid
			fmt.Println("Ledger not found - may be in the future")
		case rpc.IsLedgerArchived(err):
			// Ledger is too old and has been archived
			fmt.Println("Ledger archived - try a more recent ledger")
		case rpc.IsRateLimitError(err):
			// Too many requests - implement backoff
			fmt.Println("Rate limited - waiting before retry")
			time.Sleep(5 * time.Second)
			// Retry logic here
		default:
			// Other errors (network, etc.)
			fmt.Printf("Error: %v
", err)
		}
		return
	}

	fmt.Printf("Ledger sequence: %d
", header.Sequence)
}

// ExampleClient_GetLedgerHeader_simulation demonstrates using ledger data for simulation
func ExampleClient_GetLedgerHeader_simulation() {
	client, _ := rpc.NewClient(rpc.WithNetwork(rpc.Testnet))
	ctx := context.Background()

	// Fetch the ledger where a transaction was executed
	ledgerSeq := uint32(12345678)
	header, err := client.GetLedgerHeader(ctx, ledgerSeq)
	if err != nil {
		log.Fatalf("Failed to fetch ledger: %v", err)
	}

	// Use ledger properties for simulation
	fmt.Printf("Simulating transaction at ledger %d:
", header.Sequence)
	fmt.Printf("  Timestamp: %s
", header.CloseTime)
	fmt.Printf("  Protocol: v%d
", header.ProtocolVersion)
	fmt.Printf("  Network state: %s total coins
", header.TotalCoins)

	// The HeaderXDR can be decoded for full ledger header details
	fmt.Printf("  Header XDR available: %d bytes
", len(header.HeaderXDR))
}

// ExampleNewClient demonstrates creating clients for different networks
func ExampleNewClient() {
	// Create a testnet client
	testnetClient, _ := rpc.NewClient(rpc.WithNetwork(rpc.Testnet))
	fmt.Printf("Testnet client created: %v
", testnetClient.Network)

	// Create a mainnet client
	mainnetClient, _ := rpc.NewClient(rpc.WithNetwork(rpc.Mainnet))
	fmt.Printf("Mainnet client created: %v
", mainnetClient.Network)

	// Create a futurenet client
	futurenetClient, _ := rpc.NewClient(rpc.WithNetwork(rpc.Futurenet))
	fmt.Printf("Futurenet client created: %v
", futurenetClient.Network)
}
