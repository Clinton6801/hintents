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
	"encoding/base64"
	"fmt"

	"github.com/dotandev/hintents/internal/decoder"
	"github.com/dotandev/hintents/internal/errors"
	"github.com/spf13/cobra"
)

var (
	xdrFormat string
	xdrData   string
	xdrType   string
)

var xdrCmd = &cobra.Command{
	Use:     "xdr",
	GroupID: "utility",
	Short:   "Format and decode XDR data",
	Long:    `Decode and format XDR structures to JSON or table format for easy inspection.`,
	RunE:    xdrExec,
}

func xdrExec(cmd *cobra.Command, args []string) error {
	if xdrData == "" {
		return errors.WrapCliArgumentRequired("data")
	}

	data, err := base64.StdEncoding.DecodeString(xdrData)
	if err != nil {
		return errors.WrapValidationError(fmt.Sprintf("invalid base64 input: %v", err))
	}

	var output interface{}

	switch xdrType {
	case "ledger-entry":
		le, decodeErr := decoder.DecodeXDRBase64AsLedgerEntry(string(data))
		if decodeErr != nil {
			return errors.WrapUnmarshalFailed(decodeErr, "ledger entry")
		}
		output = le

	case "diagnostic-event":
		event, decodeErr := decoder.DecodeXDRBase64AsDiagnosticEvent(string(data))
		if decodeErr != nil {
			return errors.WrapUnmarshalFailed(decodeErr, "diagnostic event")
		}
		output = event

	default:
		return errors.WrapValidationError(fmt.Sprintf("unsupported XDR type: %s (use: ledger-entry, diagnostic-event)", xdrType))
	}

	formatter := decoder.NewXDRFormatter(decoder.FormatType(xdrFormat))
	result, err := formatter.Format(output)
	if err != nil {
		return errors.WrapValidationError(fmt.Sprintf("formatting failed: %v", err))
	}

	fmt.Println(result)
	return nil
}

func init() {
	rootCmd.AddCommand(xdrCmd)

	xdrCmd.Flags().StringVar(&xdrData, "data", "", "Base64-encoded XDR data to decode")
	xdrCmd.Flags().StringVar(&xdrFormat, "format", "json", "Output format: json or table")
	xdrCmd.Flags().StringVar(&xdrType, "type", "ledger-entry", "XDR type: ledger-entry, diagnostic-event")

	_ = xdrCmd.MarkFlagRequired("data")

	_ = xdrCmd.RegisterFlagCompletionFunc("format", completeXDRFormatFlag)
	_ = xdrCmd.RegisterFlagCompletionFunc("type", completeXDRTypeFlag)
}
