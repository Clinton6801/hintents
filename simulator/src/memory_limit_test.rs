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

//! Test module for memory limit simulation functionality

#[cfg(test)]
mod tests {
    use crate::runner::SimHost;
    use crate::types::ResourceCalibration;
    use std::panic;

    #[test]
    fn test_memory_limit_field() {
        // Test that SimHost can be created with a memory limit
        let memory_limit = Some(1000000); // 1MB limit
        let host = SimHost::new(None, None, memory_limit);

        assert_eq!(host.memory_limit, memory_limit);
    }

    #[test]
    fn test_no_memory_limit() {
        // Test that SimHost can be created without memory limit
        let host = SimHost::new(None, None, None);

        assert_eq!(host.memory_limit, None);
    }

    #[test]
    fn test_memory_limit_check_no_panic() {
        // Test memory limit checking functionality when within limits
        let memory_limit = Some(1000); // Very small limit
        let host = SimHost::new(None, None, memory_limit);

        // This should not panic as we haven't executed any operations yet
        host.check_memory_limit();
    }

    #[test]
    fn test_memory_limit_exceeded_does_not_propagate_panic() {
        // Use catch_unwind to ensure panics from check_memory_limit do not
        // abort the entire test process.
        let memory_limit = Some(100);
        let host = SimHost::new(None, None, memory_limit);

        let result = panic::catch_unwind(|| {
            host.check_memory_limit();
        });

        // We only assert that a panic is observed here, without relying on
        // the exact panic message format.
        assert!(result.is_err());
    }
}
