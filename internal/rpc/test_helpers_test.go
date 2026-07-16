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

package rpc

import (
	"encoding/base64"
	"fmt"

	"github.com/stellar/go-stellar-sdk/xdr"
)

// buildValidEntryB64 decodes a base64-encoded XDR LedgerKey, constructs a
// valid LedgerEntry whose key fields match, and returns it as base64 XDR.
// This is used by mock servers so that VerifyLedgerEntryHash passes.
func buildValidEntryB64(keyB64 string) string {
	keyBytes, err := base64.StdEncoding.DecodeString(keyB64)
	if err != nil {
		panic(fmt.Sprintf("buildValidEntryB64: bad base64: %v", err))
	}
	var lk xdr.LedgerKey
	if err := xdr.SafeUnmarshal(keyBytes, &lk); err != nil {
		panic(fmt.Sprintf("buildValidEntryB64: bad XDR key: %v", err))
	}

	var entry xdr.LedgerEntry
	entry.LastModifiedLedgerSeq = 100

	switch lk.Type {
	case xdr.LedgerEntryTypeAccount:
		entry.Data = xdr.LedgerEntryData{
			Type: xdr.LedgerEntryTypeAccount,
			Account: &xdr.AccountEntry{
				AccountId: lk.Account.AccountId,
				Balance:   1000,
			},
		}
	case xdr.LedgerEntryTypeContractCode:
		entry.Data = xdr.LedgerEntryData{
			Type: xdr.LedgerEntryTypeContractCode,
			ContractCode: &xdr.ContractCodeEntry{
				Hash: lk.ContractCode.Hash,
				Code: []byte{0xCA, 0xFE},
			},
		}
	case xdr.LedgerEntryTypeContractData:
		entry.Data = xdr.LedgerEntryData{
			Type: xdr.LedgerEntryTypeContractData,
			ContractData: &xdr.ContractDataEntry{
				Contract:   lk.ContractData.Contract,
				Key:        lk.ContractData.Key,
				Durability: lk.ContractData.Durability,
			},
		}
	case xdr.LedgerEntryTypeTrustline:
		entry.Data = xdr.LedgerEntryData{
			Type: xdr.LedgerEntryTypeTrustline,
			TrustLine: &xdr.TrustLineEntry{
				AccountId: lk.TrustLine.AccountId,
				Asset:     lk.TrustLine.Asset,
				Balance:   500,
				Limit:     1000,
			},
		}
	default:
		panic(fmt.Sprintf("buildValidEntryB64: unsupported key type %v", lk.Type))
	}

	eb, err := entry.MarshalBinary()
	if err != nil {
		panic(fmt.Sprintf("buildValidEntryB64: marshal failed: %v", err))
	}
	return base64.StdEncoding.EncodeToString(eb)
}
