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
	"strings"
)

// Event represents an emitted contract event during a contract call.
type Event struct {
	Type     string
	Metadata string
	Children []Event
}

// RenderEventTree renders a list of events as a structured ASCII tree.
func RenderEventTree(events []Event) string {
	if len(events) == 0 {
		return "No events recorded."
	}

	var sb strings.Builder
	sb.WriteString("Events
")

	renderNodes(&sb, events, "")

	return strings.TrimSuffix(sb.String(), "
")
}

// renderNodes is a recursive helper that builds the ASCII tree structure.
func renderNodes(sb *strings.Builder, events []Event, prefix string) {
	for i, event := range events {
		isLast := i == len(events)-1

		// Write the current level's prefix and branch character
		sb.WriteString(prefix)
		if isLast {
			sb.WriteString("└── ")
		} else {
			sb.WriteString("├── ")
		}

		// Write the event type
		sb.WriteString(event.Type)

		// Optionally write metadata if present
		if event.Metadata != "" {
			sb.WriteString(" (")
			sb.WriteString(event.Metadata)
			sb.WriteString(")")
		}

		sb.WriteByte('
')

		// Recursively render children with updated prefix
		if len(event.Children) > 0 {
			nextPrefix := prefix
			if isLast {
				nextPrefix += "    "
			} else {
				nextPrefix += "│   "
			}
			renderNodes(sb, event.Children, nextPrefix)
		}
	}
}
