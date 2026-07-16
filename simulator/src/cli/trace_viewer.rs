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

use std::path::PathBuf;

#[allow(dead_code)]
fn trace_viewer_temp_root() -> PathBuf {
    std::env::temp_dir().join("erst-trace-viewer")
}

#[allow(dead_code)]
pub fn trace_viewer_temp_path(file_name: &str) -> PathBuf {
    trace_viewer_temp_root().join(file_name)
}

#[allow(dead_code)]
pub fn render_trace() {
    tracing::info!(kind = "span", "User logged in");
    tracing::error!(kind = "error", "Connection failed");
}
