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
	"os"
	"path/filepath"
	"testing"
)

func TestParseContractWasmOverrideSpecs(t *testing.T) {
	tmpDir := t.TempDir()
	wasmPath := filepath.Join(tmpDir, "bridge.wasm")
	if err := os.WriteFile(wasmPath, []byte{0x00, 0x61, 0x73, 0x6d}, 0644); err != nil {
		t.Fatalf("WriteFile: %v", err)
	}

	overrides, err := parseContractWasmOverrideSpecs([]string{
		"cafebabe=" + wasmPath,
	})
	if err != nil {
		t.Fatalf("parseContractWasmOverrideSpecs: %v", err)
	}
	if got := overrides["cafebabe"]; got != wasmPath {
		t.Fatalf("expected override path %q, got %q", wasmPath, got)
	}
}

func TestParseContractWasmOverrideSpecsRejectsInvalidSpec(t *testing.T) {
	if _, err := parseContractWasmOverrideSpecs([]string{"missing-separator"}); err == nil {
		t.Fatal("expected parseContractWasmOverrideSpecs to reject malformed override")
	}
}

func TestCloneStringMap(t *testing.T) {
	original := map[string]string{"a": "1"}
	cloned := cloneStringMap(original)
	cloned["a"] = "2"

	if original["a"] != "1" {
		t.Fatalf("expected original map to stay unchanged, got %q", original["a"])
	}
}
