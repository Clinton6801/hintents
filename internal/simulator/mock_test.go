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

package simulator

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseLedgerOverrideFlags(t *testing.T) {
	overrides, err := ParseLedgerOverrideFlags([]string{
		"key1:value1",
		"anotherKey:YmFzZTY0dGVzdA==",
	})
	assert.NoError(t, err)
	assert.Equal(t, "value1", overrides["key1"])
	assert.Equal(t, "YmFzZTY0dGVzdA==", overrides["anotherKey"])
}

func TestParseLedgerOverrideFlags_InvalidFormat(t *testing.T) {
	_, err := ParseLedgerOverrideFlags([]string{"invalid-format"})
	assert.Error(t, err)
}

func TestLoadLedgerOverrideManifest(t *testing.T) {
	tmpFile := filepath.Join(t.TempDir(), "ledger_override.json")
	override := LedgerOverrideManifest{
		LedgerEntries: map[string]string{
			"keyA": "valueA",
		},
	}
	data, err := json.MarshalIndent(override, "", "  ")
	assert.NoError(t, err)
	assert.NoError(t, os.WriteFile(tmpFile, data, 0644))

	entries, err := LoadLedgerOverrideManifest(tmpFile)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(entries))
	assert.Equal(t, "valueA", entries["keyA"])
}

func TestMergeLedgerOverrides(t *testing.T) {
	base := map[string]string{"a": "1", "b": "2"}
	overrides := map[string]string{"b": "over", "c": "3"}
	merged := MergeLedgerOverrides(base, overrides)
	assert.Equal(t, "1", merged["a"])
	assert.Equal(t, "over", merged["b"])
	assert.Equal(t, "3", merged["c"])
}
