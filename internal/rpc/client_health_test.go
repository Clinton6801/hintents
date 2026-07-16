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
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetHealth_Success(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)

		var req GetHealthRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		assert.NoError(t, err)
		assert.Equal(t, "getHealth", req.Method)

		resp := GetHealthResponse{
			Jsonrpc: "2.0",
			ID:      1,
			Result: struct {
				Status                string `json:"status"`
				LatestLedger          uint32 `json:"latestLedger"`
				OldestLedger          uint32 `json:"oldestLedger"`
				LedgerRetentionWindow uint32 `json:"ledgerRetentionWindow"`
			}{
				Status:                "healthy",
				LatestLedger:          100,
				OldestLedger:          1,
				LedgerRetentionWindow: 99,
			},
		}
		_ = json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	client := &Client{
		SorobanURL: server.URL,
		AltURLs:    []string{server.URL},
	}

	resp, err := client.GetHealth(context.Background())
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "healthy", resp.Result.Status)
	assert.Equal(t, uint32(100), resp.Result.LatestLedger)
}

func TestGetHealth_Error(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := GetHealthResponse{
			Jsonrpc: "2.0",
			ID:      1,
			Error: &struct {
				Code    int    `json:"code"`
				Message string `json:"message"`
			}{
				Code:    -32601,
				Message: "Method not found",
			},
		}
		_ = json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	client := &Client{
		SorobanURL: server.URL,
		AltURLs:    []string{server.URL},
	}

	resp, err := client.GetHealth(context.Background())
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "Method not found")
}

func TestGetHealth_Failover(t *testing.T) {
	server1 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server1.Close()

	server2 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := GetHealthResponse{
			Jsonrpc: "2.0",
			ID:      1,
			Result: struct {
				Status                string `json:"status"`
				LatestLedger          uint32 `json:"latestLedger"`
				OldestLedger          uint32 `json:"oldestLedger"`
				LedgerRetentionWindow uint32 `json:"ledgerRetentionWindow"`
			}{
				Status: "healthy",
			},
		}
		_ = json.NewEncoder(w).Encode(resp)
	}))
	defer server2.Close()

	client, _ := NewClient(
		WithNetwork(Testnet),
		WithAltURLs([]string{server1.URL, server2.URL}),
	)
	// Manually set SorobanURL to server1.URL for the first attempt
	client.SorobanURL = server1.URL

	resp, err := client.GetHealth(context.Background())
	assert.NoError(t, err)
	if assert.NotNil(t, resp) && assert.NotNil(t, resp.Result) {
		assert.Equal(t, "healthy", resp.Result.Status)
	}
}
