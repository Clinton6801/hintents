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

package bindings

import (
	"testing"

	"github.com/stellar/go-stellar-sdk/xdr"
)

func TestMapTypeDefToTS(t *testing.T) {
	g := &Generator{}

	tests := []struct {
		name     string
		typeDef  xdr.ScSpecTypeDef
		expected string
	}{
		{
			name:     "Bool",
			typeDef:  xdr.ScSpecTypeDef{Type: xdr.ScSpecTypeScSpecTypeBool},
			expected: "boolean",
		},
		{
			name:     "U64",
			typeDef:  xdr.ScSpecTypeDef{Type: xdr.ScSpecTypeScSpecTypeU64},
			expected: "bigint",
		},
		{
			name:     "String",
			typeDef:  xdr.ScSpecTypeDef{Type: xdr.ScSpecTypeScSpecTypeString},
			expected: "string",
		},
		{
			name:     "Address",
			typeDef:  xdr.ScSpecTypeDef{Type: xdr.ScSpecTypeScSpecTypeAddress},
			expected: "Address",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := g.mapTypeDefToTS(tt.typeDef)
			if result != tt.expected {
				t.Errorf("mapTypeDefToTS() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestToPascalCase(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello-world", "HelloWorld"},
		{"my_contract", "MyContract"},
		{"simple", "Simple"},
		{"multi-word-test", "MultiWordTest"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := toPascalCase(tt.input)
			if result != tt.expected {
				t.Errorf("toPascalCase(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}
