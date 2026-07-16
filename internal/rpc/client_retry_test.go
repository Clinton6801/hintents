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

package rpc

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"
	"time"
)

const validKey = "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA=="

func newRetryHTTPClient() *http.Client {
	cfg := RetryConfig{
		MaxRetries:         2,
		InitialBackoff:     1 * time.Millisecond,
		MaxBackoff:         2 * time.Millisecond,
		JitterFraction:     0,
		StatusCodesToRetry: []int{http.StatusTooManyRequests},
	}
	transport := NewRetryTransport(cfg, http.DefaultTransport)
	return &http.Client{Transport: transport}
}

func TestSimulateTransactionRetriesOnRateLimit(t *testing.T) {
	var calls int32
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if atomic.AddInt32(&calls, 1) == 1 {
			w.Header().Set("Retry-After", "1")
			w.WriteHeader(http.StatusTooManyRequests)
			return
		}
		resp := SimulateTransactionResponse{
			Jsonrpc: "2.0",
			ID:      1,
		}
		resp.Result.MinResourceFee = "1"
		resp.Result.TransactionData = "AAAA"
		_ = json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	client, err := NewClient(
		WithNetwork(Testnet),
		WithHorizonURL(server.URL),
		WithSorobanURL(server.URL),
		WithHTTPClient(newRetryHTTPClient()),
		WithCacheEnabled(false),
	)
	if err != nil {
		t.Fatalf("failed to build client: %v", err)
	}

	resp, err := client.SimulateTransaction(context.Background(), "AAAA")
	if err != nil {
		t.Fatalf("expected retry to succeed, got error: %v", err)
	}
	if resp.Result.MinResourceFee != "1" {
		t.Fatalf("unexpected response: %+v", resp.Result)
	}
	if atomic.LoadInt32(&calls) < 2 {
		t.Fatalf("expected at least 2 calls, got %d", atomic.LoadInt32(&calls))
	}
}

func TestGetLedgerEntriesRetriesOnRateLimit(t *testing.T) {
	var calls int32
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if atomic.AddInt32(&calls, 1) == 1 {
			w.Header().Set("Retry-After", "1")
			w.WriteHeader(http.StatusTooManyRequests)
			return
		}
		resp := GetLedgerEntriesResponse{
			Jsonrpc: "2.0",
			ID:      1,
		}
		validXdr := buildValidEntryB64(validKey)
		resp.Result.Entries = []LedgerEntryResult{{
			Key: validKey,
			Xdr: validXdr,
		}}
		_ = json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	client, err := NewClient(
		WithNetwork(Testnet),
		WithHorizonURL(server.URL),
		WithSorobanURL(server.URL),
		WithHTTPClient(newRetryHTTPClient()),
		WithCacheEnabled(false),
	)
	if err != nil {
		t.Fatalf("failed to build client: %v", err)
	}

	entries, err := client.GetLedgerEntries(context.Background(), []string{validKey})
	if err != nil {
		t.Fatalf("expected retry to succeed, got error: %v", err)
	}
	if _, ok := entries[validKey]; !ok {
		t.Fatalf("expected entry for validKey, got: %v", entries)
	}
	if atomic.LoadInt32(&calls) < 2 {
		t.Fatalf("expected at least 2 calls, got %d", atomic.LoadInt32(&calls))
	}
}
