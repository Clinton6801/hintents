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

//go:build !darwin && !windows

package deeplink

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// checkRegistration checks xdg-mime and .desktop file presence on Linux/BSD.
func checkRegistration(selfPath string) Result {
	res := Result{
		FixSteps: genericFixSteps(),
	}

	// Ask xdg-mime which application handles the scheme.
	out, err := exec.Command("xdg-mime", "query", "default", "x-scheme-handler/erst").Output()
	if err == nil {
		handler := strings.TrimSpace(string(out))
		if handler != "" {
			res.Registered = true
			res.Handler = handler
			res.FixSteps = nil
			return res
		}
	}

	// Fallback: scan known .desktop directories for an erst handler.
	desktopDirs := []string{
		filepath.Join(os.Getenv("HOME"), ".local", "share", "applications"),
		"/usr/share/applications",
		"/usr/local/share/applications",
	}

	for _, dir := range desktopDirs {
		entries, err := os.ReadDir(dir)
		if err != nil {
			continue
		}
		for _, e := range entries {
			if !strings.HasSuffix(e.Name(), ".desktop") {
				continue
			}
			data, err := os.ReadFile(filepath.Join(dir, e.Name()))
			if err != nil {
				continue
			}
			content := string(data)
			if strings.Contains(content, "x-scheme-handler/erst") {
				res.Registered = true
				res.Handler = filepath.Join(dir, e.Name())
				res.FixSteps = nil
				return res
			}
		}
	}

	return res
}
