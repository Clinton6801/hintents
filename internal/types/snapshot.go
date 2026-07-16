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

package types

import (
	"encoding/json"
	"fmt"

	"github.com/cespare/xxhash/v2"
	"github.com/dotandev/hintents/internal/snapshot"
)

// StateSnapshot mirrors the Rust simulator snapshot payload.
// Field order is kept consistent with the Rust struct for serde/bincode parity.
type StateSnapshot struct {
	LedgerEntries    map[string]string `json:"ledger_entries"`
	Timestamp        uint64            `json:"timestamp"`
	InstructionIndex uint32            `json:"instruction_index"`
	Events           []string          `json:"events"`
}

// LedgerDelta mirrors the Rust state::StateDiff JSON shape used by the UI.
type LedgerDelta struct {
	NewKeys      []string `json:"new_keys"`
	ModifiedKeys []string `json:"modified_keys"`
	DeletedKeys  []string `json:"deleted_keys"`
}

// Fingerprint returns a 64-bit xxHash digest of the canonical serialization of
// snap's ledger entries and linear memory.
//
// xxHash-64 is used rather than SHA-256 because registry verification runs
// once per entry on every load: at thousands of snapshots the difference in
// throughput is meaningful. The goal is integrity detection (corruption,
// truncation, accidental edit), not cryptographic resistance.
//
// The input to the hash is the JSON encoding of the snapshot as produced by
// json.Marshal. Because snapshot.FromMap sorts entries by key before
// construction, this output is deterministic across equivalent snapshots.
//
// A nil snapshot returns (0, nil).
func Fingerprint(snap *snapshot.Snapshot) (uint64, error) {
	if snap == nil {
		return 0, nil
	}
	data, err := json.Marshal(snap)
	if err != nil {
		return 0, fmt.Errorf("marshal snapshot for fingerprint: %w", err)
	}
	return xxhash.Sum64(data), nil
}
