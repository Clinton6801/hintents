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



/* eslint-disable @typescript-eslint/no-unused-vars */

import { createAuditSigner } from "../src/audit/signing/factory";
import { KmsSigner } from "../src/audit/signing/kmsSigner";
import { SoftwareEd25519Signer } from "../src/audit/signing/softwareSigner";
import { Pkcs11Signer } from "../src/audit/signing/pkcs11Signer";

jest.mock("../src/audit/signing/kmsSigner");
jest.mock("../src/audit/signing/softwareSigner");
jest.mock("../src/audit/signing/pkcs11Signer");

describe("AuditSigner Factory", () => {
  const mockPrivateKey = process.env.TEST_PRIVATE_KEY_PEM || "";
  const mockPublicKey = process.env.TEST_PUBLIC_KEY_PEM || "";

  beforeEach(() => {
    process.env.ERST_KMS_KEY_ID = "arn:aws:kms:us-east-1:123456789012:key/test";
    process.env.ERST_KMS_PUBLIC_KEY_PEM = mockPublicKey;
    process.env.AWS_REGION = "us-east-1";
    jest.clearAllMocks();
  });

  afterEach(() => {
    delete process.env.ERST_KMS_KEY_ID;
    delete process.env.ERST_KMS_PUBLIC_KEY_PEM;
    delete process.env.ERST_PKCS11_MODULE;
  });

  test("creates KMS signer when provider is kms", () => {
    createAuditSigner({ hsmProvider: "kms" });
    expect(KmsSigner).toHaveBeenCalled();
  });

  test("creates KMS signer with case-insensitive provider", () => {
    createAuditSigner({ hsmProvider: "KMS" });
    expect(KmsSigner).toHaveBeenCalled();
  });

  test("creates software signer when provider is software", () => {
    createAuditSigner({
      hsmProvider: "software",
      softwarePrivateKeyPem: mockPrivateKey,
    });
    expect(SoftwareEd25519Signer).toHaveBeenCalledWith(mockPrivateKey);
  });

  test("creates software signer by default", () => {
    createAuditSigner({ softwarePrivateKeyPem: mockPrivateKey });
    expect(SoftwareEd25519Signer).toHaveBeenCalledWith(mockPrivateKey);
  });

  test("creates PKCS#11 signer when provider is pkcs11", () => {
    process.env.ERST_PKCS11_MODULE = "/usr/lib/softhsm/libsofthsm2.so";
    createAuditSigner({ hsmProvider: "pkcs11" });
    expect(Pkcs11Signer).toHaveBeenCalled();
  });

  test("throws when software provider without private key", () => {
    expect(() => createAuditSigner({ hsmProvider: "software" })).toThrow(
      "software signing selected but no private key was provided",
    );
  });
});
