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

import "testing"

func TestNavigatorHistory_PushPop(t *testing.T) {
	h := NewNavigatorHistory()

	// Pop on empty returns false
	_, ok := h.Pop()
	if ok {
		t.Fatal("Pop on empty history should return false")
	}

	h.Push(5)
	h.Push(10)
	h.Push(15)

	if h.Len() != 3 {
		t.Fatalf("expected Len() == 3, got %d", h.Len())
	}

	val, ok := h.Pop()
	if !ok || val != 15 {
		t.Fatalf("expected Pop() == (15, true), got (%d, %v)", val, ok)
	}

	val, ok = h.Pop()
	if !ok || val != 10 {
		t.Fatalf("expected Pop() == (10, true), got (%d, %v)", val, ok)
	}
}

func TestNavigatorHistory_BoundedAt20(t *testing.T) {
	h := NewNavigatorHistory()
	for i := 0; i < 30; i++ {
		h.Push(i)
	}
	if h.Len() != MaxNavigationHistory {
		t.Fatalf("expected Len() == %d after 30 pushes, got %d", MaxNavigationHistory, h.Len())
	}

	// Oldest entries (0-9) should have been trimmed; first pop should be 29
	val, ok := h.Pop()
	if !ok || val != 29 {
		t.Fatalf("expected Pop() == (29, true), got (%d, %v)", val, ok)
	}
}

func TestNavigatorHistory_Clear(t *testing.T) {
	h := NewNavigatorHistory()
	h.Push(1)
	h.Push(2)
	h.Clear()
	if h.Len() != 0 {
		t.Fatalf("expected Len() == 0 after Clear(), got %d", h.Len())
	}
	_, ok := h.Pop()
	if ok {
		t.Fatal("Pop after Clear should return false")
	}
}
