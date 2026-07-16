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

//go:build windows

package sourcemap

import (
	"fmt"
	"os"
)

func (sc *SourceCache) acquireLock(entryPath string, exclusive bool) (*os.File, error) {
	lp := sc.lockPath(entryPath)
	lf, err := os.OpenFile(lp, os.O_CREATE|os.O_RDWR, 0600)
	if err != nil {
		return nil, fmt.Errorf("failed to open lock file %q: %w", lp, err)
	}
	return lf, nil
}

func (sc *SourceCache) releaseLock(lf *os.File) {
	_ = lf.Close()
}
