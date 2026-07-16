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
	"fmt"
	"testing"
)

func BenchmarkSearchLargeTrace(b *testing.B) {
	// Create trace with 10,000 nodes
	nodes := make([]*TraceNode, 10000)
	for i := range nodes {
		nodes[i] = &TraceNode{
			ID:         fmt.Sprintf("node-%d", i),
			ContractID: fmt.Sprintf("contract-%d", i),
			Function:   fmt.Sprintf("function_%d", i%100),
			Error:      "",
			EventData:  fmt.Sprintf("event data %d", i),
		}
	}

	engine := NewSearchEngine()
	engine.SetQuery("function_42")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		engine.Search(nodes)
	}
}

func BenchmarkSearchSmallTrace(b *testing.B) {
	// Create trace with 100 nodes
	nodes := make([]*TraceNode, 100)
	for i := range nodes {
		nodes[i] = &TraceNode{
			ID:         fmt.Sprintf("node-%d", i),
			ContractID: fmt.Sprintf("contract-%d", i),
			Function:   fmt.Sprintf("function_%d", i%10),
		}
	}

	engine := NewSearchEngine()
	engine.SetQuery("function_5")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		engine.Search(nodes)
	}
}

func BenchmarkHighlightMatches(b *testing.B) {
	engine := NewSearchEngine()
	engine.SetQuery("error")

	node := &TraceNode{
		ID:    "test",
		Error: "This is an error message with multiple error occurrences and error handling",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		engine.HighlightMatches(node, "error")
	}
}

func BenchmarkNavigateMatches(b *testing.B) {
	nodes := make([]*TraceNode, 1000)
	for i := range nodes {
		nodes[i] = &TraceNode{
			ID:       fmt.Sprintf("node-%d", i),
			Function: "test_function",
		}
	}

	engine := NewSearchEngine()
	engine.SetQuery("test")
	engine.Search(nodes)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		engine.NextMatch()
	}
}

func BenchmarkFlattenLargeTree(b *testing.B) {
	// Create deep tree
	root := NewTraceNode("root", "transaction")
	current := root

	for i := 0; i < 1000; i++ {
		child := NewTraceNode(fmt.Sprintf("node-%d", i), "contract_call")
		current.AddChild(child)
		current = child
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		root.Flatten()
	}
}

func BenchmarkCaseInsensitiveSearch(b *testing.B) {
	nodes := make([]*TraceNode, 1000)
	for i := range nodes {
		nodes[i] = &TraceNode{
			ID:    fmt.Sprintf("node-%d", i),
			Error: "ERROR message with Error and error",
		}
	}

	engine := NewSearchEngine()
	engine.SetQuery("error")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		engine.Search(nodes)
	}
}

func BenchmarkCaseSensitiveSearch(b *testing.B) {
	nodes := make([]*TraceNode, 1000)
	for i := range nodes {
		nodes[i] = &TraceNode{
			ID:    fmt.Sprintf("node-%d", i),
			Error: "ERROR message with Error and error",
		}
	}

	engine := NewSearchEngine()
	engine.caseSensitive = true
	engine.SetQuery("error")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		engine.Search(nodes)
	}
}
