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

package trace

// MaxNavigationHistory is the maximum number of entries kept in the undo stack.
const MaxNavigationHistory = 20

// NavigatorHistory tracks navigation steps for undo support.
// The stack is bounded to MaxNavigationHistory entries; when full,
// the oldest entry is discarded.
type NavigatorHistory struct {
	stack []int
}

// NewNavigatorHistory creates an empty navigation history.
func NewNavigatorHistory() *NavigatorHistory {
	return &NavigatorHistory{
		stack: make([]int, 0, MaxNavigationHistory),
	}
}

// Push records a navigation index. If the stack exceeds MaxNavigationHistory,
// the oldest entry is trimmed.
func (h *NavigatorHistory) Push(index int) {
	h.stack = append(h.stack, index)
	if len(h.stack) > MaxNavigationHistory {
		trimmed := make([]int, len(h.stack)-1, MaxNavigationHistory)
		copy(trimmed, h.stack[1:])
		h.stack = trimmed
	}
}

// Pop removes and returns the most recent navigation index.
// Returns -1 and false if the stack is empty.
func (h *NavigatorHistory) Pop() (int, bool) {
	if len(h.stack) == 0 {
		return -1, false
	}
	last := h.stack[len(h.stack)-1]
	h.stack = h.stack[:len(h.stack)-1]
	return last, true
}

// Len returns the current number of entries in the history.
func (h *NavigatorHistory) Len() int {
	return len(h.stack)
}

// Clear removes all entries from the history.
func (h *NavigatorHistory) Clear() {
	h.stack = h.stack[:0]
}
