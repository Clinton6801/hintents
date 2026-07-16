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


import { Pkcs11Ed25519Signer } from '../src/audit/signing/pkcs11Signer';

jest.mock('pkcs11js');

describe('Pkcs11Ed25519Signer signing (mock module)', () => {
  const originalEnv = process.env;

  beforeEach(() => {
    jest.resetModules();
    process.env = { ...originalEnv };
    process.env.ERST_PKCS11_MODULE = '/tmp/mock-pkcs11.so';
    process.env.ERST_AUDIT_HSM_RATE_LIMIT_FILE = '/tmp/erst_audit_hsm_calls_test.json';
    process.env.ERST_PKCS11_PIN = '1234';
    process.env.ERST_PKCS11_KEY_LABEL = 'test-key';
    process.env.ERST_PKCS11_PUBLIC_KEY_PEM = '-----BEGIN PUBLIC KEY-----
mock
-----END PUBLIC KEY-----';
  });

  afterEach(() => {
    process.env = originalEnv;
  });

  test('signs payload using mocked PKCS#11 module', async () => {
    const signer = new Pkcs11Ed25519Signer();
    const payload = Buffer.from('hello-world');

    const signature = await signer.sign(payload);

    expect(signature).toBeInstanceOf(Uint8Array);
    expect(signature.length).toBe(64);
  });

  test('uses configured module path and key label for lookup', async () => {
    const signer = new Pkcs11Ed25519Signer();
    await signer.sign(Buffer.from('another-payload'));

    const pkcs11Mock = jest.requireMock('pkcs11js') as { __getState: () => { loadedModule?: string; lastTemplate?: unknown } };
    const state = pkcs11Mock.__getState();

    expect(state.loadedModule).toBe('/tmp/mock-pkcs11.so');
    expect(JSON.stringify(state.lastTemplate)).toContain('test-key');
  });
});
