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



import { xdr } from '@stellar/stellar-sdk';

export interface LedgerKey {
    type: xdr.LedgerEntryType;
    key: string;
    hash: string;
}

export interface FootprintResult {
    readOnly: LedgerKey[];
    readWrite: LedgerKey[];
    all: LedgerKey[];
}

// SDK middleware types for custom injection

export interface SDKContext {
    path: string;
    method: string;
    data?: any;
    headers?: Record<string, string>;
    metadata: Record<string, unknown>;
}

export interface SDKResponse<T = any> {
    data: T;
    status: number;
    duration: number;
    endpoint: string;
    metadata: Record<string, unknown>;
}

export type NextFn<T = any> = (ctx: SDKContext) => Promise<SDKResponse<T>>;

export type SDKMiddleware<T = any> = (
    ctx: SDKContext,
    next: NextFn<T>,
) => Promise<SDKResponse<T>>;

// Composes an array of middleware into a single chain around a core handler.
export function composeMiddleware<T = any>(
    middlewares: SDKMiddleware<T>[],
    core: NextFn<T>,
): NextFn<T> {
    return middlewares.reduceRight<NextFn<T>>(
        (next, mw) => (ctx) => mw(ctx, next),
        core,
    );
}
