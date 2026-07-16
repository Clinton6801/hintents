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

package visualizer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDescribeHostFunction(t *testing.T) {
	doc := DescribeHostFunction("require_auth")
	assert.Contains(t, doc, "ensure")

	nonexistent := DescribeHostFunction("unknown_fn")
	assert.Contains(t, nonexistent, "host function")
}

func TestFormatGasSummary(t *testing.T) {
	assert.Equal(t, "CPU: 150000 instructions · Memory: 2048 bytes", FormatGasSummary(150000, 2048))
}

func TestEstimateGasHint(t *testing.T) {
	high := EstimateGasHint(250000, 8192)
	assert.Contains(t, high, "High resource usage")

	low := EstimateGasHint(30000, 512)
	assert.Contains(t, low, "Low resource usage")
}

func TestDiagnosticsForSource(t *testing.T) {
	source := "require_auth(account)
storage_put(key, value)
unchanged_line"
	hints := DiagnosticsForSource(source)
	assert.Len(t, hints, 2)
	assert.Equal(t, 0, hints[0].Line)
	assert.Contains(t, hints[0].Message, "require_auth")
	assert.Equal(t, 1, hints[1].Line)
	assert.Contains(t, hints[1].Message, "storage_put")
}
