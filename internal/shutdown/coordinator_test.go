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

package shutdown

import (
	"context"
	"testing"
)

func TestCoordinatorRun_LIFOAndOnce(t *testing.T) {
	c := NewCoordinator()
	order := make([]string, 0, 3)

	c.Register("first", func(ctx context.Context) error {
		_ = ctx
		order = append(order, "first")
		return nil
	})
	c.Register("second", func(ctx context.Context) error {
		_ = ctx
		order = append(order, "second")
		return nil
	})
	c.Register("third", func(ctx context.Context) error {
		_ = ctx
		order = append(order, "third")
		return nil
	})

	if err := c.Run(context.Background()); err != nil {
		t.Fatalf("unexpected run error: %v", err)
	}

	want := []string{"third", "second", "first"}
	if len(order) != len(want) {
		t.Fatalf("unexpected hook count: got %d want %d", len(order), len(want))
	}
	for i := range want {
		if order[i] != want[i] {
			t.Fatalf("unexpected order at %d: got %s want %s", i, order[i], want[i])
		}
	}

	// Second run should be a no-op.
	order = order[:0]
	if err := c.Run(context.Background()); err != nil {
		t.Fatalf("unexpected second run error: %v", err)
	}
	if len(order) != 0 {
		t.Fatalf("expected no hooks on second run, got %d", len(order))
	}
}
