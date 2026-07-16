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

//go:build windows || plan9 || js || wasip1

package trace

import "os"

// getTermWidthSys returns 0 on platforms where TIOCGWINSZ is unavailable.
// getTermWidth falls back to COLUMNS or 80.
func getTermWidthSys() int { return 0 }

// watchResize is a no-op on Windows and plan9 (no SIGWINCH equivalent).
func watchResize(_ chan<- os.Signal) {}
