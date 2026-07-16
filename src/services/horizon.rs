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

use serde::Deserialize;

use crate::error::AppError;

/// Client responsible for communicating with the Stellar Horizon API.
#[derive(Clone, Debug)]
pub struct HorizonClient {
    base_url: String,
}

impl HorizonClient {
    /// Create a new Horizon client.
    pub fn new(base_url: impl Into<String>) -> Self {
        Self {
            base_url: base_url.into(),
        }
    }

    /// Get the base URL of the Horizon server.
    pub fn base_url(&self) -> &str {
        &self.base_url
    }

    /// Fetch the latest transaction from Horizon.
    ///
    /// NOTE:
    /// - Networking is added in a later issue
    /// - This stub prevents accidental Horizon usage early
    pub fn fetch_latest_transaction(
        &self,
    ) -> Result<HorizonTransaction, AppError> {
        Err(AppError::Network(
            "Horizon client not implemented yet".into(),
        ))
    }

    /// Fetch operations for a given transaction hash.
    pub fn fetch_operations(
        &self,
        _tx_hash: &str,
    ) -> Result<Vec<HorizonOperation>, AppError> {
        Err(AppError::Network(
            "Horizon client not implemented yet".into(),
        ))
    }
}

/// Represents a transaction returned by Horizon.
#[derive(Debug, Deserialize)]
pub struct HorizonTransaction {
    pub hash: String,
    pub successful: bool,
    pub fee_charged: String,
}

/// Represents an operation within a transaction.
#[derive(Debug, Deserialize)]
pub struct HorizonOperation {
    #[serde(rename = "type")]
    pub op_type: String,

    pub from: Option<String>,
    pub to: Option<String>,

    pub asset_type: Option<String>,
    pub asset_code: Option<String>,
    pub asset_issuer: Option<String>,

    pub amount: Option<String>,
}
