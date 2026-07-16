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

package analytics

import (
	"fmt"
)

type StorageGrowthReport struct {
	BeforeBytes int64
	AfterBytes  int64
	DeltaBytes  int64
	PerKeyDelta map[string]int64
}

func PrintStorageReport(report *StorageGrowthReport, fee int64) {
	fmt.Println("[PKG] Contract Storage Growth Report")
	fmt.Println("--------------------------------")
	fmt.Printf("Before: %d bytes
", report.BeforeBytes)
	fmt.Printf("After:  %d bytes
", report.AfterBytes)
	fmt.Printf("Delta:  %+d bytes
", report.DeltaBytes)
	fmt.Printf("Fee Impact: %d stroops

", fee)

	fmt.Println("Per-Key Changes:")
	for key, delta := range report.PerKeyDelta {
		if delta != 0 {
			fmt.Printf("  %s: %+d bytes
", key, delta)
		}
	}
}
