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

package protocolreg

import "testing"

func TestParseDebugURI(t *testing.T) {
	parsed, err := ParseDebugURI("erst://debug/0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef?network=testnet&operation=2&source=dashboard")
	if err != nil {
		t.Fatalf("ParseDebugURI returned error: %v", err)
	}

	if parsed.TransactionHash != "0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef" {
		t.Fatalf("unexpected transaction hash: %s", parsed.TransactionHash)
	}
	if parsed.Network != "testnet" {
		t.Fatalf("unexpected network: %s", parsed.Network)
	}
	if parsed.Operation == nil || *parsed.Operation != 2 {
		t.Fatalf("unexpected operation: %#v", parsed.Operation)
	}
	if parsed.Source != "dashboard" {
		t.Fatalf("unexpected source: %s", parsed.Source)
	}
}

func TestParseDebugURIRejectsInvalidValues(t *testing.T) {
	tests := []string{
		"",
		"https://example.com",
		"erst://decode/0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef?network=testnet",
		"erst://debug/not-a-hash?network=testnet",
		"erst://debug/0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef",
		"erst://debug/0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef?network=invalid",
		"erst://debug/0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef?network=testnet&operation=-1",
	}

	for _, tc := range tests {
		if _, err := ParseDebugURI(tc); err == nil {
			t.Fatalf("expected ParseDebugURI to fail for %q", tc)
		}
	}
}
