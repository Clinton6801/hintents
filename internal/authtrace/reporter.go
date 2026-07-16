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

package authtrace

import (
	"encoding/json"
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type DetailedReporter struct {
	trace *AuthTrace
}

var expirationLedgerPattern = regexp.MustCompile(`(?i)expiration(?:[_\s]+ledger)?\D*(\d+)`)

func NewDetailedReporter(trace *AuthTrace) *DetailedReporter {
	return &DetailedReporter{trace: trace}
}

func (r *DetailedReporter) GenerateReport() string {
	var sb strings.Builder

	status := "SUCCEEDED"
	if !r.trace.Success {
		status = "FAILED"
	}

	sb.WriteString("=== MULTI-SIGNATURE AUTHORIZATION DEBUG REPORT ===

")
	fmt.Fprintf(&sb, "Authorization: %s
", status)
	fmt.Fprintf(&sb, "Account: %s
", r.trace.AccountID)
	fmt.Fprintf(&sb, "Total Signers: %d
", r.trace.SignerCount)
	fmt.Fprintf(&sb, "Valid Signatures: %d

", r.trace.ValidSignatures)
	r.writeMultiSigRequirement(&sb)
	if expirationLedger, ok := r.findExpirationLedger(); ok {
		fmt.Fprintf(&sb, "  Expiration Ledger: %d

", expirationLedger)
	}

	if len(r.trace.Failures) > 0 {
		r.writeFailures(&sb)
	}

	if len(r.trace.AuthEvents) > 0 {
		r.writeEvents(&sb)
	}

	if len(r.trace.CustomContracts) > 0 {
		r.writeContracts(&sb)
	}

	r.writeSignatureWeightSummary(&sb)

	return sb.String()
}

func (r *DetailedReporter) findExpirationLedger() (uint32, bool) {
	for _, event := range r.trace.AuthEvents {
		if event.Details == "" {
			continue
		}
		match := expirationLedgerPattern.FindStringSubmatch(event.Details)
		if len(match) < 2 {
			continue
		}
		ledger, err := strconv.ParseUint(match[1], 10, 32)
		if err != nil {
			continue
		}
		if ledger == 0 {
			return 0, false
		}
		return uint32(ledger), true
	}
	return 0, false
}

func (r *DetailedReporter) writeMultiSigRequirement(sb *strings.Builder) {
	requiredWeight, providedWeight, ok := r.multiSigWeights()
	if !ok {
		return
	}

	requiredSigs := minSignaturesForWeight(r.trace.SignatureWeights, requiredWeight)
	if requiredSigs <= 1 {
		return
	}

	providedSigs := r.validSignerCount()
	missingSigs := requiredSigs - providedSigs
	if missingSigs < 0 {
		missingSigs = 0
	}

	fmt.Fprintf(sb, "  Signatures: %d/%d (Missing: %d)
", providedSigs, requiredSigs, missingSigs)
	fmt.Fprintf(sb, "  Required Weight: %d
", requiredWeight)
	fmt.Fprintf(sb, "  Provided Weight: %d

", providedWeight)
}

func (r *DetailedReporter) multiSigWeights() (uint32, uint32, bool) {
	var requiredWeight uint32
	var providedWeight uint32
	if len(r.trace.Failures) > 0 {
		requiredWeight = r.trace.Failures[0].RequiredWeight
		providedWeight = r.trace.Failures[0].CollectedWeight
	} else {
		requiredWeight = r.trace.Thresholds.HighThreshold
		for _, event := range r.trace.AuthEvents {
			if event.EventType == "signature_verification" && event.Status == "valid" {
				providedWeight += event.Weight
			}
		}
	}

	var maxSingleSignerWeight uint32
	for _, signer := range r.trace.SignatureWeights {
		if signer.Weight > maxSingleSignerWeight {
			maxSingleSignerWeight = signer.Weight
		}
	}

	if requiredWeight == 0 || requiredWeight <= maxSingleSignerWeight {
		return 0, 0, false
	}
	return requiredWeight, providedWeight, true
}

func (r *DetailedReporter) validSignerCount() int {
	seen := make(map[string]struct{})
	for _, event := range r.trace.AuthEvents {
		if event.EventType != "signature_verification" || event.Status != "valid" || event.SignerKey == "" {
			continue
		}
		seen[event.SignerKey] = struct{}{}
	}
	return len(seen)
}

func minSignaturesForWeight(weights []KeyWeight, required uint32) int {
	if required == 0 {
		return 0
	}

	sorted := make([]uint32, 0, len(weights))
	for _, w := range weights {
		if w.Weight > 0 {
			sorted = append(sorted, w.Weight)
		}
	}
	sort.Slice(sorted, func(i, j int) bool { return sorted[i] > sorted[j] })

	var total uint32
	for i, weight := range sorted {
		total += weight
		if total >= required {
			return i + 1
		}
	}
	return len(sorted)
}

func (r *DetailedReporter) writeFailures(sb *strings.Builder) {
	sb.WriteString("--- FAILURE DETAILS ---
")
	for i, failure := range r.trace.Failures {
		fmt.Fprintf(sb, "
Failure #%d:
", i+1)
		fmt.Fprintf(sb, "  Reason: %s
", failure.FailureReason)
		fmt.Fprintf(sb, "  Required Weight: %d
", failure.RequiredWeight)
		fmt.Fprintf(sb, "  Collected Weight: %d
", failure.CollectedWeight)
		fmt.Fprintf(sb, "  Missing Weight: %d
", failure.MissingWeight)

		if len(failure.FailedSigners) > 0 {
			sb.WriteString("  Failed Signers:
")
			for _, signer := range failure.FailedSigners {
				fmt.Fprintf(sb, "    - %s (weight: %d, type: %s)
",
					signer.SignerKey, signer.Weight, signer.SignerType)
			}
		}
	}
}

func (r *DetailedReporter) writeEvents(sb *strings.Builder) {
	sb.WriteString("
--- AUTHORIZATION TRACE ---
")
	for i, event := range r.trace.AuthEvents {
		fmt.Fprintf(sb, "
[%d] %s
", i+1, event.EventType)
		if event.SignerKey != "" {
			fmt.Fprintf(sb, "    Signer: %s
", event.SignerKey)
		}
		fmt.Fprintf(sb, "    Status: %s
", event.Status)
		if event.Weight > 0 {
			fmt.Fprintf(sb, "    Weight: %d
", event.Weight)
		}
		if event.Details != "" {
			fmt.Fprintf(sb, "    Details: %s
", event.Details)
		}
		if event.ErrorReason != "" {
			fmt.Fprintf(sb, "    Error: %s
", event.ErrorReason)
		}
	}
}

func (r *DetailedReporter) writeContracts(sb *strings.Builder) {
	sb.WriteString("
--- CUSTOM CONTRACT AUTHORIZATIONS ---
")
	for _, contract := range r.trace.CustomContracts {
		fmt.Fprintf(sb, "
Contract: %s
", contract.ContractID)
		fmt.Fprintf(sb, "  Method: %s
", contract.Method)
		fmt.Fprintf(sb, "  Result: %s
", contract.Result)
		if contract.ErrorMsg != "" {
			fmt.Fprintf(sb, "  Error: %s
", contract.ErrorMsg)
		}
	}
}

func (r *DetailedReporter) writeSignatureWeightSummary(sb *strings.Builder) {
	var totalProvided uint32
	for _, event := range r.trace.AuthEvents {
		if event.EventType == "signature_verification" && event.Status == "valid" {
			totalProvided += event.Weight
		}
	}

	required := r.trace.Thresholds.HighThreshold
	if required == 0 && len(r.trace.Failures) > 0 {
		required = r.trace.Failures[0].RequiredWeight
	}

	fmt.Fprintf(sb, "
Total Signature Weight: %d / Required: %d
", totalProvided, required)
}

func (r *DetailedReporter) GenerateJSON() ([]byte, error) {
	return json.MarshalIndent(r.trace, "", "  ")
}

func (r *DetailedReporter) GenerateJSONString() (string, error) {
	data, err := r.GenerateJSON()
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (r *DetailedReporter) SummaryMetrics() map[string]interface{} {
	metrics := map[string]interface{}{
		"success":          r.trace.Success,
		"account_id":       r.trace.AccountID,
		"total_signers":    r.trace.SignerCount,
		"valid_signatures": r.trace.ValidSignatures,
		"failure_count":    len(r.trace.Failures),
		"event_count":      len(r.trace.AuthEvents),
		"custom_contracts": len(r.trace.CustomContracts),
	}

	if len(r.trace.Failures) > 0 {
		failure := r.trace.Failures[0]
		metrics["failure_reason"] = failure.FailureReason
		metrics["required_weight"] = failure.RequiredWeight
		metrics["collected_weight"] = failure.CollectedWeight
		metrics["missing_weight"] = failure.MissingWeight
	}

	return metrics
}

func (r *DetailedReporter) IdentifyMissingKeys() []SignerInfo {
	if len(r.trace.Failures) == 0 {
		return nil
	}

	failure := r.trace.Failures[0]
	return failure.FailedSigners
}

func (r *DetailedReporter) FindSignatureByKey(key string) *AuthEvent {
	for _, event := range r.trace.AuthEvents {
		if event.SignerKey == key && event.EventType == "signature_verification" {
			return &event
		}
	}
	return nil
}

func (r *DetailedReporter) GetAuthPath(accountID string) []string {
	var path []string
	for _, event := range r.trace.AuthEvents {
		if event.AccountID == accountID {
			path = append(path, fmt.Sprintf("%s(%s)", event.EventType, event.Status))
		}
	}
	return path
}
