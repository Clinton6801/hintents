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
	"sync"
	"testing"
)

// TestMarkFailureSuccessRace verifies that concurrent calls to markHorizonFailure,
// markHorizonSuccess, markSorobanFailure, and markSorobanSuccess do not race.
// Run with: go test -race ./internal/rpc/
func TestMarkFailureSuccessRace(t *testing.T) {
	c := &Client{
		HorizonURL: "https://horizon.example.com",
		SorobanURL: "https://soroban.example.com",
	}

	const goroutines = 50
	var wg sync.WaitGroup
	wg.Add(goroutines * 4)

	for i := 0; i < goroutines; i++ {
		go func() { defer wg.Done(); c.markFailure(c.HorizonURL) }()
		go func() { defer wg.Done(); c.markSuccess(c.HorizonURL) }()
		go func() { defer wg.Done(); c.markFailure(c.SorobanURL) }()
		go func() { defer wg.Done(); c.markSuccess(c.SorobanURL) }()
	}

	wg.Wait()
}
