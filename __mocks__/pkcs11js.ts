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

const { privateKey } = generateKeyPairSync('ed25519', {
  publicKeyEncoding: { type: 'spki', format: 'pem' },
  privateKeyEncoding: { type: 'pkcs8', format: 'pem' },
});

type MockState = {
  loadedModule?: string;
  lastTemplate?: unknown;
  initialized?: boolean;
};

const state: MockState = {};

export const CKF_SERIAL_SESSION = 0x00000004;
export const CKF_RW_SESSION = 0x00000002;
export const CKA_CLASS = 0x00000000;
export const CKO_PRIVATE_KEY = 0x00000003;
export const CKA_LABEL = 0x00000003;
export const CKA_ID = 0x00000102;
export const CKM_EDDSA = 0x00001050;

export class PKCS11 {
  load(modulePath: string): void {
    state.loadedModule = modulePath;
  }

  C_Initialize(): void {
    state.initialized = true;
  }

  C_GetSlotList(_tokenPresent: boolean): number[] {
    return [1];
  }

  C_OpenSession(_slot: number, _flags: number): number {
    return 1;
  }

  C_Login(_session: number, _userType: number, _pin?: string): void {
    // Accept any pin for mock purposes.
  }

  C_FindObjectsInit(_session: number, template: unknown): void {
    state.lastTemplate = template;
  }

  C_FindObjects(_session: number, _count: number): number[] {
    return [1];
  }

  C_FindObjectsFinal(_session: number): void {
    // no-op
  }

  C_SignInit(_session: number, _mechanism: unknown, _key: unknown): void {
    // no-op
  }

  C_Sign(_session: number, payload: Buffer): Buffer {
    return nodeSign(null, payload, privateKey);
  }

  C_CloseSession(_session: number): void {
    // no-op
  }

  C_Finalize(): void {
    // no-op
  }
}

export const __getState = (): MockState => ({ ...state });

export default {
  PKCS11,
  CKF_SERIAL_SESSION,
  CKF_RW_SESSION,
  CKA_CLASS,
  CKO_PRIVATE_KEY,
  CKA_LABEL,
  CKA_ID,
  CKM_EDDSA,
  __getState,
};
