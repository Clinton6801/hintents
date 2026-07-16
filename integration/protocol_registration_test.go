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

package integration

import (
	"bytes"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
	"time"
)

func TestProtocolRegistrationLifecycle(t *testing.T) {
	if os.Getenv("ERST_RUN_PROTOCOL_REGISTRATION_TESTS") != "1" {
		t.Skip("set ERST_RUN_PROTOCOL_REGISTRATION_TESTS=1 to enable protocol registration integration tests")
	}

	if runtime.GOOS == "linux" {
		if _, err := exec.LookPath("xdg-mime"); err != nil {
			t.Skip("xdg-mime is required for protocol registration tests")
		}
	}

	_, _, _ = runErst(t, "protocol:unregister")
	t.Cleanup(func() {
		_, _, _ = runErst(t, "protocol:unregister")
	})

	stdout, stderr, err := runErst(t, "protocol:register")
	assertExitCode(t, 0, err)
	assertContains(t, "register output", stdout+stderr, "Registered ERST protocol handler")

	stdout, stderr, err = runErst(t, "protocol:verify")
	assertExitCode(t, 0, err)
	assertContains(t, "verify output", stdout+stderr, "Verified ERST protocol registration")

	scriptPath := os.Getenv("ERST_PROTOCOL_VERIFY_SCRIPT")
	if scriptPath == "" {
		t.Fatal("ERST_PROTOCOL_VERIFY_SCRIPT must be set when protocol registration tests are enabled")
	}
	runNativeVerificationScript(t, scriptPath)

	stdout, stderr, err = runErst(t, "protocol:status")
	assertExitCode(t, 0, err)
	assertContains(t, "status output", stdout+stderr, "REGISTERED")

	stdout, stderr, err = runErst(t, "protocol:unregister")
	assertExitCode(t, 0, err)
	assertContains(t, "unregister output", stdout+stderr, "Unregistered ERST protocol handler")

	stdout, stderr, err = runErst(t, "protocol:verify")
	if exitCode(err) == 0 {
		t.Fatalf("expected protocol:verify to fail after unregister, got stdout=%q stderr=%q", stdout, stderr)
	}
	assertContains(t, "verify failure output", stdout+stderr, "[FAIL]")
}

func runNativeVerificationScript(t *testing.T, scriptPath string) {
	t.Helper()

	absoluteScriptPath := scriptPath
	if !filepath.IsAbs(absoluteScriptPath) {
		absoluteScriptPath = filepath.Join(repoRoot(t), scriptPath)
	}

	ctx, cancel := timeoutCtx(t, 30*time.Second)
	defer cancel()

	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.CommandContext(ctx, "powershell", "-NoProfile", "-ExecutionPolicy", "Bypass", "-File", absoluteScriptPath)
	default:
		cmd = exec.CommandContext(ctx, "bash", absoluteScriptPath)
	}

	cmd.Env = append(os.Environ(), "ERST_BINARY="+binaryPath(t))
	var output bytes.Buffer
	cmd.Stdout = &output
	cmd.Stderr = &output

	if err := cmd.Run(); err != nil {
		t.Fatalf("native verification script failed: %v
%s", err, strings.TrimSpace(output.String()))
	}
}
