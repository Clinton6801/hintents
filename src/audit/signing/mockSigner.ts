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



import { generateKeyPairSync, sign as nodeSign } from 'crypto';
import type { AuditSigner, PublicKey, Signature, HardwareAttestation } from './types';

/**
 * In-memory signer used for unit tests; does not require a real HSM.
 */
export class MockAuditSigner implements AuditSigner {
  private readonly privateKeyPem: string;
  private readonly publicKeyPem: string;
  private readonly mockAttestation: HardwareAttestation | undefined;

  constructor(opts?: { withAttestation?: boolean }) {
    const { publicKey, privateKey } = generateKeyPairSync('ed25519', {
      publicKeyEncoding: { type: 'spki', format: 'pem' },
      privateKeyEncoding: { type: 'pkcs8', format: 'pem' },
    });
    this.privateKeyPem = privateKey;
    this.publicKeyPem = publicKey;

    if (opts?.withAttestation) {
      this.mockAttestation = {
        certificates: [
          {
            pem: '-----BEGIN CERTIFICATE-----
MOCK_LEAF_CERT
-----END CERTIFICATE-----',
            subject: 'CN=MockHSM Attestation Key',
            issuer: 'CN=MockHSM Root CA',
            serial: 'deadbeef01',
          },
          {
            pem: '-----BEGIN CERTIFICATE-----
MOCK_ROOT_CERT
-----END CERTIFICATE-----',
            subject: 'CN=MockHSM Root CA',
            issuer: 'CN=MockHSM Root CA',
            serial: 'cafebabe02',
          },
        ],
        token_info: 'MockHSM v1.0 (Test Manufacturer)',
        key_non_exportable: true,
        retrieved_at: new Date().toISOString(),
      };
    }
  }

  async sign(payload: Uint8Array): Promise<Signature> {
    return nodeSign(null, Buffer.from(payload), this.privateKeyPem);
  }

  async public_key(): Promise<PublicKey> {
    return this.publicKeyPem;
  }

  async attestation_chain(): Promise<HardwareAttestation | undefined> {
    return this.mockAttestation;
  }
}
