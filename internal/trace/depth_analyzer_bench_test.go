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

import (
	"testing"
)

func BenchmarkAnalyzeDepth(b *testing.B) {
	root := createDeepTrace(100)
	da := NewDepthAnalyzer(10)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		da.AnalyzeDepth(root)
	}
}

func BenchmarkOptimizeForDisplay(b *testing.B) {
	root := createDeepTrace(100)
	da := NewDepthAnalyzer(10)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		da.OptimizeForDisplay(root)
	}
}

func BenchmarkFocusOnErrors(b *testing.B) {
	root := createDeepTrace(100)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FocusOnErrors(root)
	}
}

func createDeepTrace(depth int) *TraceNode {
	root := NewTraceNode("root", "transaction")
	current := root

	for i := 0; i < depth; i++ {
		child := NewTraceNode("node-"+string(rune(i)), "contract_call")
		child.ContractID = "CONTRACT"
		child.Function = "call"
		cpu := uint64(1000)
		mem := uint64(512)
		child.CPUDelta = &cpu
		child.MemoryDelta = &mem
		current.AddChild(child)
		current = child
	}

	errorNode := NewTraceNode("error", "error")
	errorNode.Error = "Deep error"
	current.AddChild(errorNode)

	return root
}
