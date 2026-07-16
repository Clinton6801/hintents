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

// Issue #1273: Ledger Entry Mocking CLI & API
package simulator

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type LedgerOverrideManifest struct {
	LedgerEntries map[string]string `json:"ledger_entries,omitempty"`
}

func LoadLedgerOverrideManifest(path string) (map[string]string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var manifest LedgerOverrideManifest
	if err := json.Unmarshal(data, &manifest); err != nil {
		return nil, err
	}

	return manifest.LedgerEntries, nil
}

func ParseLedgerOverrideFlags(entries []string) (map[string]string, error) {
	overrides := make(map[string]string)
	for _, entry := range entries {
		parts := strings.SplitN(entry, ":", 2)
		if len(parts) != 2 || parts[0] == "" {
			return nil, fmt.Errorf("invalid ledger override format: %q, expected key:value", entry)
		}
		overrides[parts[0]] = parts[1]
	}

	return overrides, nil
}

func MergeLedgerOverrides(base map[string]string, overrides map[string]string) map[string]string {
	if len(overrides) == 0 {
		return base
	}

	if base == nil {
		base = make(map[string]string)
	}

	for key, value := range overrides {
		base[key] = value
	}

	return base
}
