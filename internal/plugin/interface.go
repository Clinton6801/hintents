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

package plugin

import "encoding/json"

// Version is the semantic versioning for the plugin API
const Version = "1.0.0"

// DecoderPlugin defines the interface for custom decoder plugins
type DecoderPlugin interface {
	// Name returns the plugin identifier
	Name() string

	// Version returns the plugin version following semver
	Version() string

	// CanDecode returns true if this plugin can handle the given event
	CanDecode(eventType string) bool

	// Decode processes the event and returns decoded data
	Decode(data []byte) (json.RawMessage, error)

	// Metadata returns plugin capabilities and requirements
	Metadata() Metadata
}

// Metadata describes plugin capabilities
type Metadata struct {
	Name        string   `json:"name"`
	Version     string   `json:"version"`
	APIVersion  string   `json:"api_version"`
	EventTypes  []string `json:"event_types"`
	Description string   `json:"description"`
}

// Factory creates a plugin instance
type Factory interface {
	Create() (DecoderPlugin, error)
}

// FactorySymbol is the exported symbol name for dynamic loading
const FactorySymbol = "NewPluginFactory"
