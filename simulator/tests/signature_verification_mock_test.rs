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

//! Tests for signature verification mocking functionality

use erst_sim::types::SimulationRequest;

#[test]
fn test_signature_verification_mock_true() {
    let request = SimulationRequest {
        envelope_xdr: String::new(),
        result_meta_xdr: String::new(),
        ledger_entries: None,
        control_command: None,
        rewind_step: None,
        fork_params: None,
        harness_reset: false,
        ledger_entries_zstd: None,
        contract_wasm: None,
        wasm_path: None,
        no_cache: false,
        enable_optimization_advisor: false,
        profile: None,
        _timestamp: None,
        timestamp: String::new(),
        mock_base_fee: None,
        mock_gas_price: None,
        mock_signature_verification: Some(true),
        enable_coverage: false,
        coverage_lcov_path: None,
        resource_calibration: None,
        memory_limit: None,
        restore_preamble: None,
        include_linear_memory: false,
        enable_asset_safety: false,
    };

    assert_eq!(request.mock_signature_verification, Some(true));
}

#[test]
fn test_signature_verification_mock_false() {
    let request = SimulationRequest {
        envelope_xdr: String::new(),
        result_meta_xdr: String::new(),
        ledger_entries: None,
        control_command: None,
        rewind_step: None,
        fork_params: None,
        harness_reset: false,
        ledger_entries_zstd: None,
        contract_wasm: None,
        wasm_path: None,
        no_cache: false,
        enable_optimization_advisor: false,
        profile: None,
        _timestamp: None,
        timestamp: String::new(),
        mock_base_fee: None,
        mock_gas_price: None,
        mock_signature_verification: Some(false),
        enable_coverage: false,
        coverage_lcov_path: None,
        resource_calibration: None,
        memory_limit: None,
        restore_preamble: None,
        include_linear_memory: false,
        enable_asset_safety: false,
    };

    assert_eq!(request.mock_signature_verification, Some(false));
}

#[test]
fn test_signature_verification_mock_disabled() {
    let request = SimulationRequest {
        envelope_xdr: String::new(),
        result_meta_xdr: String::new(),
        ledger_entries: None,
        control_command: None,
        rewind_step: None,
        fork_params: None,
        harness_reset: false,
        ledger_entries_zstd: None,
        contract_wasm: None,
        wasm_path: None,
        no_cache: false,
        enable_optimization_advisor: false,
        profile: None,
        _timestamp: None,
        timestamp: String::new(),
        mock_base_fee: None,
        mock_gas_price: None,
        mock_signature_verification: None,
        enable_coverage: false,
        coverage_lcov_path: None,
        resource_calibration: None,
        memory_limit: None,
        restore_preamble: None,
        include_linear_memory: false,
        enable_asset_safety: false,
    };

    assert_eq!(request.mock_signature_verification, None);
}
