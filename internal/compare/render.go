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

package compare

import (
	"fmt"
	"strings"

	"github.com/dotandev/hintents/internal/simulator"
	"github.com/dotandev/hintents/internal/visualizer"
)

const (
	colWidth  = 52 // width of each column in the side-by-side table
	columnSep = " | "
)

// Render prints a human-readable side-by-side diff of a DiffResult to stdout.
// It uses the visualizer package for theme-aware colours.
func Render(result *DiffResult) {
	if result == nil {
		return
	}

	printHeader()

	// ── Status ────────────────────────────────────────────────────────────────
	fmt.Println(sectionTitle("Execution Status"))
	renderStatus(result.StatusDiff)

	// ── Budget / Resource Usage ───────────────────────────────────────────────
	if result.BudgetDiff != nil {
		fmt.Println()
		fmt.Println(sectionTitle("Resource Usage (Local vs On-Chain)"))
		renderBudget(result.BudgetDiff)
	}

	// ── Raw Event Diff ────────────────────────────────────────────────────────
	if len(result.EventDiffs) > 0 {
		fmt.Println()
		fmt.Println(sectionTitle("Event Log Diff"))
		renderEventDiffs(result.EventDiffs)
	}

	// ── Diagnostic Event Diff ─────────────────────────────────────────────────
	if len(result.DiagnosticDiffs) > 0 {
		fmt.Println()
		fmt.Println(sectionTitle("Diagnostic Event Diff"))
		renderDiagnosticDiffs(result.DiagnosticDiffs)
	}

	// ── Divergent Call Paths ──────────────────────────────────────────────────
	if len(result.CallPathDivergences) > 0 {
		fmt.Println()
		fmt.Println(sectionTitle("Divergent Call Paths"))
		renderCallPaths(result.CallPathDivergences)
	}

	// ── Summary ───────────────────────────────────────────────────────────────
	fmt.Println()
	renderSummary(result)
}

// ─── internal renderers ───────────────────────────────────────────────────────

