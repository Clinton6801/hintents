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

package profile

import (
	"bytes"
	"testing"

	"github.com/dotandev/hintents/internal/trace"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTraceToPprof_NilTrace(t *testing.T) {
	_, err := TraceToPprof(nil)
	assert.Error(t, err)
}

func TestTraceToPprof_EmptyTrace(t *testing.T) {
	execTrace := trace.NewExecutionTrace("tx1", 10)
	p, err := TraceToPprof(execTrace)
	require.NoError(t, err)
	require.NotNil(t, p)
	assert.Len(t, p.SampleType, 1)
	assert.Equal(t, SampleTypeGas, p.SampleType[0].Type)
	assert.Equal(t, SampleUnitCount, p.SampleType[0].Unit)
	assert.Empty(t, p.Sample)
}

func TestTraceToPprof_WithGas(t *testing.T) {
	execTrace := trace.NewExecutionTrace("tx1", 10)
	execTrace.AddState(trace.ExecutionState{
		Step:       0,
		Operation:  "contract_call",
		ContractID: "CDLZFC3SYJYDZT7K67VZ75HPJVIEUVNIXF47ZG2FB2RMQQAHHAGCN4B2",
		Function:   "transfer",
		HostState:  map[string]interface{}{"gas_used": float64(15000)},
	})
	execTrace.AddState(trace.ExecutionState{
		Step:       1,
		Operation:  "contract_call",
		ContractID: "CDLZFC3SYJYDZT7K67VZ75HPJVIEUVNIXF47ZG2FB2RMQQAHHAGCN4B2",
		Function:   "mint",
		HostState:  map[string]interface{}{"gas_used": 8000},
	})

	p, err := TraceToPprof(execTrace)
	require.NoError(t, err)
	require.NotNil(t, p)
	assert.Len(t, p.SampleType, 1)
	assert.Equal(t, SampleTypeGas, p.SampleType[0].Type)
	require.Len(t, p.Sample, 2)
	assert.Equal(t, []int64{15000}, p.Sample[0].Value)
	assert.Equal(t, []int64{8000}, p.Sample[1].Value)
	assert.Len(t, p.Function, 2)
	assert.Len(t, p.Location, 2)
}

func TestTraceToPprof_WritePprof(t *testing.T) {
	execTrace := trace.NewExecutionTrace("tx1", 10)
	execTrace.AddState(trace.ExecutionState{
		Operation:  "contract_call",
		ContractID: "C1",
		Function:   "foo",
		HostState:  map[string]interface{}{"gas_used": int64(1000)},
	})

	var buf bytes.Buffer
	err := WritePprof(execTrace, &buf)
	require.NoError(t, err)
	assert.Greater(t, buf.Len(), 0)
}

func TestTraceToPprof_GasTypes(t *testing.T) {
	execTrace := trace.NewExecutionTrace("tx1", 10)
	execTrace.AddState(trace.ExecutionState{
		Operation: "call", ContractID: "C1", Function: "f",
		HostState: map[string]interface{}{"gas_used": float64(100)},
	})
	execTrace.AddState(trace.ExecutionState{
		Operation: "call", ContractID: "C1", Function: "g",
		HostState: map[string]interface{}{"gas_used": 200},
	})
	execTrace.AddState(trace.ExecutionState{
		Operation: "call", ContractID: "C1", Function: "h",
		HostState: map[string]interface{}{"gas_used": int64(300)},
	})

	p, err := TraceToPprof(execTrace)
	require.NoError(t, err)
	require.Len(t, p.Sample, 3)
	assert.Equal(t, []int64{100}, p.Sample[0].Value)
	assert.Equal(t, []int64{200}, p.Sample[1].Value)
	assert.Equal(t, []int64{300}, p.Sample[2].Value)
}
