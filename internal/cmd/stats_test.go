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
	"testing"

	"github.com/dotandev/hintents/internal/simulator"
)

func makeResponse(events []simulator.CategorizedEvent) *simulator.SimulationResponse {
	return &simulator.SimulationResponse{
		Status:            "success",
		CategorizedEvents: events,
	}
}

func TestBuildContractStats_Empty(t *testing.T) {
	resp := makeResponse(nil)
	stats := buildContractStats(resp)
	if len(stats) != 0 {
		t.Errorf("expected 0 stats, got %d", len(stats))
	}
}

func TestBuildContractStats_SingleContract(t *testing.T) {
	cid := "CONTRACT_A"
	resp := makeResponse([]simulator.CategorizedEvent{
		{EventType: "storage_write", ContractID: &cid},
		{EventType: "require_auth", ContractID: &cid},
	})

	stats := buildContractStats(resp)

	if len(stats) != 1 {
		t.Fatalf("expected 1 stat, got %d", len(stats))
	}

	s := stats[0]
	wantCost := uint64(costWeightStorageWrite + costWeightAuth)
	if s.estimatedCost != wantCost {
		t.Errorf("estimatedCost = %d, want %d", s.estimatedCost, wantCost)
	}
}

func TestBuildContractStats_Sorted(t *testing.T) {
	cheap := "B"
	expensive := "A"

	resp := makeResponse([]simulator.CategorizedEvent{
		{EventType: "contract_call", ContractID: &cheap},
		{EventType: "storage_write", ContractID: &expensive},
	})

	stats := buildContractStats(resp)

	if stats[0].contractID != expensive {
		t.Errorf("expected %s first, got %s", expensive, stats[0].contractID)
	}
}

func TestEventCost(t *testing.T) {
	cases := []struct {
		eventType string
		want      uint64
	}{
		{"storage_write", uint64(costWeightStorageWrite)},
		{"require_auth", uint64(costWeightAuth)},
		{"other", uint64(costWeightDefault)},
	}

	for _, tc := range cases {
		got := eventCost(tc.eventType)
		if got != tc.want {
			t.Errorf("eventCost(%q) = %d, want %d", tc.eventType, got, tc.want)
		}
	}
}
