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
	"testing"
)

func TestColorBlindThemes(t *testing.T) {
	_ = os.Setenv("FORCE_COLOR", "1")
	defer func() { _ = os.Unsetenv("FORCE_COLOR") }()

	tests := []struct {
		name  string
		theme Theme
	}{
		{"deuteranopia", ThemeDeuteranopia},
		{"protanopia", ThemeProtanopia},
		{"tritanopia", ThemeTritanopia},
		{"high-contrast", ThemeHighContrast},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetTheme(tt.theme)

			success := Success()
			if success == "" {
				t.Error("Success() returned empty string")
			}

			warning := Warning()
			if warning == "" {
				t.Error("Warning() returned empty string")
			}

			errorMsg := Error()
			if errorMsg == "" {
				t.Error("Error() returned empty string")
			}

			info := Info()
			if info == "" {
				t.Error("Info() returned empty string")
			}
		})
	}
}

func TestThemeConsistency(t *testing.T) {
	_ = os.Setenv("FORCE_COLOR", "1")
	defer func() { _ = os.Unsetenv("FORCE_COLOR") }()

	themes := []Theme{
		ThemeDefault,
		ThemeDeuteranopia,
		ThemeProtanopia,
		ThemeTritanopia,
		ThemeHighContrast,
	}

	for _, theme := range themes {
		t.Run(string(theme), func(t *testing.T) {
			SetTheme(theme)

			successColor := themeColors("success")
			errorColor := themeColors("error")
			warningColor := themeColors("warning")

			if successColor == errorColor {
				t.Error("success and error colors should be different")
			}
			if successColor == warningColor {
				t.Error("success and warning colors should be different")
			}
		})
	}
}
