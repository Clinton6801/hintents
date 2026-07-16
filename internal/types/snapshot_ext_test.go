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

package types_test

import (
	"testing"

	"github.com/dotandev/hintents/internal/snapshot"
	"github.com/dotandev/hintents/internal/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFingerprint_NilSnapshot(t *testing.T) {
	fp, err := types.Fingerprint(nil)
	require.NoError(t, err)
	assert.Equal(t, uint64(0), fp)
}

func TestFingerprint_Deterministic(t *testing.T) {
	entries := map[string]string{
		"key_b": "val_b",
		"key_a": "val_a",
	}
	snap := snapshot.FromMap(entries)

	fp1, err := types.Fingerprint(snap)
	require.NoError(t, err)

	fp2, err := types.Fingerprint(snap)
	require.NoError(t, err)

	assert.Equal(t, fp1, fp2, "identical snapshots must produce identical fingerprints")
}

func TestFingerprint_ChangedEntryProducesDifferentHash(t *testing.T) {
	snap1 := snapshot.FromMap(map[string]string{"key_a": "val_a"})
	snap2 := snapshot.FromMap(map[string]string{"key_a": "val_b"})

	fp1, err := types.Fingerprint(snap1)
	require.NoError(t, err)

	fp2, err := types.Fingerprint(snap2)
	require.NoError(t, err)

	assert.NotEqual(t, fp1, fp2, "different ledger values must produce different fingerprints")
}

func TestFingerprint_LinearMemoryAffectsHash(t *testing.T) {
	entries := map[string]string{"k": "v"}

	snapWithoutMem := snapshot.FromMapWithOptions(entries, snapshot.BuildOptions{})
	snapWithMem := snapshot.FromMapWithOptions(entries, snapshot.BuildOptions{LinearMemory: []byte("memory data")})

	fp1, err := types.Fingerprint(snapWithoutMem)
	require.NoError(t, err)

	fp2, err := types.Fingerprint(snapWithMem)
	require.NoError(t, err)

	assert.NotEqual(t, fp1, fp2, "presence of linear memory must change the fingerprint")
}

func TestFingerprint_EmptySnapshot(t *testing.T) {
	snap := snapshot.FromMap(nil)

	fp, err := types.Fingerprint(snap)
	require.NoError(t, err)
	assert.NotEqual(t, uint64(0), fp, "empty but non-nil snapshot should produce a non-zero fingerprint")
}
