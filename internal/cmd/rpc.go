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
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/dotandev/hintents/internal/config"
	"github.com/spf13/cobra"
)

var (
	rpcHealthURLFlag string
)

var rpcCmd = &cobra.Command{
	Use:     "rpc",
	GroupID: "utility",
	Short:   "Manage and monitor RPC endpoints",
}

var rpcHealthCmd = &cobra.Command{
	Use:     "health",
	Aliases: []string{"rpc:health"},
	Short:   "Check the health of configured RPC endpoints",
	RunE: func(cmd *cobra.Command, args []string) error {
		urls := []string{}
		cfg, cfgErr := config.Load()
		timeout := 15 * time.Second
		if rpcHealthURLFlag != "" {
			urls = strings.Split(rpcHealthURLFlag, ",")
		} else {
			if cfgErr == nil {
				if len(cfg.RpcUrls) > 0 {
					urls = cfg.RpcUrls
				} else if cfg.RpcUrl != "" {
					urls = []string{cfg.RpcUrl}
				}
				if cfg.RequestTimeout > 0 {
					timeout = time.Duration(cfg.RequestTimeout) * time.Second
				}
			}
		}

		if len(urls) == 0 {
			return fmt.Errorf("no RPC URLs configured and none provided via --rpc")
		}

		fmt.Println("[STATS] RPC Endpoint Status:")
		fmt.Println()

		client := &http.Client{Timeout: timeout}

		for i, url := range urls {
			url = strings.TrimSpace(url)
			if url == "" {
				continue
			}

			start := time.Now()
			status := "[OK]"
			success := true
			errStr := ""

			resp, err := client.Get(url)
			if err != nil {
				status = "[FAIL]"
				success = false
				errStr = err.Error()
			} else {
				defer resp.Body.Close()
				if resp.StatusCode >= 400 {
					status = "[FAIL]"
					success = false
					errStr = fmt.Sprintf("HTTP %d", resp.StatusCode)
				}
			}

			duration := time.Since(start)

			if success {
				fmt.Printf("  [%d]  %s
", i+1, url)
				fmt.Printf("      Status: %s
", status)
				fmt.Printf("      Latency: %v
", duration.Round(time.Millisecond))
			} else {
				fmt.Printf("  [%d] %s %s
", i+1, status, url)
				fmt.Printf("      Error: %s
", errStr)
			}
			fmt.Println()
		}

		return nil
	},
}

func init() {
	rpcHealthCmd.Flags().StringVar(&rpcHealthURLFlag, "rpc", "", "RPC URLs to check (comma-separated)")
	rpcCmd.AddCommand(rpcHealthCmd)

	// Add the rpc:health as a top-level command for compatibility
	rpcHealthAliasCmd := *rpcHealthCmd
	rpcHealthAliasCmd.Use = "rpc:health"
	rpcHealthAliasCmd.Hidden = true
	rootCmd.AddCommand(&rpcHealthAliasCmd)

	rootCmd.AddCommand(rpcCmd)
}
