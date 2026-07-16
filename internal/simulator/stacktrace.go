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

package simulator

type WasmStackTrace struct {
	TrapKind       interface{}  `json:"trap_kind"`
	RawMessage     string       `json:"raw_message"`
	Frames         []StackFrame `json:"frames"`
	SorobanWrapped bool         `json:"soroban_wrapped"`
}

type StackFrame struct {
	Index          int             `json:"index"`
	FuncIndex      *uint32         `json:"func_index,omitempty"`
	FuncName       *string         `json:"func_name,omitempty"`
	WasmOffset     *uint64         `json:"wasm_offset,omitempty"`
	Module         *string         `json:"module,omitempty"`
	SourceLocation *SourceLocation `json:"source_location,omitempty"`
}

type SourceLocation struct {
	File      string `json:"file"`
	Line      uint   `json:"line"`
	Column    uint   `json:"column"`
	ColumnEnd *uint  `json:"column_end,omitempty"`
}
