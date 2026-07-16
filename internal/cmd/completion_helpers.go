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
	"github.com/spf13/cobra"
)

var networkAliases = []string{"testnet	Stellar test network", "mainnet	Stellar public network", "futurenet	Stellar future network"}
var initNetworkAliases = []string{"public	Stellar public network", "testnet	Stellar test network", "futurenet	Stellar future network", "standalone	Local standalone network"}
var themeNames = []string{"default	Standard terminal colors", "deuteranopia	Red-green color blind friendly", "protanopia	Red color blind friendly", "tritanopia	Blue-yellow color blind friendly", "high-contrast	High contrast for low-vision"}
var xdrFormats = []string{"json	JSON output", "table	Tabular output"}
var xdrTypes = []string{"ledger-entry	Ledger entry XDR", "diagnostic-event	Diagnostic event XDR"}
var reportFormats = []string{"html	HTML report", "pdf	PDF report", "json	JSON report", "html,pdf	Both HTML and PDF"}

func completeNetworkFlag(_ *cobra.Command, _ []string, _ string) ([]string, cobra.ShellCompDirective) {
	return networkAliases, cobra.ShellCompDirectiveNoFileComp
}

func completeInitNetworkFlag(_ *cobra.Command, _ []string, _ string) ([]string, cobra.ShellCompDirective) {
	return initNetworkAliases, cobra.ShellCompDirectiveNoFileComp
}

func completeThemeFlag(_ *cobra.Command, _ []string, _ string) ([]string, cobra.ShellCompDirective) {
	return themeNames, cobra.ShellCompDirectiveNoFileComp
}

func completeXDRFormatFlag(_ *cobra.Command, _ []string, _ string) ([]string, cobra.ShellCompDirective) {
	return xdrFormats, cobra.ShellCompDirectiveNoFileComp
}

func completeXDRTypeFlag(_ *cobra.Command, _ []string, _ string) ([]string, cobra.ShellCompDirective) {
	return xdrTypes, cobra.ShellCompDirectiveNoFileComp
}

func completeReportFormatFlag(_ *cobra.Command, _ []string, _ string) ([]string, cobra.ShellCompDirective) {
	return reportFormats, cobra.ShellCompDirectiveNoFileComp
}

func completeNoOp(_ *cobra.Command, _ []string, _ string) ([]string, cobra.ShellCompDirective) {
	return nil, cobra.ShellCompDirectiveNoFileComp
}
