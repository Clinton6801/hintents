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

package visualizer

import (
	"os"
	"strings"
	"testing"
)

func TestNoColorDisablesColors(t *testing.T) {
	_ = os.Setenv("NO_COLOR", "1")
	defer func() { _ = os.Unsetenv("NO_COLOR") }()

	if ColorEnabled() {
		t.Error("ColorEnabled() should be false when NO_COLOR is set")
	}

	out := Colorize("hello", "red")
	if strings.Contains(out, "") {
		t.Errorf("Colorize should not contain ANSI when NO_COLOR set, got: %q", out)
	}
	if out != "hello" {
		t.Errorf("Colorize should return plain text, got: %q", out)
	}
}

func TestTermDumbDisablesColors(t *testing.T) {
	oldTerm := os.Getenv("TERM")
	defer func() { _ = os.Setenv("TERM", oldTerm) }()
	_ = os.Unsetenv("NO_COLOR")
	_ = os.Setenv("TERM", "dumb")

	if ColorEnabled() {
		t.Error("ColorEnabled() should be false when TERM=dumb")
	}
}

func TestSymbolReturnsPlainASCIIWhenDisabled(t *testing.T) {
	_ = os.Setenv("NO_COLOR", "1")
	defer func() { _ = os.Unsetenv("NO_COLOR") }()

	for name, wantPlain := range map[string]string{
		"check":   "[OK]",
		"cross":   "[X]",
		"warn":    "[!]",
		"arrow_r": "->",
	} {
		got := Symbol(name)
		if got != wantPlain {
			t.Errorf("Symbol(%q) = %q, want %q (NO_COLOR should force plain ASCII)", name, got, wantPlain)
		}
	}
}

func TestSuccessWarningErrorNoEscapeWhenDisabled(t *testing.T) {
	_ = os.Setenv("NO_COLOR", "1")
	defer func() { _ = os.Unsetenv("NO_COLOR") }()

	for _, s := range []string{Success(), Warning(), Error()} {
		if strings.Contains(s, "") {
			t.Errorf("Output contains ANSI escape when disabled: %q", s)
		}
	}
}

func TestNoColorOverridesForceColor(t *testing.T) {
	_ = os.Setenv("NO_COLOR", "1")
	_ = os.Setenv("FORCE_COLOR", "1")
	defer func() {
		_ = os.Unsetenv("NO_COLOR")
		_ = os.Unsetenv("FORCE_COLOR")
	}()

	if ColorEnabled() {
		t.Error("NO_COLOR must take precedence over FORCE_COLOR")
	}
	out := Colorize("test", "red")
	if out != "test" {
		t.Errorf("NO_COLOR+FORCE_COLOR: expected plain text, got %q", out)
	}
}

func TestForceColorEnablesColorsWhenSet(t *testing.T) {
	_ = os.Unsetenv("NO_COLOR")
	_ = os.Setenv("FORCE_COLOR", "1")
	defer func() { _ = os.Unsetenv("FORCE_COLOR") }()

	// FORCE_COLOR=1 should enable colors (even when piped / not TTY)
	if !ColorEnabled() {
		t.Error("ColorEnabled() should be true when FORCE_COLOR=1 and NO_COLOR unset")
	}
	out := Colorize("hello", "red")
	if !strings.Contains(out, "") {
		t.Errorf("FORCE_COLOR=1: Colorize should emit ANSI, got plain: %q", out)
	}
}

func TestContractBoundaryPlainText(t *testing.T) {
	_ = os.Setenv("NO_COLOR", "1")
	defer func() { _ = os.Unsetenv("NO_COLOR") }()

	out := ContractBoundary("CABC", "CXYZ")
	expected := "--- contract boundary: CABC -> CXYZ ---"
	if out != expected {
		t.Errorf("ContractBoundary() = %q, want %q", out, expected)
	}
	if strings.Contains(out, "") {
		t.Errorf("ContractBoundary should not contain ANSI when NO_COLOR set, got: %q", out)
	}
}

func TestContractBoundaryWithColor(t *testing.T) {
	_ = os.Unsetenv("NO_COLOR")
	_ = os.Setenv("FORCE_COLOR", "1")
	defer func() { _ = os.Unsetenv("FORCE_COLOR") }()

	out := ContractBoundary("CABC", "CXYZ")
	if !strings.Contains(out, "CABC") || !strings.Contains(out, "CXYZ") {
		t.Errorf("ContractBoundary should contain both contract IDs, got: %q", out)
	}
	if !strings.Contains(out, "") {
		t.Errorf("ContractBoundary should contain ANSI codes when colors enabled, got: %q", out)
	}
}
