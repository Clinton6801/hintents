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

// Package secutil provides stateless security utility primitives.
package secutil

import "runtime"

// Memzero overwrites b with zeros to reduce the window during which sensitive
// key material is readable in a heap or core dump.
//
// The gc compiler does not currently elide stores to heap-allocated slices;
// runtime.KeepAlive provides an additional safety margin by signalling that b
// is still live at this point, preventing any future compiler optimisation from
// removing the zeroing loop.
//
// Callers are responsible for minimising the lifetime of any string containing
// the same key material (e.g. a hex-encoded private key), as Go strings are
// immutable and cannot be cleared by the caller.
func Memzero(b []byte) {
	for i := range b {
		b[i] = 0
	}
	runtime.KeepAlive(b)
}
