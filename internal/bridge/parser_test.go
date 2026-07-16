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

import "testing"

func TestParseFrameSnapshot_Normal(t *testing.T) {
	data := make([]byte, 1024)
	result := ParseFrameSnapshot(0, data)

	if result.Oversized {
		t.Fatal("expected non-oversized frame")
	}
	if result.Data == nil {
		t.Fatal("expected data to be set")
	}
	if result.Message != "Snapshot captured" {
		t.Fatalf("unexpected message: %s", result.Message)
	}
}

func TestParseFrameSnapshot_Oversized(t *testing.T) {
	data := make([]byte, MaxSnapshotSize+1)
	result := ParseFrameSnapshot(42, data)

	if !result.Oversized {
		t.Fatal("expected oversized frame")
	}
	if result.Data != nil {
		t.Fatal("expected data to be nil for oversized frame")
	}
	if !result.IsOversized() {
		t.Fatal("IsOversized() should return true")
	}
}

func TestParseFrameSnapshot_ExactLimit(t *testing.T) {
	data := make([]byte, MaxSnapshotSize)
	result := ParseFrameSnapshot(1, data)

	if result.Oversized {
		t.Fatal("frame at exact limit should not be oversized")
	}
}
