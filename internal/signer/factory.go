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


// Copyright 2026 Erst Users

package signer

import "os"

// NewFromEnv creates a Signer based on the ERST_SIGNER_TYPE environment
// variable. When the variable is absent or set to "software", an
// InMemorySigner is returned using the hex key from
// ERST_SOFTWARE_PRIVATE_KEY_HEX. When set to "pkcs11", a Pkcs11Signer
// is created from ERST_PKCS11_* environment variables.
func NewFromEnv() (Signer, error) {
	signerType := os.Getenv("ERST_SIGNER_TYPE")
	if signerType == "" {
		signerType = "software"
	}

	switch signerType {
	case "software":
		keyHex := os.Getenv("ERST_SOFTWARE_PRIVATE_KEY_HEX")
		if keyHex == "" {
			return nil, &Error{Op: "factory", Msg: "ERST_SOFTWARE_PRIVATE_KEY_HEX is required for software signer"}
		}
		return NewInMemorySigner(keyHex)

	case "pkcs11":
		cfg, err := Pkcs11ConfigFromEnv()
		if err != nil {
			return nil, err
		}
		return NewPkcs11Signer(*cfg)

	default:
		return nil, &Error{Op: "factory", Msg: "unsupported ERST_SIGNER_TYPE: " + signerType}
	}
}
