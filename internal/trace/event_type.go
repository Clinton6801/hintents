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

package trace

import "strings"

// Event type filter constants for trace viewer filtering
const (
	EventTypeTrap         = "trap"
	EventTypeContractCall = "contract_call"
	EventTypeHostFunction = "host_function"
	EventTypeAuth         = "auth"
	EventTypeOther        = "other"
)

// AllFilterableEventTypes returns the list of event types that can be used for filtering
func AllFilterableEventTypes() []string {
	return []string{EventTypeTrap, EventTypeContractCall, EventTypeHostFunction, EventTypeAuth}
}

// ClassifyEventType returns the event type for a given execution state.
// If state.EventType is set, it is used; otherwise the type is inferred from
// Operation, Function, and Error fields.
func ClassifyEventType(state *ExecutionState) string {
	if state == nil {
		return EventTypeOther
	}
	if state.EventType != "" {
		return normalizeEventType(state.EventType)
	}
	op := strings.ToLower(strings.TrimSpace(state.Operation))
	fn := strings.ToLower(strings.TrimSpace(state.Function))
	err := strings.ToLower(strings.TrimSpace(state.Error))

	if state.Error != "" && (strings.Contains(err, "trap") || strings.Contains(err, "panic")) {
		return EventTypeTrap
	}
	if strings.Contains(op, "trap") {
		return EventTypeTrap
	}

	if strings.Contains(fn, "require_auth") || strings.Contains(fn, "authorize") ||
		strings.Contains(op, "auth") || strings.Contains(op, "require_auth") {
		return EventTypeAuth
	}

	if strings.Contains(op, "host") || strings.Contains(op, "host_fn") ||
		strings.Contains(op, "host function") || isKnownHostFunction(fn) {
		return EventTypeHostFunction
	}

	if strings.Contains(op, "contract_call") || strings.Contains(op, "contract_init") ||
		(strings.Contains(op, "invoke") && !strings.Contains(op, "host")) {
		return EventTypeContractCall
	}
	if state.ContractID != "" && state.Function != "" && state.Error == "" {
		return EventTypeContractCall
	}

	return EventTypeOther
}

func normalizeEventType(s string) string {
	switch strings.ToLower(strings.TrimSpace(s)) {
	case EventTypeTrap, "traps":
		return EventTypeTrap
	case EventTypeContractCall, "contract call", "contractcall":
		return EventTypeContractCall
	case EventTypeHostFunction, "host function", "host_fn", "hostfn":
		return EventTypeHostFunction
	case EventTypeAuth, "auth_event", "auth event":
		return EventTypeAuth
	default:
		return EventTypeOther
	}
}

func isKnownHostFunction(fn string) bool {
	hostFns := []string{
		"require_auth", "put_ledger_entry", "get_ledger_entry", "del_ledger_entry",
		"get_contract_instance", "get_contract_code", "get_ledger_key_durability",
		"invoke_contract", "create_contract", "deployer", "builtin",
	}
	for _, h := range hostFns {
		if fn == h || strings.HasPrefix(fn, h+" ") {
			return true
		}
	}
	return false
}

// StepMatchesFilter returns true if the state at the given step index matches the filter.
// Empty filter means no filtering (all steps match).
func (t *ExecutionTrace) StepMatchesFilter(stepIndex int, filter string) bool {
	if stepIndex < 0 || stepIndex >= len(t.States) {
		return false
	}
	if filter == "" {
		return true
	}
	return ClassifyEventType(&t.States[stepIndex]) == filter
}
