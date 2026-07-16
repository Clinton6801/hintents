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

	"github.com/stretchr/testify/assert"
)

func TestTraceNode_BudgetMetrics(t *testing.T) {
	node := NewTraceNode("test-node", "contract_call")

	// Initially, budget metrics should be nil
	assert.Nil(t, node.CPUDelta)
	assert.Nil(t, node.MemoryDelta)

	// Set budget metrics
	cpuDelta := uint64(150000)
	memDelta := uint64(2048)
	node.CPUDelta = &cpuDelta
	node.MemoryDelta = &memDelta

	assert.NotNil(t, node.CPUDelta)
	assert.NotNil(t, node.MemoryDelta)
	assert.Equal(t, uint64(150000), *node.CPUDelta)
	assert.Equal(t, uint64(2048), *node.MemoryDelta)
}

func TestCreateMockTrace_WithBudgetMetrics(t *testing.T) {
	root := CreateMockTrace()

	assert.NotNil(t, root)

	// Check that contract calls have budget metrics
	hasNodeWithBudget := false
	for _, child := range root.Children {
		if child.Type == "contract_call" && child.CPUDelta != nil && child.MemoryDelta != nil {
			hasNodeWithBudget = true
			assert.Greater(t, *child.CPUDelta, uint64(0), "CPU delta should be greater than 0")
			assert.Greater(t, *child.MemoryDelta, uint64(0), "Memory delta should be greater than 0")
		}
	}

	assert.True(t, hasNodeWithBudget, "Mock trace should have at least one node with budget metrics")
}

func TestTraceNode_BudgetMetrics_NestedNodes(t *testing.T) {
	// Create a tree with nested budget metrics
	root := NewTraceNode("root", "transaction")

	call1 := NewTraceNode("call-1", "contract_call")
	cpu1 := uint64(100000)
	mem1 := uint64(1024)
	call1.CPUDelta = &cpu1
	call1.MemoryDelta = &mem1
	root.AddChild(call1)

	call2 := NewTraceNode("call-2", "contract_call")
	cpu2 := uint64(200000)
	mem2 := uint64(2048)
	call2.CPUDelta = &cpu2
	call2.MemoryDelta = &mem2
	call1.AddChild(call2)

	// Flatten and check all nodes
	flattened := root.FlattenAll()

	assert.Equal(t, 3, len(flattened))

	// Root should not have budget (transaction node)
	assert.Nil(t, flattened[0].CPUDelta)

	// Child nodes should have budget
	assert.NotNil(t, flattened[1].CPUDelta)
	assert.Equal(t, uint64(100000), *flattened[1].CPUDelta)

	assert.NotNil(t, flattened[2].CPUDelta)
	assert.Equal(t, uint64(200000), *flattened[2].CPUDelta)
}
