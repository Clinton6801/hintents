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

package bridge

import "fmt"

// MaxSnapshotSize is the Go-side threshold matching the Rust simulator's
// MAX_SNAPSHOT_SIZE (10 MB). Frames exceeding this size are treated as
// oversized and receive a safe fallback message.
const MaxSnapshotSize = 10 * 1024 * 1024

// FrameResult represents the outcome of processing a single simulation frame's
// snapshot data on the Go bridge side.
type FrameResult struct {
	// Step is the simulation step index.
	Step int `json:"step"`

	// Data holds the snapshot bytes when available.
	// nil when the frame was oversized and could not be captured.
	Data []byte `json:"data,omitempty"`

	// Oversized is true when the frame exceeded MaxSnapshotSize.
	Oversized bool `json:"oversized"`

	// Message is a human-readable status string.
	Message string `json:"message"`
}

// ParseFrameSnapshot inspects raw frame data from the simulator and returns
// a FrameResult. If the data exceeds MaxSnapshotSize the frame is flagged
// as oversized with a safe fallback message instead of attempting to
// deserialise (which could cause OOM or excessive allocation).
func ParseFrameSnapshot(step int, data []byte) FrameResult {
	if len(data) > MaxSnapshotSize {
		return FrameResult{
			Step:      step,
			Data:      nil,
			Oversized: true,
			Message:   fmt.Sprintf("Frame %d too large to capture (%d bytes, limit %d bytes)", step, len(data), MaxSnapshotSize),
		}
	}

	return FrameResult{
		Step:      step,
		Data:      data,
		Oversized: false,
		Message:   "Snapshot captured",
	}
}

// IsOversized is a convenience check for callers that just need to know
// whether a frame's snapshot is usable.
func (fr FrameResult) IsOversized() bool {
	return fr.Oversized
}
