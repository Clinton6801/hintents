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

package rpc

import "fmt"

const horizonPageMaxLimit = 200

func normalizePageSize(limit int) int {
	if limit <= 0 {
		return horizonPageMaxLimit
	}
	if limit > horizonPageMaxLimit {
		return horizonPageMaxLimit
	}
	return limit
}

type pageIterator[P any, R any] struct {
	first   func() (P, error)
	next    func(P) (P, error)
	records func(P) []R
	max     int
}

func (it pageIterator[P, R]) collect() ([]R, error) {
	page, err := it.first()
	if err != nil {
		return nil, err
	}

	out := make([]R, 0)
	for {
		rows := it.records(page)
		if len(rows) == 0 {
			return out, nil
		}

		if it.max > 0 {
			remaining := it.max - len(out)
			if remaining <= 0 {
				return out, nil
			}
			if len(rows) > remaining {
				rows = rows[:remaining]
			}
		}
		out = append(out, rows...)

		if it.max > 0 && len(out) >= it.max {
			return out, nil
		}

		page, err = it.next(page)
		if err != nil {
			return out, fmt.Errorf("fetch next page: %w", err)
		}
	}
}
