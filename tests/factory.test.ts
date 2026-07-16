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



import { createAuditSigner } from '../src/audit/signing/factory';
import { KmsSigner } from '../src/audit/signing/kmsSigner';

describe('Audit signer factory', () => {
  const mockKeyId = 'arn:aws:kms:us-east-1:123456789012:key/12345678-1234-1234-1234-123456789012';
  const mockPublicKeyPem = `-----BEGIN PUBLIC KEY-----
MFMwEwYHKoZIzj0CAQYIKoZIzj0DAQcDOgAEWdp8vGtXxyGkftJoJphBnwvlvVfc
6xwvSMu00nWXrF5bUegdisSGSF3567890123456789abcdefg
-----END PUBLIC KEY-----`;

  beforeEach(() => {
    process.env.ERST_KMS_KEY_ID = mockKeyId;
    process.env.ERST_KMS_PUBLIC_KEY_PEM = mockPublicKeyPem;
    process.env.AWS_REGION = 'us-east-1';
  });

  afterEach(() => {
    delete process.env.ERST_KMS_KEY_ID;
    delete process.env.ERST_KMS_PUBLIC_KEY_PEM;
  });

  test('creates KMS signer when provider is kms', () => {
    const signer = createAuditSigner({ hsmProvider: 'kms' });
    expect(signer).toBeInstanceOf(KmsSigner);
  });

  test('respects case-insensitive provider selection', () => {
    const signer = createAuditSigner({ hsmProvider: 'KMS' });
    expect(signer).toBeInstanceOf(KmsSigner);
  });

  test('defaults to software signer when no provider specified', () => {
    const privateKey = process.env.TEST_PRIVATE_KEY_PEM || 'test-key-placeholder';
    const signer = createAuditSigner({ softwarePrivateKeyPem: privateKey });
    expect(signer).toBeDefined();
  });
});
