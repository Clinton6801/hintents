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
	"bufio"
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestScaffoldErstProjectCreatesFilesAndDirs(t *testing.T) {
	root := t.TempDir()

	err := scaffoldErstProject(root, initScaffoldOptions{Network: "testnet"})
	require.NoError(t, err)

	for _, rel := range []string{
		"erst.toml",
		".gitignore",
		".erst/cache",
		".erst/snapshots",
		".erst/traces",
		"overrides",
		"wasm",
	} {
		_, statErr := os.Stat(filepath.Join(root, rel))
		assert.NoError(t, statErr, "expected %s to exist", rel)
	}

	erstToml, err := os.ReadFile(filepath.Join(root, "erst.toml"))
	require.NoError(t, err)
	assert.Contains(t, string(erstToml), `network = "testnet"`)
	assert.Contains(t, string(erstToml), `network_passphrase = "Test SDF Network ; September 2015"`)
	assert.Contains(t, string(erstToml), `cache_path = ".erst/cache"`)

	gitignore, err := os.ReadFile(filepath.Join(root, ".gitignore"))
	require.NoError(t, err)
	assert.Contains(t, string(gitignore), "# Erst local debugging artifacts")
	assert.Contains(t, string(gitignore), ".erst/traces/")
}

func TestScaffoldErstProjectDoesNotOverwriteErstTomlWithoutForce(t *testing.T) {
	root := t.TempDir()
	path := filepath.Join(root, "erst.toml")
	require.NoError(t, os.WriteFile(path, []byte("network = \"public\"
"), 0644))

	err := scaffoldErstProject(root, initScaffoldOptions{Network: "testnet"})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "erst.toml already exists")

	content, readErr := os.ReadFile(path)
	require.NoError(t, readErr)
	assert.Equal(t, "network = \"public\"
", string(content))
}

func TestEnsureGitignoreBlockIsIdempotent(t *testing.T) {
	root := t.TempDir()
	path := filepath.Join(root, ".gitignore")
	initial := "node_modules/
"
	require.NoError(t, os.WriteFile(path, []byte(initial), 0644))

	block := renderProjectGitignoreBlock()
	require.NoError(t, ensureGitignoreBlock(path, block))
	require.NoError(t, ensureGitignoreBlock(path, block))

	content, err := os.ReadFile(path)
	require.NoError(t, err)

	text := string(content)
	assert.True(t, strings.Contains(text, initial))
	assert.Equal(t, 1, strings.Count(text, "# Erst local debugging artifacts"))
}

func TestRenderProjectErstTomlStandaloneNetwork(t *testing.T) {
	content := renderProjectErstToml(initScaffoldOptions{Network: "standalone"})
	assert.Contains(t, content, `rpc_url = "http://localhost:8000"`)
	assert.Contains(t, content, `network = "standalone"`)
	assert.Contains(t, content, `network_passphrase = "Standalone Network ; February 2017"`)
}

func TestRenderProjectErstTomlWithOverrides(t *testing.T) {
	content := renderProjectErstToml(initScaffoldOptions{
		Network:           "testnet",
		RPCURL:            "https://rpc.example.org",
		NetworkPassphrase: "Example Network Passphrase",
	})

	assert.Contains(t, content, `rpc_url = "https://rpc.example.org"`)
	assert.Contains(t, content, `network_passphrase = "Example Network Passphrase"`)
}

func TestPromptWithDefaultUsesDefaultOnEmptyInput(t *testing.T) {
	reader := bufio.NewReader(strings.NewReader("
"))
	out := &bytes.Buffer{}

	value, err := promptWithDefault(reader, out, "Preferred Soroban RPC URL", "https://rpc.default")
	require.NoError(t, err)
	assert.Equal(t, "https://rpc.default", value)
	assert.Contains(t, out.String(), "Preferred Soroban RPC URL [https://rpc.default]: ")
}

func TestPromptWithDefaultUsesUserInput(t *testing.T) {
	reader := bufio.NewReader(strings.NewReader("https://rpc.custom
"))
	out := &bytes.Buffer{}

	value, err := promptWithDefault(reader, out, "Preferred Soroban RPC URL", "https://rpc.default")
	require.NoError(t, err)
	assert.Equal(t, "https://rpc.custom", value)
}
