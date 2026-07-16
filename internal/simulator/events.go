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
	"encoding/base64"
	"fmt"

	"github.com/stellar/go-stellar-sdk/xdr"
)

type DiagnosticEvent struct {
	EventType                string   `json:"event_type"`
	ContractID               *string  `json:"contract_id,omitempty"`
	Topics                   []string `json:"topics"`
	Data                     string   `json:"data"`
	InSuccessfulContractCall bool     `json:"in_successful_contract_call"`
	WasmInstruction          *string  `json:"wasm_instruction,omitempty"`
	CPU                      *uint64  `json:"cpu,omitempty"`
	Memory                   *uint64  `json:"mem,omitempty"`
}

// ParseData decodes the base64-encoded XDR Data into an xdr.ScVal
func (e *DiagnosticEvent) ParseData() (xdr.ScVal, error) {
	var val xdr.ScVal
	if e.Data == "" {
		return val, nil
	}
	raw, err := base64.StdEncoding.DecodeString(e.Data)
	if err != nil {
		return val, fmt.Errorf("decode data base64: %w", err)
	}
	if err := xdr.SafeUnmarshal(raw, &val); err != nil {
		return val, fmt.Errorf("unmarshal data xdr: %w", err)
	}
	return val, nil
}

// ParseTopics decodes the base64-encoded XDR Topics into a slice of xdr.ScVal
func (e *DiagnosticEvent) ParseTopics() ([]xdr.ScVal, error) {
	var vals []xdr.ScVal
	for i, t := range e.Topics {
		var val xdr.ScVal
		raw, err := base64.StdEncoding.DecodeString(t)
		if err != nil {
			return nil, fmt.Errorf("decode topic[%d] base64: %w", i, err)
		}
		if err := xdr.SafeUnmarshal(raw, &val); err != nil {
			return nil, fmt.Errorf("unmarshal topic[%d] xdr: %w", i, err)
		}
		vals = append(vals, val)
	}
	return vals, nil
}

type CategorizedEvent struct {
	Category                 string   `json:"category"`
	EventType                string   `json:"event_type"`
	ContractID               *string  `json:"contract_id,omitempty"`
	Topics                   []string `json:"topics"`
	Data                     string   `json:"data"`
	InSuccessfulContractCall bool     `json:"in_successful_contract_call"`
	WasmInstruction          *string  `json:"wasm_instruction,omitempty"`
	CPU                      *uint64  `json:"cpu,omitempty"`
	Memory                   *uint64  `json:"memory,omitempty"`
}
