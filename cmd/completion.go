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
	"os"

	"github.com/spf13/cobra"
)

var completionCmd = &cobra.Command{
	Use:   "completion [bash|zsh|fish|powershell]",
	Short: "Generate completion script for your shell",
	Long: `Generate shell completion scripts for erst commands.

The completion script must be evaluated to provide interactive completion of erst commands.
This can be done by sourcing it from your shell profile or piping it to the appropriate location.

Installation instructions:

  Bash:
    $ erst completion bash > /etc/bash_completion.d/erst
    $ source ~/.bashrc

  Zsh:
    $ erst completion zsh > "${fpath[1]}/_erst"
    # or place in your custom completions directory:
    $ mkdir -p ~/.zsh/completions
    $ erst completion zsh > ~/.zsh/completions/_erst
    # then add to your ~/.zshrc: fpath=(~/.zsh/completions $fpath)

  Fish:
    $ erst completion fish > ~/.config/fish/completions/erst.fish
    $ source ~/.config/fish/config.fish

  PowerShell:
    PS> erst completion powershell | Out-String | Invoke-Expression
    # To load completions for every new session, add to your PowerShell profile:
    PS> erst completion powershell >> $PROFILE

For detailed instructions on setting up completions for your shell, consult your shell's documentation.`,
	DisableFlagsInUseLine: true,
	ValidArgs:             []string{"bash", "zsh", "fish", "powershell"},
	Args:                  cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	RunE: func(_ *cobra.Command, args []string) error {
		shell := args[0]

		switch shell {
		case "bash":
			return rootCmd.GenBashCompletionV2(os.Stdout, true)
		case "zsh":
			return rootCmd.GenZshCompletion(os.Stdout)
		case "fish":
			return rootCmd.GenFishCompletion(os.Stdout, true)
		case "powershell":
			return rootCmd.GenPowerShellCompletionWithDesc(os.Stdout)
		default:
			return fmt.Errorf("unsupported shell type %q. Valid shells: bash, zsh, fish, powershell", shell)
		}
	},
}

func init() {
	rootCmd.AddCommand(completionCmd)
}
