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

//! Zstd decompression for ledger snapshot payloads sent over IPC.
//!
//! The Go side compresses `ledger_entries` with Zstd and base64-encodes the
//! result into `ledger_entries_zstd`.  This module decodes and decompresses
//! that field back into the plain `HashMap<String, String>` the simulator
//! expects.

use base64::Engine as _;
use std::collections::HashMap;

use super::types::IpcError;

/// Decodes a base64-encoded Zstd blob produced by `bridge.CompressRequest`
/// and returns the original `ledger_entries` map.
pub fn decompress_ledger_entries(b64: &str) -> Result<HashMap<String, String>, IpcError> {
    let compressed = base64::engine::general_purpose::STANDARD
        .decode(b64)
        .map_err(|e| IpcError::Decompress(format!("base64 decode: {e}")))?;

    let raw = zstd::decode_all(compressed.as_slice())
        .map_err(|e| IpcError::Decompress(format!("zstd decode: {e}")))?;

    serde_json::from_slice(&raw).map_err(IpcError::from)
}
