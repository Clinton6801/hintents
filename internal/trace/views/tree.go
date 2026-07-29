// Copyright 2026 Erst Users
// SPDX-License-Identifier: Apache-2.0

package views

import "github.com/dotandev/hintents/internal/trace"

// BuildTraceTreeView creates a collapsible tree from a flat execution trace.
func BuildTraceTreeView(trace *trace.ExecutionTrace) *trace.TraceNode {
	return trace.BuildTraceTree()
}
