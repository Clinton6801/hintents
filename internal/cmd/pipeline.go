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
	"io"
	"os"

	"github.com/dotandev/hintents/internal/errors"
	"github.com/dotandev/hintents/internal/pipeline"
	"github.com/spf13/cobra"
)

var pipelineCmd = &cobra.Command{
	Use:   "pipeline",
	Short: "Manage and run Programmable Transaction Pipelines (PTB-style)",
}

var pipelineRunCmd = &cobra.Command{
	Use:   "run [file.json]",
	Short: "Run a pipeline from a JSON file (or stdin)",
	RunE: func(cmd *cobra.Command, args []string) error {
		var data []byte
		var err error

		if len(args) > 0 {
			data, err = os.ReadFile(args[0])
			if err != nil {
				return errors.WrapValidationError(fmt.Sprintf("failed to read file: %v", err))
			}
		} else {
			stat, _ := os.Stdin.Stat()
			if (stat.Mode() & os.ModeCharDevice) == 0 {
				data, err = io.ReadAll(os.Stdin)
				if err != nil {
					return errors.WrapValidationError("failed to read from stdin")
				}
			} else {
				return errors.WrapValidationError("must provide file or pipe JSON to stdin")
			}
		}

		p, err := pipeline.FromJSON(data)
		if err != nil {
			return err
		}

		fmt.Printf("Pipeline loaded with %d commands.
", len(p.Commands))
		fmt.Println("Executing Pipeline (Simulation)...")
		for i, c := range p.Commands {
			fmt.Printf("[%d] %s -> %s (Args: %v)
", i, c.Type, c.Target, c.Args)
		}
		fmt.Println("Pipeline execution finished.")
		return nil
	},
}

func init() {
	pipelineCmd.AddCommand(pipelineRunCmd)
	rootCmd.AddCommand(pipelineCmd)
}
