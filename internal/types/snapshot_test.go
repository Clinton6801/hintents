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
	"testing"
)

func TestStateSnapshotUnmarshal(t *testing.T) {
	t.Parallel()

	raw := []byte(`{
		"ledger_entries": {
			"AAAAAQAAAAAAAAA=": "AAAAAgAAAAAAAAA=",
			"AAAAAQAAAAAAAAE=": "AAAAAgAAAAAAAAE="
		},
		"timestamp": 1711651200,
		"instruction_index": 42,
		"events": [
			"ev:invoke:contract",
			"ev:return:ok"
		]
	}`)

	var got StateSnapshot
	if err := json.Unmarshal(raw, &got); err != nil {
		t.Fatalf("unmarshal StateSnapshot: %v", err)
	}

	if got.Timestamp != 1711651200 {
		t.Fatalf("unexpected timestamp: got %d", got.Timestamp)
	}

	if got.InstructionIndex != 42 {
		t.Fatalf("unexpected instruction_index: got %d", got.InstructionIndex)
	}

	if len(got.LedgerEntries) != 2 {
		t.Fatalf("unexpected ledger_entries length: got %d", len(got.LedgerEntries))
	}

	if got.LedgerEntries["AAAAAQAAAAAAAAA="] != "AAAAAgAAAAAAAAA=" {
		t.Fatalf("unexpected ledger entry value for first key")
	}

	if len(got.Events) != 2 {
		t.Fatalf("unexpected events length: got %d", len(got.Events))
	}
}

func TestLedgerDeltaUnmarshal(t *testing.T) {
	t.Parallel()

	raw := []byte(`{
		"new_keys": ["deadbeef", "cafebabe"],
		"modified_keys": ["010203"],
		"deleted_keys": ["f00d"]
	}`)

	var got LedgerDelta
	if err := json.Unmarshal(raw, &got); err != nil {
		t.Fatalf("unmarshal LedgerDelta: %v", err)
	}

	if len(got.NewKeys) != 2 || got.NewKeys[0] != "deadbeef" || got.NewKeys[1] != "cafebabe" {
		t.Fatalf("unexpected new_keys: %#v", got.NewKeys)
	}

	if len(got.ModifiedKeys) != 1 || got.ModifiedKeys[0] != "010203" {
		t.Fatalf("unexpected modified_keys: %#v", got.ModifiedKeys)
	}

	if len(got.DeletedKeys) != 1 || got.DeletedKeys[0] != "f00d" {
		t.Fatalf("unexpected deleted_keys: %#v", got.DeletedKeys)
	}
}
