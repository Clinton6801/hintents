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
	"context"
	"testing"
)

func TestSearchCommand_RequiresQueryArg(t *testing.T) {
	if err := searchCmd.Args(searchCmd, []string{}); err == nil {
		t.Fatal("expected missing query arg to fail")
	}
	if err := searchCmd.Args(searchCmd, []string{"usdc"}); err != nil {
		t.Fatalf("expected one query arg to pass: %v", err)
	}
}

func TestSearchCommand_InvalidNetwork(t *testing.T) {
	prevNetwork := searchNetworkFlag
	t.Cleanup(func() {
		searchNetworkFlag = prevNetwork
	})
	searchNetworkFlag = "invalid-network"

	err := searchCmd.RunE(searchCmd, []string{"usdc"})
	if err == nil {
		t.Fatal("expected invalid network error")
	}
}

func TestSearchCommand_EmptyQueryRejectedByRPCLayer(t *testing.T) {
	prevNetwork := searchNetworkFlag
	prevHorizonURL := searchHorizonURLFlag
	t.Cleanup(func() {
		searchNetworkFlag = prevNetwork
		searchHorizonURLFlag = prevHorizonURL
	})

	searchNetworkFlag = "testnet"
	searchHorizonURLFlag = "http://127.0.0.1:1"

	// Ensure we call RunE directly (cobra args validation is tested separately).
	searchCmd.SetContext(context.Background())
	err := searchCmd.RunE(searchCmd, []string{""})
	if err == nil {
		t.Fatal("expected empty query to fail")
	}
}
