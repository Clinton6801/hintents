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
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/mattn/go-isatty"
	"github.com/spf13/cobra"
)

var statusFixFlag bool

var statusCmd = &cobra.Command{
	Use:     "status",
	GroupID: "utility",
	Short:   "Check protocol registration and system health",
	Long: `Inspect the health of the erst:// protocol handler registration.

This command verifies that the custom URI scheme (erst://) is properly
registered with the operating system so that deep links work correctly.

On failure it can interactively offer to repair the registration:
  - Windows: rewrites HKCU\Software\Classes\erst registry keys
  - macOS:   rewrites ~/Library/LaunchAgents/com.erst.protocol.plist
  - Linux:   rewrites ~/.local/share/applications/erst-protocol.desktop

Use --fix to skip the interactive prompt and repair automatically.`,
	Example: `  # Check protocol registration status
  erst status

  # Automatically repair without prompting
  erst status --fix`,
	Args: cobra.NoArgs,
	RunE: runStatus,
}

func runStatus(cmd *cobra.Command, args []string) error {
	out := cmd.OutOrStdout()

	fmt.Fprintln(out, "Erst Protocol Registration Status")
	fmt.Fprintln(out, "==================================")
	fmt.Fprintln(out)

	result := checkProtocolRegistration()

	if result.Registered {
		fmt.Fprintf(out, "[32m[OK][0m Protocol handler (erst://) is registered
")
		if result.Detail != "" {
			fmt.Fprintf(out, "  Path: %s
", result.Detail)
		}
		fmt.Fprintln(out)
		fmt.Fprintln(out, "[32mAll checks passed.[0m")
		return nil
	}

	// Registration is broken
	fmt.Fprintf(out, "[31m[FAIL][0m Protocol handler (erst://) is not registered
")
	if result.Detail != "" {
		fmt.Fprintf(out, "  [33m→ %s[0m
", result.Detail)
	}
	fmt.Fprintln(out)

	// Determine whether to attempt repair
	shouldFix := statusFixFlag
	if !shouldFix && isInteractiveTTY(cmd) {
		shouldFix = promptYesNo(cmd, "Would you like ERST to repair the protocol registration? [y/n]: ")
	}

	if !shouldFix {
		fmt.Fprintln(out, "Skipping repair. Run 'erst status --fix' to repair automatically.")
		return nil
	}

	// Attempt repair
	fmt.Fprintln(out, "Repairing protocol registration...")
	if err := repairProtocolRegistration(out); err != nil {
		fmt.Fprintf(out, "[31m[FAIL][0m Repair failed: %v
", err)
		return fmt.Errorf("protocol repair failed: %w", err)
	}

	// Verify the fix
	fmt.Fprintln(out)
	fmt.Fprintln(out, "Verifying repair...")
	verify := checkProtocolRegistration()
	if verify.Registered {
		fmt.Fprintf(out, "[32m[OK][0m Protocol registration repaired successfully
")
		return nil
	}

	fmt.Fprintf(out, "[31m[FAIL][0m Verification failed — registration still broken
")
	if verify.Detail != "" {
		fmt.Fprintf(out, "  [33m→ %s[0m
", verify.Detail)
	}
	return fmt.Errorf("protocol registration repair could not be verified")
}

// ProtocolCheckResult holds the outcome of a protocol registration check.
type ProtocolCheckResult struct {
	Registered bool
	Detail     string
}

// checkProtocolRegistration inspects OS-specific artefacts to determine whether
// the erst:// URI scheme handler is registered.
func checkProtocolRegistration() ProtocolCheckResult {
	switch runtime.GOOS {
	case "windows":
		return checkProtocolWindows()
	case "darwin":
		return checkProtocolDarwin()
	case "linux":
		return checkProtocolLinux()
	default:
		return ProtocolCheckResult{Detail: fmt.Sprintf("unsupported platform: %s", runtime.GOOS)}
	}
}

func checkProtocolWindows() ProtocolCheckResult {
	regPath := `HKEY_CURRENT_USER\Software\Classes\erst`
	out, err := exec.Command("reg", "query", regPath).CombinedOutput()
	if err != nil {
		return ProtocolCheckResult{Detail: "Registry key HKCU\Software\Classes\erst not found"}
	}
	if !strings.Contains(string(out), "URL Protocol") {
		return ProtocolCheckResult{Detail: "Registry key exists but missing URL Protocol value"}
	}
	return ProtocolCheckResult{Registered: true, Detail: regPath}
}

func checkProtocolDarwin() ProtocolCheckResult {
	plistPath := filepath.Join(os.Getenv("HOME"), "Library", "LaunchAgents", "com.erst.protocol.plist")
	info, err := os.Stat(plistPath)
	if err != nil {
		return ProtocolCheckResult{Detail: fmt.Sprintf("Plist not found at %s", plistPath)}
	}
	if info.Size() == 0 {
		return ProtocolCheckResult{Detail: fmt.Sprintf("Plist is empty at %s", plistPath)}
	}

	// Verify the plist contains the expected protocol scheme
	data, err := os.ReadFile(plistPath)
	if err != nil {
		return ProtocolCheckResult{Detail: fmt.Sprintf("Cannot read plist: %v", err)}
	}
	if !strings.Contains(string(data), "<string>erst</string>") {
		return ProtocolCheckResult{Detail: "Plist exists but does not contain erst:// scheme"}
	}

	return ProtocolCheckResult{Registered: true, Detail: plistPath}
}

func checkProtocolLinux() ProtocolCheckResult {
	desktopPath := filepath.Join(os.Getenv("HOME"), ".local", "share", "applications", "erst-protocol.desktop")
	info, err := os.Stat(desktopPath)
	if err != nil {
		return ProtocolCheckResult{Detail: fmt.Sprintf("Desktop file not found at %s", desktopPath)}
	}
	if info.Size() == 0 {
		return ProtocolCheckResult{Detail: fmt.Sprintf("Desktop file is empty at %s", desktopPath)}
	}

	data, err := os.ReadFile(desktopPath)
	if err != nil {
		return ProtocolCheckResult{Detail: fmt.Sprintf("Cannot read desktop file: %v", err)}
	}
	if !strings.Contains(string(data), "x-scheme-handler/erst") {
		return ProtocolCheckResult{Detail: "Desktop file exists but does not contain erst scheme handler"}
	}

	return ProtocolCheckResult{Registered: true, Detail: desktopPath}
}

// repairProtocolRegistration writes the correct protocol handler artefacts for
// the current platform. On macOS/Linux this may require elevated permissions
// for certain paths, so we attempt the write and surface any permission errors.
func repairProtocolRegistration(out io.Writer) error {
	cliPath, err := os.Executable()
	if err != nil {
		return fmt.Errorf("cannot determine executable path: %w", err)
	}

	switch runtime.GOOS {
	case "windows":
		return repairProtocolWindows(cliPath, out)
	case "darwin":
		return repairProtocolDarwin(cliPath, out)
	case "linux":
		return repairProtocolLinux(cliPath, out)
	default:
		return fmt.Errorf("unsupported platform: %s", runtime.GOOS)
	}
}

func repairProtocolWindows(cliPath string, out io.Writer) error {
	regPath := `HKEY_CURRENT_USER\Software\Classes\erst`

	commands := []struct {
		desc string
		args []string
	}{
		{"Setting URL protocol type", []string{"reg", "add", regPath, "/ve", "/d", "URL:ERST Protocol", "/f"}},
		{"Setting URL Protocol value", []string{"reg", "add", regPath, "/v", "URL Protocol", "/d", "", "/f"}},
		{"Setting shell open command", []string{"reg", "add", regPath + `\shell\open
