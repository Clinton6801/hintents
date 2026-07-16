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

package tests

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/dotandev/hintents/internal/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestConfigPriority verifies that:
// 1. Env vars override config file
// 2. We can simulate Flag overrides (conceptually, though testing main flags is harder in unit tests,
// we can test the precedence logic if we extract it, or rely on end-to-end.
// Here we verify at least Config vs Env using the same primitives used in command)
func TestConfigPriority(t *testing.T) {
	// Setup generic temporary home directory for config
	tmpDir := t.TempDir()

	// Mock home directory to point to tmpDir so config.LoadConfig uses it
	// On Unix, os.UserHomeDir() uses HOME; on Windows, it uses USERPROFILE
	originalHome := os.Getenv("HOME")
	originalUserProfile := os.Getenv("USERPROFILE")
	defer func() {
		os.Setenv("HOME", originalHome)
		os.Setenv("USERPROFILE", originalUserProfile)
	}()
	os.Setenv("HOME", tmpDir)
	os.Setenv("USERPROFILE", tmpDir)

	// Create a dummy TOML config file
	configFile := filepath.Join(tmpDir, ".erst.toml")
	configData := []byte(`rpc_token = "CONFIG_TOKEN"`)
	err := os.WriteFile(configFile, configData, 0600)
	require.NoError(t, err)

	// Helper to resolve token with same logic as debug.go
	resolveToken := func(flagVal string) string {
		token := flagVal
		if token == "" {
			token = os.Getenv("ERST_RPC_TOKEN")
		}
		if token == "" {
			cfg, err := config.Load()
			if err == nil && cfg.RPCToken != "" {
				token = cfg.RPCToken
			}
		}
		return token
	}

	t.Run("ConfigOnly", func(t *testing.T) {
		// Unset ENV
		os.Unsetenv("ERST_RPC_TOKEN")

		token := resolveToken("")
		assert.Equal(t, "CONFIG_TOKEN", token, "Should load from config file when flag and env are empty")
	})

	t.Run("EnvOverridesConfig", func(t *testing.T) {
		os.Setenv("ERST_RPC_TOKEN", "ENV_TOKEN")
		defer os.Unsetenv("ERST_RPC_TOKEN")

		token := resolveToken("")
		assert.Equal(t, "ENV_TOKEN", token, "Env var should override config file")
	})

	t.Run("FlagOverridesAll", func(t *testing.T) {
		os.Setenv("ERST_RPC_TOKEN", "ENV_TOKEN")
		defer os.Unsetenv("ERST_RPC_TOKEN")

		token := resolveToken("FLAG_TOKEN")
		assert.Equal(t, "FLAG_TOKEN", token, "Flag should override env and config")
	})
}
