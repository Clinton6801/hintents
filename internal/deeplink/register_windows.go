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

// checkRegistration queries the Windows registry for the erst:// URL handler.
func checkRegistration(selfPath string) Result {
	res := Result{
		FixSteps: genericFixSteps(),
	}

	// Query HKEY_CLASSES_ROOT\erst to see if the key exists.
	out, err := exec.Command(
		"reg", "query", `HKEY_CLASSES_ROOT\erst`, "/ve",
	).Output()

	if err != nil {
		return res
	}

	value := strings.ToLower(string(out))
	if strings.Contains(value, "url:erst") || strings.Contains(value, "url protocol") {
		res.Registered = true
		res.Handler = strings.TrimSpace(string(out))
		res.FixSteps = nil
	}

	return res
}
