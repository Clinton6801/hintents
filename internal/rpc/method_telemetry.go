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

package rpc

import "context"

// MethodTelemetry is an optional SDK hook for timing method execution.
// Implementations can forward timings to metrics/telemetry backends.
type MethodTelemetry interface {
	StartMethodTimer(ctx context.Context, method string, attributes map[string]string) MethodTimer
}

// MethodTimer represents a started method execution timer.
type MethodTimer interface {
	Stop(err error)
}

var (
	_ MethodTelemetry = (*noopMethodTelemetry)(nil)
	_ MethodTimer     = (*noopMethodTimer)(nil)
)

type noopMethodTelemetry struct{}

func (noopMethodTelemetry) StartMethodTimer(_ context.Context, _ string, _ map[string]string) MethodTimer {
	return noopMethodTimer{}
}

type noopMethodTimer struct{}

func (noopMethodTimer) Stop(_ error) {}

func defaultMethodTelemetry() MethodTelemetry {
	return noopMethodTelemetry{}
}
