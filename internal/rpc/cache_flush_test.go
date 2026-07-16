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

import (
	"context"
	"testing"
)

func TestFlush_NoOp(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	if err := Flush(ctx); err != nil {
		t.Fatalf("expected nil, got %v", err)
	}
}
