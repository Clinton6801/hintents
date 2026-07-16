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

package cmd

import (
	"context"
	"sync"
	"time"

	"github.com/dotandev/hintents/internal/logger"
	"github.com/dotandev/hintents/internal/rpc"
	"github.com/dotandev/hintents/internal/shutdown"
	"github.com/dotandev/hintents/internal/simulator"
)

const shutdownTimeout = 3 * time.Second

var shutdownState struct {
	mu          sync.RWMutex
	coordinator *shutdown.Coordinator
}

func setShutdownCoordinator(c *shutdown.Coordinator) {
	shutdownState.mu.Lock()
	defer shutdownState.mu.Unlock()
	shutdownState.coordinator = c
}

func clearShutdownCoordinator() {
	shutdownState.mu.Lock()
	defer shutdownState.mu.Unlock()
	shutdownState.coordinator = nil
}

func registerShutdownHook(name string, fn shutdown.HookFunc) {
	shutdownState.mu.RLock()
	c := shutdownState.coordinator
	shutdownState.mu.RUnlock()
	if c == nil {
		return
	}
	c.Register(name, fn)
}

func runShutdownHooksWithTimeout(c *shutdown.Coordinator, timeout time.Duration) {
	if c == nil {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	if err := c.Run(ctx); err != nil {
		logger.Logger.Warn("Shutdown hooks completed with errors", "error", err)
	}
}

func registerCacheFlushHook() {
	registerShutdownHook("rpc-cache-flush", func(ctx context.Context) error {
		if err := rpc.Flush(ctx); err != nil {
			logger.Logger.Warn("Failed to flush RPC cache during shutdown", "error", err)
		}
		return nil
	})
}

func registerRunnerCloseHook(name string, runner simulator.RunnerInterface) {
	if runner == nil {
		return
	}

	registerShutdownHook(name, func(ctx context.Context) error {
		_ = ctx
		return runner.Close()
	})
}
