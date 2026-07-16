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

package fuzz

import (
	"fmt"
	"time"
)

// FuzzingStats represents statistics from a fuzzing run
type FuzzingStats struct {
	StartTime          time.Time
	EndTime            time.Time
	ExecutionCount     uint64
	CrashCount         uint64
	NewCoverageCount   uint64
	CorpusSize         int
	CoverageEntryCount int
	UniqueInputsCount  int
}

// Duration returns the total execution time
func (s *FuzzingStats) Duration() time.Duration {
	if s.EndTime.IsZero() {
		return 0
	}
	return s.EndTime.Sub(s.StartTime)
}

// ExecutionsPerSecond calculates throughput
func (s *FuzzingStats) ExecutionsPerSecond() float64 {
	d := s.Duration()
	if d == 0 {
		return 0
	}
	return float64(s.ExecutionCount) / d.Seconds()
}

// CoverageStatistics represents aggregate coverage statistics
type CoverageStatistics struct {
	CorpusSize                int
	UniqueCoverageCount       int
	CrashCount                int
	ExecutionCount            uint64
	MaxCoverage               uint32
	AvgCoverage               uint32
	TimeSinceLastCoverageGrow time.Duration
}

// String returns a human-readable representation of coverage statistics
func (cs CoverageStatistics) String() string {
	return fmt.Sprintf(
		"CoverageStats{corpus_size=%d, unique_coverage=%d, crashes=%d, executions=%d, max_coverage=%d, avg_coverage=%d, time_since_grow=%v}",
		cs.CorpusSize, cs.UniqueCoverageCount, cs.CrashCount, cs.ExecutionCount,
		cs.MaxCoverage, cs.AvgCoverage, cs.TimeSinceLastCoverageGrow,
	)
}

// String returns a human-readable representation of fuzzing statistics
func (s *FuzzingStats) String() string {
	return fmt.Sprintf(
		"FuzzingStats{duration=%v, executions=%d, crashes=%d, new_coverage=%d, corpus_size=%d, coverage_entries=%d, exec/sec=%.2f}",
		s.Duration(), s.ExecutionCount, s.CrashCount, s.NewCoverageCount,
		s.CorpusSize, s.CoverageEntryCount, s.ExecutionsPerSecond(),
	)
}
