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

package cmd

import (
	"context"
	"fmt"

	"github.com/dotandev/hintents/internal/rpc"
	"github.com/dotandev/hintents/internal/simulator"
	"github.com/spf13/cobra"
)

var gasAdvisorNetwork string

var gasAdvisorCmd = &cobra.Command{
	Use:     "gas-advisor",
	GroupID: "utility",
	Short:   "Gas optimization advisor with dynamic network baselines",
	Long: `Analyse transaction gas usage against real-world network averages.

Examples:
  erst gas-advisor sync --network testnet
  erst gas-advisor sync --network mainnet`,
}

var gasAdvisorSyncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Sync baselines with current network state",
	RunE: func(cmd *cobra.Command, args []string) error {
		net := rpc.Network(gasAdvisorNetwork)
		client := rpc.NewClientDefault(net, "")
		if client == nil {
			return fmt.Errorf("failed to create RPC client for network: %s", gasAdvisorNetwork)
		}

		runner, err := simulator.NewRunner("", false)
		if err != nil {
			return fmt.Errorf("failed to create simulator runner: %w", err)
		}

		provider := simulator.NewNetworkBaselineProvider(runner)
		advisor := simulator.NewGasOptimizationAdvisor(provider)

		fmt.Printf("Syncing baselines from %s network...
", gasAdvisorNetwork)
		if err := advisor.SyncBaselines(context.Background()); err != nil {
			return fmt.Errorf("sync failed: %w", err)
		}

		bl := advisor.CurrentBaselines()
		fmt.Printf("Baselines synced successfully!
")
		fmt.Printf("  Source:          %s
", bl.Source)
		fmt.Printf("  Avg CPU/op:      %d instructions
", bl.AvgCPUPerOp)
		fmt.Printf("  Avg Memory/op:   %d bytes
", bl.AvgMemoryPerOp)
		fmt.Printf("  Avg Fee:         %d stroops
", bl.AvgFeeStroops)
		fmt.Printf("  Synced at:       %s
", bl.SyncedAt.Format("2006-01-02 15:04:05 UTC"))
		return nil
	},
}

func init() {
	gasAdvisorSyncCmd.Flags().StringVarP(&gasAdvisorNetwork, "network", "n", "testnet", "Network to sync baselines from (testnet|mainnet|futurenet)")
	gasAdvisorCmd.AddCommand(gasAdvisorSyncCmd)
	rootCmd.AddCommand(gasAdvisorCmd)
}
