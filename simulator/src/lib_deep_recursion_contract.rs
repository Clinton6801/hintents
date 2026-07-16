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

//! Intentionally faulty contract: triggers unbounded deep recursion.
//!
//! The Soroban host enforces a call-stack depth limit.  This contract
//! exceeds that limit by calling itself recursively until the host traps.
//! It is used exclusively by the simulator safety test-suite.

#![no_std]

use soroban_sdk::{contract, contractimpl, Env};

#[contract]
pub struct DeepRecursionContract;

#[contractimpl]
impl DeepRecursionContract {
    /// Recurses `depth` times.  Pass a value larger than the host's maximum
    /// call-stack depth (typically 10) to trigger a trap.
    pub fn recurse(env: Env, depth: u32) -> u32 {
        if depth == 0 {
            return 0;
        }
        // Re-invoke the current contract to consume call-stack depth.
        let current = env.current_contract_address();
        let client = DeepRecursionContractClient::new(&env, &current);
        1 + client.recurse(&(depth - 1))
    }
}
