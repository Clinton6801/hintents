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
	"testing"

	interrors "github.com/dotandev/hintents/internal/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSimulatePartial_HaltsOnMissingKey(t *testing.T) {
	req := &SimulationRequest{
		EnvelopeXdr:   "envelope-xdr",
		ResultMetaXdr: "result-meta-xdr",
		LedgerEntries: map[string]string{
			"key-a": "value-a",
			"key-b": "value-b",
		},
	}

	state := map[string]string{
		"key-a": "value-a",
		// "key-b" intentionally absent
	}

	result, err := SimulatePartial(req, state)

	require.Error(t, err)
	var missingErr *interrors.MissingLedgerKeyError
	require.ErrorAs(t, err, &missingErr)
	assert.Equal(t, "key-b", missingErr.Key)
	assert.False(t, result.Completed)
	assert.Equal(t, "key-b", result.HaltedAtKey)
	assert.True(t, interrors.Is(err, interrors.ErrMissingLedgerKey))
}

func TestSimulatePartial_CompletesWhenAllKeysPresent(t *testing.T) {
	req := &SimulationRequest{
		EnvelopeXdr:   "envelope-xdr",
		ResultMetaXdr: "result-meta-xdr",
		LedgerEntries: map[string]string{
			"key-a": "value-a",
			"key-b": "value-b",
		},
	}

	state := map[string]string{
		"key-a": "value-a",
		"key-b": "value-b",
	}

	result, err := SimulatePartial(req, state)

	require.NoError(t, err)
	assert.True(t, result.Completed)
	assert.Empty(t, result.HaltedAtKey)
	assert.Equal(t, 2, result.OpsApplied)
}

func TestSimulatePartial_EmptyLedgerEntriesCompletes(t *testing.T) {
	req := &SimulationRequest{
		EnvelopeXdr:   "envelope-xdr",
		ResultMetaXdr: "result-meta-xdr",
		LedgerEntries: map[string]string{},
	}

	result, err := SimulatePartial(req, map[string]string{})

	require.NoError(t, err)
	assert.True(t, result.Completed)
	assert.Equal(t, 0, result.OpsApplied)
}

func TestSimulatePartial_NilRequestReturnsValidationError(t *testing.T) {
	result, err := SimulatePartial(nil, map[string]string{})

	require.Error(t, err)
	assert.False(t, result.Completed)
	assert.True(t, interrors.Is(err, interrors.ErrValidationFailed))
}

func TestSimulatePartial_AllKeysMissingHaltsAtFirst(t *testing.T) {
	req := &SimulationRequest{
		EnvelopeXdr:   "envelope-xdr",
		ResultMetaXdr: "result-meta-xdr",
		LedgerEntries: map[string]string{
			"key-a": "value-a",
		},
	}

	result, err := SimulatePartial(req, map[string]string{})

	require.Error(t, err)
	var missingErr *interrors.MissingLedgerKeyError
	require.ErrorAs(t, err, &missingErr)
	assert.False(t, result.Completed)
	assert.Equal(t, 0, result.OpsApplied)
}
