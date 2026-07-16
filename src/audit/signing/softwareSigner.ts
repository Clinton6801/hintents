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



import { sign, createPublicKey } from 'crypto';
import type { AuditSigner, PublicKey, Signature } from './types';

/**
 * Default signer that uses a local Ed25519 private key (PKCS#8 PEM).
 */
export class SoftwareEd25519Signer implements AuditSigner {
  constructor(private readonly privateKeyPem: string) {}

  async sign(payload: Uint8Array): Promise<Signature> {
    try {
      return sign(null, Buffer.from(payload), this.privateKeyPem);
    } catch (e) {
      const msg = e instanceof Error ? e.message : String(e);
      throw new Error(`software signing failed: ${msg}`);
    }
  }

  async public_key(): Promise<PublicKey> {
    try {
      const pub = createPublicKey(this.privateKeyPem);
      return pub.export({ type: 'spki', format: 'pem' }).toString();
    } catch (e) {
      const msg = e instanceof Error ? e.message : String(e);
      throw new Error(`software public key derivation failed: ${msg}`);
    }
  }
}
