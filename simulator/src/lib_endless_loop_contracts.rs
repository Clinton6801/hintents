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

//! Intentionally faulty contract: triggers CPU budget exhaustion via an
//! endless loop.
//!
//! The Soroban host meters every Wasm instruction.  A non-terminating loop
//! will consume all available CPU instructions and cause the host to trap
//! with a budget-exceeded error.  This contract is used exclusively by the
//! simulator safety test-suite.

#![no_std]

use soroban_sdk::{contract, contractimpl, Env};

#[contract]
pub struct EndlessLoopContract;

#[contractimpl]
impl EndlessLoopContract {
    /// Enters a loop that never terminates.  The host will trap once the CPU
    /// instruction budget is exhausted.
    pub fn run(_env: Env) {
        // A plain `loop {}` is sufficient; the host metering will fire before
        // the Wasm engine can spin indefinitely on real hardware.
        #[allow(clippy::empty_loop)]
        loop {}
    }
}
