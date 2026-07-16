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

#![allow(clippy::pedantic, clippy::nursery, dead_code)]

pub mod asset_tracker;
pub mod context;
pub mod gas_optimizer;
pub mod git_detector;
pub mod hsm;
pub mod ipc;
pub mod runner;
pub mod snapshot;
pub mod source_map_cache;
pub mod source_mapper;
pub mod stack_trace;
pub mod state;
pub mod types;
pub mod wasm_types;

#[cfg(test)]
mod tests;
