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

package telemetry

import (
	"context"
	"testing"
)

func TestInit(t *testing.T) {
	ctx := context.Background()

	// Test with tracing disabled
	cleanup, err := Init(ctx, Config{
		Enabled: false,
	})
	if err != nil {
		t.Fatalf("Failed to initialize telemetry with disabled config: %v", err)
	}
	cleanup()

	// Graceful degradation: Init must never fail even when collector is unreachable
	cleanup, err = Init(ctx, Config{
		Enabled:     true,
		ExporterURL: "http://localhost:4318",
		ServiceName: "test-service",
	})
	if err != nil {
		t.Fatalf("Init must not fail when collector is down (graceful degradation): %v", err)
	}
	cleanup()

	// Tracer is always available (no-op if collector was unreachable)
	tracer := GetTracer()
	if tracer == nil {
		t.Fatal("Tracer should not be nil after initialization")
	}
	_, span := tracer.Start(ctx, "test-span")
	span.End()
}

func TestGetTracer(t *testing.T) {
	// Should not panic even if not initialized
	tracer := GetTracer()
	if tracer == nil {
		t.Fatal("GetTracer should never return nil")
	}

	// Should be able to create spans (no-op if not initialized)
	ctx := context.Background()
	_, span := tracer.Start(ctx, "test-span")
	span.End()
}

// TestInit_UnreachableCollector proves graceful degradation: with tracing enabled
// and an unreachable OTLP endpoint, Init succeeds and core paths (GetTracer, spans)
// work without blocking or error. Run with: go test ./internal/telemetry/... -v -run TestInit_UnreachableCollector
func TestInit_UnreachableCollector(t *testing.T) {
	ctx := context.Background()
	// Use a port that nothing listens on so the collector is "down"
	cleanup, err := Init(ctx, Config{
		Enabled:     true,
		ExporterURL: "http://127.0.0.1:37999",
		ServiceName: "test-service",
	})
	if err != nil {
		t.Fatalf("graceful degradation: Init must not fail when collector is down, got: %v", err)
	}
	defer cleanup()

	tracer := GetTracer()
	if tracer == nil {
		t.Fatal("GetTracer must never return nil")
	}
	_, span := tracer.Start(ctx, "telemetry-test-span")
	span.End()
	// If we get here without blocking or panic, telemetry fails silently as intended
}
