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

package secutil

import "testing"

func TestMemzero(t *testing.T) {
	data := []byte{1, 2, 3, 4, 5}
	Memzero(data)
	for i, b := range data {
		if b != 0 {
			t.Errorf("index %d: got %d, want 0", i, b)
		}
	}
}

func TestMemzeroEmpty(t *testing.T) {
	Memzero([]byte{})
}

func TestMemzeroNil(t *testing.T) {
	Memzero(nil)
}

func TestMemzeroRetainsLength(t *testing.T) {
	data := []byte{0xDE, 0xAD, 0xBE, 0xEF}
	Memzero(data)
	if len(data) != 4 {
		t.Errorf("slice length changed: got %d, want 4", len(data))
	}
}
