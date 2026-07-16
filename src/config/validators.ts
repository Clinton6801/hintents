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


import { RPCConfig } from './rpc-config';

export interface RPCConfigValidator {
    validate(cfg: RPCConfig): void;
}

export class UrlsValidator implements RPCConfigValidator {
    validate(cfg: RPCConfig): void {
        if (!cfg.urls || !Array.isArray(cfg.urls) || cfg.urls.length === 0) {
            throw new Error('No RPC URLs configured');
        }
        // each url already validated by parser, but double-check shape
        if (!cfg.urls.every(u => typeof u === 'string' && u.length > 0)) {
            throw new Error('Invalid RPC URLs');
        }
    }
}

export class NumericValidator implements RPCConfigValidator {
    validate(cfg: RPCConfig): void {
        if (!Number.isInteger(cfg.timeout) || cfg.timeout <= 0) {
            throw new Error('timeout must be a positive integer');
        }

        if (!Number.isInteger(cfg.retries) || cfg.retries < 0) {
            throw new Error('retries must be a non-negative integer');
        }

        if (!Number.isInteger(cfg.retryDelay) || cfg.retryDelay <= 0) {
            throw new Error('retryDelay must be a positive integer');
        }

        if (!Number.isInteger(cfg.circuitBreakerThreshold) || cfg.circuitBreakerThreshold <= 0) {
            throw new Error('circuitBreakerThreshold must be a positive integer');
        }

        if (!Number.isInteger(cfg.circuitBreakerTimeout) || cfg.circuitBreakerTimeout <= 0) {
            throw new Error('circuitBreakerTimeout must be a positive integer');
        }

        if (!Number.isInteger(cfg.maxRedirects) || cfg.maxRedirects < 0) {
            throw new Error('maxRedirects must be a non-negative integer');
        }
    }
}
