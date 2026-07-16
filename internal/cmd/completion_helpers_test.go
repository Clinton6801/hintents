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
	"testing"

	"github.com/spf13/cobra"
)

func TestCompleteNetworkFlag(t *testing.T) {
	completions, directive := completeNetworkFlag(nil, nil, "")
	if directive != cobra.ShellCompDirectiveNoFileComp {
		t.Fatalf("expected ShellCompDirectiveNoFileComp, got %v", directive)
	}
	if len(completions) != 3 {
		t.Fatalf("expected 3 network completions, got %d", len(completions))
	}
}

func TestCompleteInitNetworkFlag(t *testing.T) {
	completions, directive := completeInitNetworkFlag(nil, nil, "")
	if directive != cobra.ShellCompDirectiveNoFileComp {
		t.Fatalf("expected ShellCompDirectiveNoFileComp, got %v", directive)
	}
	if len(completions) != 4 {
		t.Fatalf("expected 4 init network completions, got %d", len(completions))
	}
}

func TestCompleteThemeFlag(t *testing.T) {
	completions, directive := completeThemeFlag(nil, nil, "")
	if directive != cobra.ShellCompDirectiveNoFileComp {
		t.Fatalf("expected ShellCompDirectiveNoFileComp, got %v", directive)
	}
	if len(completions) != 5 {
		t.Fatalf("expected 5 theme completions, got %d", len(completions))
	}
}

func TestCompleteXDRFormatFlag(t *testing.T) {
	completions, directive := completeXDRFormatFlag(nil, nil, "")
	if directive != cobra.ShellCompDirectiveNoFileComp {
		t.Fatalf("expected ShellCompDirectiveNoFileComp, got %v", directive)
	}
	if len(completions) != 2 {
		t.Fatalf("expected 2 xdr format completions, got %d", len(completions))
	}
}

func TestCompleteXDRTypeFlag(t *testing.T) {
	completions, directive := completeXDRTypeFlag(nil, nil, "")
	if directive != cobra.ShellCompDirectiveNoFileComp {
		t.Fatalf("expected ShellCompDirectiveNoFileComp, got %v", directive)
	}
	if len(completions) != 2 {
		t.Fatalf("expected 2 xdr type completions, got %d", len(completions))
	}
}

func TestCompleteReportFormatFlag(t *testing.T) {
	completions, directive := completeReportFormatFlag(nil, nil, "")
	if directive != cobra.ShellCompDirectiveNoFileComp {
		t.Fatalf("expected ShellCompDirectiveNoFileComp, got %v", directive)
	}
	if len(completions) != 4 {
		t.Fatalf("expected 4 report format completions, got %d", len(completions))
	}
}

func TestCompleteNoOp(t *testing.T) {
	completions, directive := completeNoOp(nil, nil, "")
	if directive != cobra.ShellCompDirectiveNoFileComp {
		t.Fatalf("expected ShellCompDirectiveNoFileComp, got %v", directive)
	}
	if completions != nil {
		t.Fatalf("expected nil completions, got %v", completions)
	}
}
