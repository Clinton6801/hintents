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

package analyzer

import (
	"fmt"

	"github.com/dotandev/hintents/internal/simulator"
	"github.com/dotandev/hintents/internal/visualizer"
)

// PrintAssetAnomalies formats and prints Move-Level asset safety anomalies.
func PrintAssetAnomalies(anomalies []simulator.AssetAnomaly) {
	if len(anomalies) == 0 {
		return
	}

	fmt.Println()
	fmt.Println(visualizer.Colorize("=== Move-Level Asset Safety Anomalies Detected ===", "red"))
	fmt.Println(visualizer.Colorize("These mathematical violations were detected during simulation:", "yellow"))

	for i, anomaly := range anomalies {
		fmt.Printf("
%d. [%s] in Contract %s
", i+1, visualizer.Colorize(anomaly.AnomalyType, "red"), anomaly.ContractID)
		fmt.Printf("   Details: %s
", anomaly.Message)
		fmt.Printf("   Amount involved: %d
", anomaly.Amount)
	}

	fmt.Println(visualizer.Colorize("==================================================", "red"))
	fmt.Println()
}
