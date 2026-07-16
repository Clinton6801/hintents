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

package audit

import (
	"encoding/json"
	"fmt"
)

// Dump is the raw {input, state, events} JSON payload produced by AuditLogger.
type Dump struct {
	Input     map[string]interface{} `json:"input"`
	State     map[string]interface{} `json:"state"`
	Events    []interface{}          `json:"events"`
	Timestamp string                 `json:"timestamp"`
}

// SignedDump extends Dump with signing metadata (matches SignedAuditLog from TS).
type SignedDump struct {
	Trace     Dump   `json:"trace"`
	Hash      string `json:"hash"`
	Signature string `json:"signature"`
	Algorithm string `json:"algorithm"`
	PublicKey string `json:"publicKey"`
	Signer    struct {
		Provider string `json:"provider"`
	} `json:"signer"`
}

// ParseDump deserialises raw JSON into an Dump.
func ParseDump(data []byte) (*Dump, error) {
	var d Dump
	if err := json.Unmarshal(data, &d); err != nil {
		return nil, fmt.Errorf("failed to parse audit dump: %w", err)
	}
	return &d, nil
}

// ParseSignedDump deserialises raw JSON into a SignedDump.
func ParseSignedDump(data []byte) (*SignedDump, error) {
	var d SignedDump
	if err := json.Unmarshal(data, &d); err != nil {
		return nil, fmt.Errorf("failed to parse signed audit dump: %w", err)
	}
	return &d, nil
}
