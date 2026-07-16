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
	"strings"
	"testing"
)

func TestDryRunCmd_NetworkValidation_ValidNetworks(t *testing.T) {
	validNetworks := []string{"testnet", "mainnet", "futurenet"}
	for _, network := range validNetworks {
		t.Run(network, func(t *testing.T) {
			prev := dryRunNetworkFlag
			t.Cleanup(func() { dryRunNetworkFlag = prev })
			dryRunNetworkFlag = network

			err := dryRunCmd.PreRunE(dryRunCmd, []string{})
			if err != nil {
				t.Errorf("expected network %q to pass validation, got error: %v", network, err)
			}
		})
	}
}

func TestDryRunCmd_NetworkValidation_InvalidNetwork(t *testing.T) {
	prev := dryRunNetworkFlag
	t.Cleanup(func() { dryRunNetworkFlag = prev })
	dryRunNetworkFlag = "invalidnet"

	err := dryRunCmd.PreRunE(dryRunCmd, []string{})
	if err == nil {
		t.Fatal("expected validation to fail for invalid network, got nil error")
	}
	errMsg := err.Error()
	if !strings.Contains(errMsg, "invalidnet") {
		t.Errorf("expected error to contain the invalid value %q, got: %s", "invalidnet", errMsg)
	}
	for _, valid := range []string{"testnet", "mainnet", "futurenet"} {
		if !strings.Contains(errMsg, valid) {
			t.Errorf("expected error to list supported value %q, got: %s", valid, errMsg)
		}
	}
}

func TestDryRunCmd_NetworkValidation_EmptyNetwork(t *testing.T) {
	prev := dryRunNetworkFlag
	t.Cleanup(func() { dryRunNetworkFlag = prev })
	dryRunNetworkFlag = ""

	err := dryRunCmd.PreRunE(dryRunCmd, []string{})
	if err == nil {
		t.Fatal("expected validation to fail for empty network, got nil error")
	}
}
