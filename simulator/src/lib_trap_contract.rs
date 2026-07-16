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

//! Intentionally faulty contract: executes an explicit Wasm trap.
//!
//! The `core::arch::wasm32::unreachable()` intrinsic compiles to a Wasm
//! `unreachable` instruction, which causes the host to trap immediately with
//! a VM-trap error.  This contract is used exclusively by the simulator
//! safety test-suite to verify that the simulator surfaces Wasm traps as
//! structured `HostError` responses rather than panicking.

#![no_std]

use soroban_sdk::{contract, contractimpl, Env};

#[contract]
pub struct TrapContract;

#[contractimpl]
impl TrapContract {
    /// Executes a Wasm `unreachable` instruction unconditionally.
    pub fn run(_env: Env) {
        // SAFETY: intentionally triggers a Wasm trap for testing purposes.
        unsafe { core::arch::wasm32::unreachable() }
    }
}
