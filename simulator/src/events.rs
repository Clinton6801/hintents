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

use crate::types::SnapshotMetadata;
use std::sync::atomic::{AtomicU64, Ordering};
use std::time::{SystemTime, UNIX_EPOCH};

static SNAPSHOT_COUNTER: AtomicU64 = AtomicU64::new(1);

pub fn generate_snapshot_id() -> String {
    let ts_micros = SystemTime::now()
        .duration_since(UNIX_EPOCH)
        .map(|d| d.as_micros() as u64)
        .unwrap_or(0);
    let counter = SNAPSHOT_COUNTER.fetch_add(1, Ordering::Relaxed);
    format!("snap-{ts_micros:016x}-{counter:08x}")
}

pub fn build_snapshot_metadata(gas_consumed: u64, call_stack_depth: u32) -> SnapshotMetadata {
    SnapshotMetadata {
        id: generate_snapshot_id(),
        gas_consumed,
        call_stack_depth,
    }
}
