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

package terminal

// Renderer defines the interface for terminal drawing and interaction.
type Renderer interface {
	Print(a ...any)
	Printf(format string, a ...any)
	Println(a ...any)
	Colorize(text, color string) string
	Success() string
	Warning() string
	Error() string
	Symbol(name string) string
	ClearLine()
	Scanln(a ...any) (int, error)
	IsTTY() bool
}
