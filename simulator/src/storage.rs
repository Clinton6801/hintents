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

use soroban_env_host::xdr::{LedgerEntry, LedgerEntryChange};
use std::collections::HashMap;

fn merge_storage_state(before: &[LedgerEntry], changes: &[LedgerEntryChange]) -> Vec<LedgerEntry> {
    let mut state: HashMap<String, LedgerEntry> = HashMap::new();

    // Load BEFORE state
    for entry in before {
        state.insert(format!("{:?}", entry.data), entry.clone());
    }

    // Apply ResultMeta changes
    for change in changes {
        match change {
            LedgerEntryChange::Created(e) | LedgerEntryChange::Updated(e) => {
                state.insert(format!("{:?}", e.data), e.clone());
            }
            LedgerEntryChange::Removed(key) => {
                state.remove(&format!("{:?}", key));
            }
            _ => {}
        }
    }

    state.into_values().collect()
}
