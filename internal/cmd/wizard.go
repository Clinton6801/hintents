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

	"github.com/dotandev/hintents/internal/errors"
	"github.com/dotandev/hintents/internal/rpc"
	"github.com/dotandev/hintents/internal/wizard"
	"github.com/spf13/cobra"
)

var wizardCmd = &cobra.Command{
	Use:     "wizard",
	GroupID: "development",
	Short:   "Interactive transaction selection wizard",
	Long:    "Find and select recent failed transactions for debugging.",
	RunE: func(cmd *cobra.Command, args []string) error {
		account, _ := cmd.Flags().GetString("account")
		network, _ := cmd.Flags().GetString("network")

		if account == "" {
			return errors.WrapCliArgumentRequired("account")
		}

		client, err := rpc.NewClient(rpc.WithNetwork(rpc.Network(network)))
		if err != nil {
			return errors.WrapValidationError(err.Error())
		}

		w := wizard.New(client)
		result, err := w.SelectTransaction(cmd.Context(), account)
		if err != nil {
			return err
		}

		fmt.Printf("
Selected: %s
Status: %s
Created: %s

Run: erst debug %s
",
			result.Hash, result.Status, result.CreatedAt, result.Hash)
		return nil
	},
}

func init() {
	wizardCmd.Flags().StringP("account", "a", "", "Stellar account address")
	wizardCmd.Flags().StringP("network", "n", string(rpc.Mainnet), "Network (testnet, mainnet, futurenet)")

	_ = wizardCmd.RegisterFlagCompletionFunc("network", completeNetworkFlag)

	rootCmd.AddCommand(wizardCmd)
}
