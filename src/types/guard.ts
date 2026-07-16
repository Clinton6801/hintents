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



import { ContractID } from './sdk';

type PublicKeyPEM = string & { readonly __brand: 'PublicKeyPEM' };

/**
 * Validates a raw string as a ContractID.
 */
export function isContractID(id: string): id is ContractID {
  // Enforces stricter format: must start with 'C' and be 56 chars
  return /^C[A-Z0-9]{55}$/.test(id);
}

/**
 * Assertion guard to enforce strict linting at the entry point of operations.
 */
export function assertContractID(id: string): asserts id is ContractID {
  if (!isContractID(id)) {
    throw new Error(`Invalid ContractID: ${id}. Strict linting rules require a 56-character 'C' address.`);
  }
}

/**
 * Validates if a string is a valid SPKI PEM Public Key.
 */
export function isPublicKeyPEM(key: string): key is PublicKeyPEM {
  return key.startsWith('-----BEGIN PUBLIC KEY-----') && key.endsWith('-----END PUBLIC KEY-----
');
}
