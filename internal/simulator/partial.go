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

package simulator

import (
	"github.com/dotandev/hintents/internal/errors"
)

// PartialSimResult holds the outcome of a partial simulation run.
// When Completed is false, HaltedAtKey identifies the first missing ledger key
// that caused the simulation to stop. OpsApplied counts how many ledger entries
// were successfully validated before the halt.
type PartialSimResult struct {
	Completed   bool
	HaltedAtKey string // empty when Completed is true
	OpsApplied  int
}

// SimulatePartial runs a simulation request against the provided ledger state,
// halting gracefully at the first missing footprint key rather than propagating
// a hard error. It returns a PartialSimResult describing how far execution
// reached, paired with a *errors.MissingLedgerKeyError when halted early.
//
// If all keys required by the request are present in state, the result will
// have Completed set to true and a nil error.
func SimulatePartial(req *SimulationRequest, state map[string]string) (*PartialSimResult, error) {
	result := &PartialSimResult{}

	if req == nil {
		return result, errors.WrapValidationError("simulation request must not be nil")
	}

	for key := range req.LedgerEntries {
		if _, ok := state[key]; !ok {
			result.HaltedAtKey = key
			result.Completed = false
			return result, errors.WrapMissingLedgerKey(key)
		}
		result.OpsApplied++
	}

	result.Completed = true
	return result, nil
}
