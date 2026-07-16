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
	"fmt"
	"sort"
	"strings"
)

// TraceStep models one callgraph row for markdown export.
type TraceStep struct {
	Step     int
	Contract string
	Function string
	Caller   string
}

// Trace is a deterministic markdown-exportable execution trace.
type Trace struct {
	Steps []TraceStep
}

// ExportMarkdown renders a trace as a GitHub-friendly markdown table.
func ExportMarkdown(trace Trace) string {
	if len(trace.Steps) == 0 {
		return "| Step | Contract | Function | Caller |
|------|----------|----------|--------|
"
	}

	steps := make([]TraceStep, len(trace.Steps))
	copy(steps, trace.Steps)
	sort.SliceStable(steps, func(i, j int) bool {
		if steps[i].Step != steps[j].Step {
			return steps[i].Step < steps[j].Step
		}
		if steps[i].Contract != steps[j].Contract {
			return steps[i].Contract < steps[j].Contract
		}
		if steps[i].Function != steps[j].Function {
			return steps[i].Function < steps[j].Function
		}
		return steps[i].Caller < steps[j].Caller
	})

	var b strings.Builder
	b.WriteString("| Step | Contract | Function | Caller |
")
	b.WriteString("|------|----------|----------|--------|
")
	for _, step := range steps {
		fmt.Fprintf(
			&b,
			"| %d | %s | %s | %s |
",
			step.Step,
			escapeMarkdownCell(step.Contract),
			escapeMarkdownCell(step.Function),
			escapeMarkdownCell(step.Caller),
		)
	}

	return b.String()
}

func escapeMarkdownCell(v string) string {
	escaped := strings.ReplaceAll(v, "|", "\|")
	return strings.TrimSpace(escaped)
}
