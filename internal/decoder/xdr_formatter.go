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

package decoder

import (
	"bytes"
	"encoding/json"
	"fmt"
	"text/tabwriter"

	"github.com/stellar/go-stellar-sdk/xdr"
)

type FormatType string

const (
	FormatJSON  FormatType = "json"
	FormatTable FormatType = "table"
)

type XDRFormatter struct {
	format FormatType
}

func NewXDRFormatter(format FormatType) *XDRFormatter {
	return &XDRFormatter{format: format}
}

func (f *XDRFormatter) Format(data interface{}) (string, error) {
	switch f.format {
	case FormatJSON:
		return f.formatJSON(data)
	case FormatTable:
		return f.formatTable(data)
	default:
		return "", fmt.Errorf("unsupported format: %s", f.format)
	}
}

func (f *XDRFormatter) formatJSON(data interface{}) (string, error) {
	output, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return "", fmt.Errorf("failed to marshal JSON: %w", err)
	}
	return string(output), nil
}

func (f *XDRFormatter) formatTable(data interface{}) (string, error) {
	switch v := data.(type) {
	case *xdr.LedgerEntry:
		return formatLedgerEntryTable(v)
	case *xdr.TransactionEnvelope:
		return formatTransactionEnvelopeTable(v)
	case *xdr.DiagnosticEvent:
		return formatDiagnosticEventTable(v)
	case []interface{}:
		return formatGenericTable(v)
	default:
		var buf bytes.Buffer
		w := tabwriter.NewWriter(&buf, 0, 0, 2, ' ', 0)
		_, _ = fmt.Fprintf(w, "Type:	%T
", v)
		_, _ = fmt.Fprintf(w, "Value:	%v
", v)
		_ = w.Flush()
		return buf.String(), nil
	}
}

func formatLedgerEntryTable(entry *xdr.LedgerEntry) (string, error) {
	var buf bytes.Buffer
	w := tabwriter.NewWriter(&buf, 0, 0, 2, ' ', 0)

	_, _ = fmt.Fprintf(w, "Type:	%v
", entry.Data.Type)
	_, _ = fmt.Fprintf(w, "Last Modified Ledger:	%d
", entry.LastModifiedLedgerSeq)

	switch entry.Data.Type {
	case xdr.LedgerEntryTypeAccount:
		if entry.Data.Account != nil {
			acc := entry.Data.Account
			_, _ = fmt.Fprintf(w, "Account ID:	%s
", acc.AccountId.Address())
			_, _ = fmt.Fprintf(w, "Balance:	%d
", acc.Balance)
			_, _ = fmt.Fprintf(w, "Sequence:	%d
", acc.SeqNum)
			_, _ = fmt.Fprintf(w, "Flags:	%d
", acc.Flags)
		}

	case xdr.LedgerEntryTypeTrustline:
		if entry.Data.TrustLine != nil {
			tl := entry.Data.TrustLine
			_, _ = fmt.Fprintf(w, "Account:	%s
", tl.AccountId.Address())
			_, _ = fmt.Fprintf(w, "Asset Type:	%v
", tl.Asset.Type)
			_, _ = fmt.Fprintf(w, "Balance:	%d
", tl.Balance)
			_, _ = fmt.Fprintf(w, "Flags:	%d
", tl.Flags)
		}

	case xdr.LedgerEntryTypeOffer:
		if entry.Data.Offer != nil {
			offer := entry.Data.Offer
			_, _ = fmt.Fprintf(w, "Seller:	%s
", offer.SellerId.Address())
			_, _ = fmt.Fprintf(w, "Offer ID:	%d
", offer.OfferId)
			_, _ = fmt.Fprintf(w, "Amount:	%d
", offer.Amount)
		}

	case xdr.LedgerEntryTypeData:
		if entry.Data.Data != nil {
			data := entry.Data.Data
			_, _ = fmt.Fprintf(w, "Account:	%s
", data.AccountId.Address())
			_, _ = fmt.Fprintf(w, "Data Name:	%s
", data.DataName)
			_, _ = fmt.Fprintf(w, "Data Value (bytes):	%d
", len(data.DataValue))
		}

	case xdr.LedgerEntryTypeClaimableBalance:
		if entry.Data.ClaimableBalance != nil {
			cb := entry.Data.ClaimableBalance
			if cb.BalanceId.V0 != nil {
				_, _ = fmt.Fprintf(w, "Balance ID:	%x
", *cb.BalanceId.V0)
			}
			_, _ = fmt.Fprintf(w, "Amount:	%d
", cb.Amount)
		}

	case xdr.LedgerEntryTypeContractData:
		if entry.Data.ContractData != nil {
			cd := entry.Data.ContractData
			_, _ = fmt.Fprintf(w, "Durability:	%v
", cd.Durability)
		}

	case xdr.LedgerEntryTypeContractCode:
		if entry.Data.ContractCode != nil {
			cc := entry.Data.ContractCode
			_, _ = fmt.Fprintf(w, "Code Hash:	%x
", cc.Hash)
			_, _ = fmt.Fprintf(w, "Code Size:	%d bytes
", len(cc.Code))
		}
	}

	_ = w.Flush()
	return buf.String(), nil
}

func formatTransactionEnvelopeTable(env *xdr.TransactionEnvelope) (string, error) {
	var buf bytes.Buffer
	w := tabwriter.NewWriter(&buf, 0, 0, 2, ' ', 0)

	_, _ = fmt.Fprintf(w, "Envelope Type:	%v
", env.Type)

	switch env.Type {
	case xdr.EnvelopeTypeEnvelopeTypeTxV0:
		if env.V0 != nil {
			tx := env.V0.Tx
			_, _ = fmt.Fprintf(w, "Fee:	%d
", tx.Fee)
			_, _ = fmt.Fprintf(w, "Sequence Num:	%d
", tx.SeqNum)
			_, _ = fmt.Fprintf(w, "Operations:	%d
", len(tx.Operations))
		}

	case xdr.EnvelopeTypeEnvelopeTypeTx:
		if env.V1 != nil {
			tx := env.V1.Tx
			_, _ = fmt.Fprintf(w, "Source Account:	%s
", tx.SourceAccount.Address())
			_, _ = fmt.Fprintf(w, "Fee:	%d
", tx.Fee)
			_, _ = fmt.Fprintf(w, "Sequence Num:	%d
", tx.SeqNum)
			_, _ = fmt.Fprintf(w, "Operations:	%d
", len(tx.Operations))
		}

	case xdr.EnvelopeTypeEnvelopeTypeTxFeeBump:
		if env.FeeBump != nil {
			feeBump := env.FeeBump.Tx
			_, _ = fmt.Fprintf(w, "Fee Source:	%s
", feeBump.FeeSource.Address())
			_, _ = fmt.Fprintf(w, "Fee:	%d
", feeBump.Fee)
		}
	}

	_ = w.Flush()
	return buf.String(), nil
}

func formatDiagnosticEventTable(event *xdr.DiagnosticEvent) (string, error) {
	var buf bytes.Buffer
	w := tabwriter.NewWriter(&buf, 0, 0, 2, ' ', 0)

	_, _ = fmt.Fprintf(w, "Successful:	%v
", event.InSuccessfulContractCall)
	_, _ = fmt.Fprintf(w, "Event Type:	%v
", event.Event.Type)

	if event.Event.ContractId != nil {
		_, _ = fmt.Fprintf(w, "Contract ID:	%x
", event.Event.ContractId)
	}

	_ = w.Flush()
	return buf.String(), nil
}

func formatGenericTable(items []interface{}) (string, error) {
	var buf bytes.Buffer
	w := tabwriter.NewWriter(&buf, 0, 0, 2, ' ', 0)

	_, _ = fmt.Fprintf(w, "Items:	%d
", len(items))

	for i, item := range items {
		if i > 0 {
			_, _ = fmt.Fprintf(w, "
")
		}
		_, _ = fmt.Fprintf(w, "Item %d:	%T
", i, item)
	}

	_ = w.Flush()
	return buf.String(), nil
}

func DecodeXDRBase64AsLedgerEntry(data string) (*xdr.LedgerEntry, error) {
	var entry xdr.LedgerEntry
	if err := entry.UnmarshalBinary([]byte(data)); err != nil {
		return nil, fmt.Errorf("failed to decode ledger entry: %w", err)
	}
	return &entry, nil
}

func DecodeXDRBase64AsDiagnosticEvent(data string) (*xdr.DiagnosticEvent, error) {
	var event xdr.DiagnosticEvent
	if err := event.UnmarshalBinary([]byte(data)); err != nil {
		return nil, fmt.Errorf("failed to decode diagnostic event: %w", err)
	}
	return &event, nil
}

func SummarizeXDRObject(data interface{}) string {
	switch v := data.(type) {
	case *xdr.LedgerEntry:
		if v == nil {
			return "empty ledger entry"
		}
		return fmt.Sprintf("LedgerEntry(%v)", v.Data.Type)

	case *xdr.TransactionEnvelope:
		if v == nil {
			return "empty transaction envelope"
		}
		return fmt.Sprintf("TransactionEnvelope(%v)", v.Type)

	case *xdr.DiagnosticEvent:
		if v == nil {
			return "empty diagnostic event"
		}
		return fmt.Sprintf("DiagnosticEvent(successful=%v)", v.InSuccessfulContractCall)

	default:
		return fmt.Sprintf("%T", v)
	}
}
