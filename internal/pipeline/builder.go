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

package pipeline

import (
	"encoding/json"
	"fmt"
)

// Pipeline simulates a Programmable Transaction Block (PTB).
type Pipeline struct {
	Commands []Command `json:"commands"`
}

// Command represents a single operation in a pipeline.
type Command struct {
	Type     string            `json:"type"`   // "MoveCall", "TransferObjects", etc.
	Target   string            `json:"target"` // e.g. "0x2::coin::mint"
	Args     []string          `json:"args"`   // arguments or references to previous results
	Metadata map[string]string `json:"metadata"`
}

// NewBuilder creates a new pipeline builder.
func NewBuilder() *Pipeline {
	return &Pipeline{
		Commands: make([]Command, 0),
	}
}

// AddCommand appends a command to the pipeline.
func (p *Pipeline) AddCommand(cmdType, target string, args []string) {
	p.Commands = append(p.Commands, Command{
		Type:   cmdType,
		Target: target,
		Args:   args,
	})
}

// ToJSON serializes the pipeline to JSON.
func (p *Pipeline) ToJSON() (string, error) {
	b, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// FromJSON parses a pipeline from JSON.
func FromJSON(data []byte) (*Pipeline, error) {
	var p Pipeline
	if err := json.Unmarshal(data, &p); err != nil {
		return nil, fmt.Errorf("failed to parse pipeline JSON: %w", err)
	}
	return &p, nil
}
