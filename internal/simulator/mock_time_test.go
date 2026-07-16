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

import "testing"

func TestMockTimeInjection_OverridesRequestTimestamp(t *testing.T) {
	const fixedTime int64 = 1700000000

	runner := &Runner{
		BinaryPath: "unused-in-this-test",
		Debug:      false,
		MockTime:   fixedTime,
	}

	req := &SimulationRequest{
		EnvelopeXdr:   "test-envelope",
		ResultMetaXdr: "test-meta",
		Timestamp:     9999999,
	}

	if runner.MockTime != 0 {
		req.Timestamp = runner.MockTime
	}

	if req.Timestamp != fixedTime {
		t.Errorf("expected Timestamp %d after mock-time injection, got %d", fixedTime, req.Timestamp)
	}
}

func TestMockTimeInjection_ZeroDoesNotOverride(t *testing.T) {
	const originalTime int64 = 1234567890

	runner := &Runner{
		BinaryPath: "unused-in-this-test",
		Debug:      false,
		MockTime:   0,
	}

	req := &SimulationRequest{
		EnvelopeXdr:   "test-envelope",
		ResultMetaXdr: "test-meta",
		Timestamp:     originalTime,
	}

	if runner.MockTime != 0 {
		req.Timestamp = runner.MockTime
	}

	if req.Timestamp != originalTime {
		t.Errorf("expected Timestamp %d unchanged when MockTime is 0, got %d", originalTime, req.Timestamp)
	}
}

func TestNewRunnerWithMockTime_SetsField(t *testing.T) {
	const wantMockTime int64 = 1700000000

	r := &Runner{
		BinaryPath: "/fake/path",
		Debug:      false,
		MockTime:   wantMockTime,
	}

	if r.MockTime != wantMockTime {
		t.Errorf("expected MockTime %d, got %d", wantMockTime, r.MockTime)
	}
}
