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

//go:build !windows && !plan9 && !js && !wasip1

package trace

import (
	"os"
	"os/signal"
	"syscall"
	"unsafe"
)

// ioctlWinsize mirrors the kernel winsize struct used by TIOCGWINSZ.
type ioctlWinsize struct {
	Row, Col       uint16
	Xpixel, Ypixel uint16
}

// getTermWidthSys queries the terminal width via TIOCGWINSZ ioctl.
// Returns 0 if stdout is not a terminal or the call fails.
func getTermWidthSys() int {
	var ws ioctlWinsize
	_, _, errno := syscall.Syscall(
		syscall.SYS_IOCTL,
		os.Stdout.Fd(),
		syscall.TIOCGWINSZ,
		uintptr(unsafe.Pointer(&ws)),
	)
	if errno != 0 || ws.Col == 0 {
		return 0
	}
	return int(ws.Col)
}

// watchResize registers ch to receive os.Signal notifications on SIGWINCH
// (terminal window resize). Call signal.Stop(ch) to deregister.
func watchResize(ch chan<- os.Signal) {
	signal.Notify(ch, syscall.SIGWINCH)
}
