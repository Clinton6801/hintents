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

//! Intentionally faulty contract: triggers an out-of-memory condition.
//!
//! This contract allocates progressively larger vectors until the Soroban
//! host budget for memory bytes is exhausted.  It is used exclusively by
//! the simulator safety test-suite and must never be deployed on-chain.

#![no_std]

use soroban_sdk::{contract, contractimpl, Env, Vec};

#[contract]
pub struct OomContract;

#[contractimpl]
impl OomContract {
    /// Allocates `iterations` nested vectors, each of increasing size, to
    /// exhaust the host memory budget as quickly as possible.
    pub fn run(env: Env, iterations: u32) {
        let mut outer: Vec<Vec<u32>> = Vec::new(&env);
        for i in 0..iterations {
            let mut inner: Vec<u32> = Vec::new(&env);
            for j in 0..i {
                inner.push_back(j);
            }
            outer.push_back(inner);
        }
        // Prevent the compiler from optimising the allocation away.
        let _ = outer.len();
    }
}
