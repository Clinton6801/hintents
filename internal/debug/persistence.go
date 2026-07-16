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

package debug

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"os"
)

// SaveToFile writes the Registry as gzip-compressed JSON to path.
// The caller should use FileExtension (".erstsnap") as the file suffix
// so that the file is recognisable as a snapshot registry.
func (r *Registry) SaveToFile(path string) error {
	data, err := json.Marshal(r)
	if err != nil {
		return fmt.Errorf("marshal registry: %w", err)
	}

	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("create %s: %w", path, err)
	}
	defer f.Close()

	gz := gzip.NewWriter(f)
	if _, err := gz.Write(data); err != nil {
		_ = gz.Close()
		return fmt.Errorf("write compressed registry: %w", err)
	}
	// Close must be called explicitly to flush the gzip footer.
	if err := gz.Close(); err != nil {
		return fmt.Errorf("close gzip writer: %w", err)
	}
	return nil
}

// LoadFromFile reads and decompresses a Registry from a gzip-compressed JSON
// file previously created by SaveToFile.
func LoadFromFile(path string) (*Registry, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("open %s: %w", path, err)
	}
	defer f.Close()

	gz, err := gzip.NewReader(f)
	if err != nil {
		return nil, fmt.Errorf("create gzip reader for %s: %w", path, err)
	}
	defer gz.Close()

	var reg Registry
	if err := json.NewDecoder(gz).Decode(&reg); err != nil {
		return nil, fmt.Errorf("decode registry from %s: %w", path, err)
	}
	return &reg, nil
}
