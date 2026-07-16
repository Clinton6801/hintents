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

import (
	"context"
)

var _ RunnerInterface = (*MockRunner)(nil)

type MockRunner struct {
	RunFunc   func(ctx context.Context, req *SimulationRequest) (*SimulationResponse, error)
	CloseFunc func() error
}

func (m *MockRunner) Run(ctx context.Context, req *SimulationRequest) (*SimulationResponse, error) {
	if m.RunFunc != nil {
		return m.RunFunc(ctx, req)
	}
	return &SimulationResponse{Status: "success"}, nil
}

func NewMockRunner(fn func(ctx context.Context, req *SimulationRequest) (*SimulationResponse, error)) *MockRunner {
	return &MockRunner{RunFunc: fn}
}

func NewDefaultMockRunner() *MockRunner {
	return &MockRunner{
		RunFunc: func(ctx context.Context, req *SimulationRequest) (*SimulationResponse, error) {
			return &SimulationResponse{
				Status: "success",
				Events: []string{},
				Logs:   []string{},
			}, nil
		},
		CloseFunc: func() error { return nil },
	}
}

func (m *MockRunner) Close() error {
	if m.CloseFunc != nil {
		return m.CloseFunc()
	}
	return nil
}
