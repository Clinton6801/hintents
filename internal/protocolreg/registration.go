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

package protocolreg

import (
	"errors"
	"fmt"
	ersterrors "github.com/dotandev/hintents/internal/errors"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

const (
	Scheme             = "erst"
	windowsRegistryKey = `HKEY_CURRENT_USER\Software\Classes\erst`
	linuxDesktopFile   = "erst-protocol.desktop"
	linuxMimeType      = "x-scheme-handler/erst"
	macOSAppName       = "Erst Protocol.app"
	macOSBundleID      = "dev.dotan.erst.protocol"
	macOSExecutable    = "erst-protocol-handler"
)

type Registrar struct {
	executablePath string
	homeDir        string
}

type VerificationReport struct {
	Platform string
	Scheme   string
	Checks   []string
	Issues   []string
}

func NewRegistrar() (*Registrar, error) {
	executablePath, err := os.Executable()
	if err != nil {
		return nil, fmt.Errorf("resolve executable path: %w", err)
	}

	executablePath, err = filepath.Abs(executablePath)
	if err != nil {
		return nil, fmt.Errorf("resolve absolute executable path: %w", err)
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("resolve user home directory: %w", err)
	}

	return &Registrar{
		executablePath: executablePath,
		homeDir:        homeDir,
	}, nil
}

func (r *Registrar) Register() error {
	switch runtime.GOOS {
	case "windows":
		return r.registerWindows()
	case "darwin":
		return r.registerDarwin()
	case "linux":
		return r.registerLinux()
	default:
		return fmt.Errorf("protocol registration is not supported on %s", runtime.GOOS)
	}
}

func (r *Registrar) Unregister() error {
	switch runtime.GOOS {
	case "windows":
		return r.unregisterWindows()
	case "darwin":
		return r.unregisterDarwin()
	case "linux":
		return r.unregisterLinux()
	default:
		return fmt.Errorf("protocol registration is not supported on %s", runtime.GOOS)
	}
}

func (r *Registrar) IsRegistered() bool {
	_, err := r.Verify()
	return err == nil
}

func (r *Registrar) Verify() (*VerificationReport, error) {
	report := &VerificationReport{
		Platform: runtime.GOOS,
		Scheme:   Scheme,
	}

	switch runtime.GOOS {
	case "windows":
		r.verifyWindows(report)
	case "darwin":
		r.verifyDarwin(report)
	case "linux":
		r.verifyLinux(report)
	default:
		report.Issues = append(report.Issues, fmt.Sprintf("protocol verification is not supported on %s", runtime.GOOS))
	}

	if len(report.Issues) > 0 {
		return report, verificationError(report.Issues)
	}

	return report, nil
}

func (r *Registrar) registerWindows() error {
	// Detect Protocol Registry Conflicts (Issue #1198)
	registryOutput, err := runCommand("reg", "query", windowsRegistryKey, "/ve")
	if err == nil && !strings.Contains(registryOutput, "erst") {
		// If the key exists (err == nil) but (Default) doesn't contain 'erst', it's a conflict
		return errors.Join(fmt.Errorf("registry conflict for %s", Scheme), ersterrors.ErrRegistryConflict)
	}

	commands := [][]string{
		{"add", windowsRegistryKey, "/ve", "/d", "URL:ERST Protocol", "/f"},
		{"add", windowsRegistryKey, "/v", "URL Protocol", "/d", "", "/f"},
		{"add", windowsRegistryKey + `\shell\open
