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

package terminal

import (
	"os"
	"strings"
	"testing"
)

func unsetEnv(t *testing.T, key string) {
	t.Helper()
	value, exists := os.LookupEnv(key)
	_ = os.Unsetenv(key)
	t.Cleanup(func() {
		if exists {
			_ = os.Setenv(key, value)
			return
		}
		_ = os.Unsetenv(key)
	})
}

func TestANSIRenderer_IsTTY(t *testing.T) {
	t.Run("NO_COLOR disables color", func(t *testing.T) {
		t.Setenv("NO_COLOR", "1")
		r := NewANSIRenderer()
		if r.IsTTY() {
			t.Error("IsTTY() should be false when NO_COLOR is set")
		}
	})

	t.Run("FORCE_COLOR enables color", func(t *testing.T) {
		unsetEnv(t, "NO_COLOR")
		t.Setenv("FORCE_COLOR", "1")
		r := NewANSIRenderer()
		if !r.IsTTY() {
			t.Error("IsTTY() should be true when FORCE_COLOR is set")
		}
	})

	t.Run("TERM=dumb disables color", func(t *testing.T) {
		unsetEnv(t, "FORCE_COLOR")
		t.Setenv("TERM", "dumb")
		r := NewANSIRenderer()
		if r.IsTTY() {
			t.Error("IsTTY() should be false when TERM=dumb")
		}
	})
}

func TestANSIRenderer_Colorize(t *testing.T) {
	t.Run("colorized when forced", func(t *testing.T) {
		unsetEnv(t, "NO_COLOR")
		t.Setenv("FORCE_COLOR", "1")
		r := NewANSIRenderer()

		colored := r.Colorize("hello", "red")
		if !strings.Contains(colored, "[31m") {
			t.Errorf("Expected red color code, got %q", colored)
		}
	})

	t.Run("NO_COLOR overrides force", func(t *testing.T) {
		t.Setenv("FORCE_COLOR", "1")
		t.Setenv("NO_COLOR", "1")
		r := NewANSIRenderer()

		plain := r.Colorize("hello", "red")
		if strings.Contains(plain, "") {
			t.Errorf("Expected plain text when NO_COLOR is set, got %q", plain)
		}
	})
}

func TestANSIRenderer_Symbols(t *testing.T) {
	t.Run("forced color symbols", func(t *testing.T) {
		t.Setenv("FORCE_COLOR", "1")
		r := NewANSIRenderer()

		if r.Symbol("check") != "[OK]" {
			t.Errorf("Expected [OK] for check symbol, got %q", r.Symbol("check"))
		}
	})

	t.Run("NO_COLOR symbols", func(t *testing.T) {
		t.Setenv("NO_COLOR", "1")
		r := NewANSIRenderer()

		if r.Symbol("check") != "[OK]" {
			t.Errorf("Expected [OK] for check symbol when NO_COLOR, got %q", r.Symbol("check"))
		}
	})
}
