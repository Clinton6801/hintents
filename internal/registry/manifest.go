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

package registry

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// Manifest represents the erst.json project file that tracks dependencies.
type Manifest struct {
	Name         string            `json:"name"`
	Version      string            `json:"version"`
	Dependencies map[string]string `json:"dependencies"`
}

// LoadManifest reads the erst.json from the current directory.
func LoadManifest(dir string) (*Manifest, error) {
	path := filepath.Join(dir, "erst.json")
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("no erst.json found in %s", dir)
		}
		return nil, fmt.Errorf("failed to read erst.json: %w", err)
	}

	var m Manifest
	if err := json.Unmarshal(data, &m); err != nil {
		return nil, fmt.Errorf("failed to parse erst.json: %w", err)
	}

	if m.Dependencies == nil {
		m.Dependencies = make(map[string]string)
	}

	return &m, nil
}

// SaveManifest writes the manifest back to disk.
func SaveManifest(dir string, m *Manifest) error {
	path := filepath.Join(dir, "erst.json")
	data, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to serialize manifest: %w", err)
	}

	if err := os.WriteFile(path, data, 0644); err != nil {
		return fmt.Errorf("failed to write erst.json: %w", err)
	}

	return nil
}

// AddDependency adds a new package to the manifest dependencies.
func (m *Manifest) AddDependency(pkg string) error {
	ref, err := ParsePackage(pkg)
	if err != nil {
		return err
	}

	name := fmt.Sprintf("@%s/%s", ref.Org, ref.Repo)
	m.Dependencies[name] = ref.Version
	return nil
}