func printHeader() {
	sep := strings.Repeat("─", colWidth*2+len(columnSep))
	fmt.Println()
	fmt.Println(visualizer.Colorize("╔"+strings.Repeat("═", len(sep))+"╗", "cyan"))
	title := "  COMPARE REPLAY  ─  Local WASM  vs  On-Chain WASM  "
	pad := len(sep) - len(title)
	if pad < 0 {
		pad = 0
	}
	fmt.Printf(visualizer.Colorize("║", "cyan")+"%s"+strings.Repeat(" ", pad)+visualizer.Colorize("║", "cyan")+"
", title)
	fmt.Println(visualizer.Colorize("╚"+strings.Repeat("═", len(sep))+"╝", "cyan"))
	fmt.Println()
}

func sectionTitle(title string) string {
	line := "── " + title + " " + strings.Repeat("─", max(0, 60-len(title)))
	return visualizer.Colorize(line, "bold")
}

func renderStatus(sd StatusDiff) {
	leftLabel := "LOCAL"
	rightLabel := "ON-CHAIN"
	fmt.Printf("  %-*s%s%-*s
", colWidth, leftLabel, columnSep, colWidth, rightLabel)
	fmt.Printf("  %s
", strings.Repeat("-", colWidth*2+len(columnSep)))

	localStatus := statusLine(sd.LocalStatus, sd.LocalError)
	onChainStatus := statusLine(sd.OnChainStatus, sd.OnChainError)

	if sd.Match {
		fmt.Printf("  %-*s%s%-*s  %s
",
			colWidth, localStatus, columnSep, colWidth, onChainStatus,
			visualizer.Colorize("[MATCH]", "green"))
	} else {
		fmt.Printf("  %-*s%s%-*s  %s
",
			colWidth, localStatus, columnSep, colWidth, onChainStatus,
			visualizer.Colorize("[DIFF]", "red"))
	}
}

func statusLine(status, errMsg string) string {
	s := status
	if errMsg != "" {
		s += " – " + truncate(errMsg, 30)
	}
	return s
}

func renderBudget(bd *BudgetDiff) {
	fmt.Printf("  %-22s  %-15s  %-15s  %s
", "Metric", "Local", "On-Chain", "Delta")
	fmt.Printf("  %s
", strings.Repeat("-", 70))

	cpuDeltaStr := formatDelta(bd.CPUDelta)
	memDeltaStr := formatDelta(bd.MemoryDelta)
	opsDeltaStr := formatDeltaInt(bd.OpsDelta)

	fmt.Printf("  %-22s  %-15d  %-15d  %s
",
		"CPU Instructions", bd.LocalCPU, bd.OnChainCPU, colorizeDelta(cpuDeltaStr, bd.CPUDelta))
	fmt.Printf("  %-22s  %-15d  %-15d  %s
",
		"Memory Bytes", bd.LocalMem, bd.OnChainMem, colorizeDelta(memDeltaStr, bd.MemoryDelta))
	fmt.Printf("  %-22s  %-15d  %-15d  %s
",
		"Operations", bd.LocalOps, bd.OnChainOps, colorizeDelta(opsDeltaStr, int64(bd.OpsDelta)))
}

func renderEventDiffs(diffs []EventDiff) {
	fmt.Printf("  %-6s  %-*s%s%-*s
", "#", colWidth, "LOCAL", columnSep, colWidth, "ON-CHAIN")
	fmt.Printf("  %s
", strings.Repeat("-", colWidth*2+len(columnSep)+8))

	for _, d := range diffs {
		localEvt := truncate(d.LocalEvent, colWidth)
		onChainEvt := truncate(d.OnChainEvent, colWidth)
		var marker string
		if d.Divergent {
			marker = visualizer.Colorize("[!]", "yellow") + " "
		} else {
			marker = visualizer.Colorize("[=]", "dim") + " "
		}
		fmt.Printf("%s[%3d]  %-*s%s%-*s
",
			marker, d.Index+1, colWidth, localEvt, columnSep, colWidth, onChainEvt)
	}
}

func renderDiagnosticDiffs(diffs []DiagnosticDiff) {
	fmt.Printf("  %-6s  %-*s%s%-*s
", "#", colWidth, "LOCAL", columnSep, colWidth, "ON-CHAIN")
	fmt.Printf("  %s
", strings.Repeat("-", colWidth*2+len(columnSep)+8))

	for _, d := range diffs {
		localDesc := diagnosticSummary(d.Local)
		onChainDesc := diagnosticSummary(d.OnChain)

		var marker string
		switch {
		case d.DivergentPath:
			marker = visualizer.Colorize("[PATH]", "red") + " "
		case d.Divergent:
			marker = visualizer.Colorize("[DIFF]", "yellow") + " "
		default:
			marker = visualizer.Colorize("[=]   ", "dim") + " "
		}

		fmt.Printf("%s[%3d]  %-*s%s%-*s
",
			marker, d.Index+1, colWidth, truncate(localDesc, colWidth),
			columnSep, colWidth, truncate(onChainDesc, colWidth))

		// Show topic diff inline if both sides have the event but topics differ
		if d.Local != nil && d.OnChain != nil && d.Divergent && !d.DivergentPath {
			renderTopicDiff(d.Local.Topics, d.OnChain.Topics)
		}
	}
}

func renderTopicDiff(local, onChain []string) {
	maxLen := len(local)
	if len(onChain) > maxLen {
		maxLen = len(onChain)
	}
	for i := 0; i < maxLen; i++ {
		var lt, ot string
		if i < len(local) {
			lt = local[i]
		}
		if i < len(onChain) {
			ot = onChain[i]
		}
		if lt != ot {
			fmt.Printf("        %s topic[%d]: %q  →  %q
",
				visualizer.Colorize("↳", "yellow"), i, lt, ot)
		}
	}
}

func renderCallPaths(divs []CallPathDivergence) {
	for i, div := range divs {
		fmt.Printf("  %s  Divergence #%d at event [%d]
",
			visualizer.Colorize("[PATH]", "red"), i+1, div.EventIndex+1)
		fmt.Printf("       Reason    : %s
", div.Reason)
		fmt.Printf("       Local     : %s
", visualizer.Colorize(div.LocalSummary, "cyan"))
		fmt.Printf("       On-Chain  : %s
", visualizer.Colorize(div.OnChainSummary, "magenta"))
		fmt.Println()
	}
}

func renderSummary(result *DiffResult) {
	fmt.Println(sectionTitle("Summary"))
	fmt.Println()

	if !result.HasDivergence {
		fmt.Printf("  %s  Local and on-chain execution are IDENTICAL
", visualizer.Success())
	} else {
		fmt.Printf("  %s  Divergence detected between local and on-chain execution
", visualizer.Warning())
	}

	fmt.Println()
	fmt.Printf("  %-30s  %d
", "Total events compared:", result.TotalEvents)
	fmt.Printf("  %-30s  %s
", "Identical events:",
		visualizer.Colorize(fmt.Sprintf("%d", result.IdenticalEvents), "green"))
	fmt.Printf("  %-30s  %s
", "Divergent events:",
		colorizeDivergentCount(result.DivergentEvents))
	fmt.Printf("  %-30s  %s
", "Call-path divergences:",
		colorizeDivergentCount(len(result.CallPathDivergences)))

	if result.BudgetDiff != nil {
		fmt.Println()
		cpuPct := budgetDeltaPct(result.BudgetDiff.CPUDelta, result.BudgetDiff.OnChainCPU)
		memPct := budgetDeltaPct(result.BudgetDiff.MemoryDelta, result.BudgetDiff.OnChainMem)
		fmt.Printf("  %-30s  %s
", "CPU delta vs on-chain:", colorizePct(cpuPct))
		fmt.Printf("  %-30s  %s
", "Memory delta vs on-chain:", colorizePct(memPct))
	}

	fmt.Println()
	sep := strings.Repeat("─", colWidth*2+len(columnSep))
	fmt.Println(visualizer.Colorize(sep, "dim"))
}

// ─── formatting helpers ───────────────────────────────────────────────────────

func diagnosticSummary(e *simulator.DiagnosticEvent) string {
	if e == nil {
		return "<absent>"
	}
	cid := ""
	if e.ContractID != nil {
		cid = truncate(*e.ContractID, 12)
	}
	topics := ""
	if len(e.Topics) > 0 {
		topics = truncate(e.Topics[0], 16)
	}
	return fmt.Sprintf("%s/%s %s", e.EventType, cid, topics)
}

func truncate(s string, n int) string {
	if len(s) <= n {
		return s
	}
	if n <= 3 {
		return s[:n]
	}
	return s[:n-3] + "..."
}

func formatDelta(v int64) string {
	if v > 0 {
		return fmt.Sprintf("+%d", v)
	}
	return fmt.Sprintf("%d", v)
}

func formatDeltaInt(v int) string {
	if v > 0 {
		return fmt.Sprintf("+%d", v)
	}
	return fmt.Sprintf("%d", v)
}

func colorizeDelta(s string, v int64) string {
	switch {
	case v > 0:
		return visualizer.Colorize(s, "yellow")
	case v < 0:
		return visualizer.Colorize(s, "green")
	default:
		return visualizer.Colorize(s, "dim")
	}
}

func colorizeDivergentCount(n int) string {
	if n == 0 {
		return visualizer.Colorize("0", "green")
	}
	return visualizer.Colorize(fmt.Sprintf("%d", n), "red")
}

func budgetDeltaPct(delta int64, base uint64) float64 {
	if base == 0 {
		return 0
	}
	return float64(delta) / float64(base) * 100.0
}

func colorizePct(pct float64) string {
	s := fmt.Sprintf("%.2f%%", pct)
	switch {
	case pct > 10:
		return visualizer.Colorize(s, "red")
	case pct > 0:
		return visualizer.Colorize(s, "yellow")
	case pct < 0:
		return visualizer.Colorize(s, "green")
	default:
		return visualizer.Colorize(s, "dim")
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
