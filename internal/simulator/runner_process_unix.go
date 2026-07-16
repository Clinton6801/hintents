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

//go:build !windows

package simulator

import (
	"errors"
	"os"
	"os/exec"
	"syscall"
	"time"
)

func prepareCommand(cmd *exec.Cmd) {
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
}

func terminateCommand(cmd *exec.Cmd, graceTimeout time.Duration) error {
	if cmd == nil || cmd.Process == nil {
		return nil
	}

	targetPID := cmd.Process.Pid
	if pgid, err := syscall.Getpgid(cmd.Process.Pid); err == nil {
		targetPID = -pgid
	}

	if err := syscall.Kill(targetPID, syscall.SIGTERM); err != nil && !errors.Is(err, syscall.ESRCH) {
		return err
	}

	deadline := time.Now().Add(graceTimeout)
	for time.Now().Before(deadline) {
		if processExited(cmd.Process) {
			return nil
		}
		time.Sleep(25 * time.Millisecond)
	}

	if err := syscall.Kill(targetPID, syscall.SIGKILL); err != nil && !errors.Is(err, syscall.ESRCH) {
		return err
	}
	return nil
}

func processExited(process *os.Process) bool {
	if process == nil {
		return true
	}
	err := process.Signal(syscall.Signal(0))
	return errors.Is(err, syscall.ESRCH)
}
