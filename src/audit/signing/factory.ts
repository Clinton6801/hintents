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



import type { AuditSigner } from './types';
import { SoftwareEd25519Signer } from './softwareSigner';
import { Pkcs11Signer } from './pkcs11Signer';
import { KmsSigner } from './kmsSigner';

export type SigningProvider = 'software' | 'pkcs11' | 'kms';

export interface CreateAuditSignerOpts {
  hsmProvider?: string;
  softwarePrivateKeyPem?: string;
  kmsKeyId?: string;
  kmsSigningAlgorithm?: string;
}

export function createAuditSigner(opts: CreateAuditSignerOpts): AuditSigner {
  const provider = (opts.hsmProvider?.toLowerCase() ?? 'software') as SigningProvider;

  switch (provider) {
    case 'kms':
      // Return KMS signer with algorithm support
      return new KmsSigner();

    case 'pkcs11':
      // The Pkcs11Signer now handles algorithm choice via ERST_PKCS11_ALGORITHM env var
      return new Pkcs11Signer();

    case 'software':
      if (!opts.softwarePrivateKeyPem) {
        throw new Error('software signing selected but no private key was provided');
      }
      return new SoftwareEd25519Signer(opts.softwarePrivateKeyPem);

    default:
      throw new Error(`unknown signing provider: "${provider}". Valid options: software, pkcs11, kms`);
  }
}
