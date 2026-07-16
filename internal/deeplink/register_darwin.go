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

package deeplink

import (
	"os/exec"
	"strings"
)

// checkRegistration queries Launch Services to find the handler for erst://.
func checkRegistration(selfPath string) Result {
	res := Result{
		FixSteps: genericFixSteps(),
	}

	// `lsregister -dump` is heavy; use the lighter `open -Ra` probe instead.
	// We ask the system which app handles the scheme by querying with a dry-run.
	out, err := exec.Command("bash", "-c",
		`/System/Library/Frameworks/CoreServices.framework/Frameworks/LaunchServices.framework/Support/lsregister -dump 2>/dev/null | grep -i "erst://" | head -5`,
	).Output()

	if err == nil && strings.Contains(strings.ToLower(string(out)), "erst") {
		res.Registered = true
		res.Handler = strings.TrimSpace(string(out))
		res.FixSteps = nil
		return res
	}

	// Fallback: try `open -Ra erst://` which exits 0 when a handler exists.
	if err2 := exec.Command("open", "-Ra", "erst://").Run(); err2 == nil {
		res.Registered = true
		res.Handler = "registered (via open -Ra)"
		res.FixSteps = nil
		return res
	}

	return res
}
