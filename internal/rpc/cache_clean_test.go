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
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	_ "modernc.org/sqlite"
)

func setupCleanTestDB(t *testing.T) {
	t.Helper()
	db, err := sql.Open("sqlite", ":memory:")
	require.NoError(t, err)
	require.NoError(t, InitCacheWithDB(db))
	t.Cleanup(func() { _ = CloseCache() })
}

func TestCleanByFilter_NoFilter(t *testing.T) {
	setupCleanTestDB(t)
	_, err := CleanByFilter(CleanFilter{})
	require.Error(t, err)
}

func TestCleanByFilter_All(t *testing.T) {
	setupCleanTestDB(t)

	require.NoError(t, SetWithTTLAndNetwork("k1", "v1", time.Hour, "mainnet"))
	require.NoError(t, SetWithTTLAndNetwork("k2", "v2", time.Hour, "testnet"))

	removed, err := CleanByFilter(CleanFilter{All: true})
	require.NoError(t, err)
	require.Equal(t, 2, removed)
}

func TestCleanByFilter_ByNetwork(t *testing.T) {
	setupCleanTestDB(t)

	require.NoError(t, SetWithTTLAndNetwork("k1", "v1", time.Hour, "mainnet"))
	require.NoError(t, SetWithTTLAndNetwork("k2", "v2", time.Hour, "testnet"))

	removed, err := CleanByFilter(CleanFilter{Network: "testnet"})
	require.NoError(t, err)
	require.Equal(t, 1, removed)

	_, found, err := Get("k1")
	require.NoError(t, err)
	require.True(t, found)
}

func TestCleanByFilter_ByAge(t *testing.T) {
	setupCleanTestDB(t)

	require.NoError(t, SetWithTTLAndNetwork("old", "v", time.Hour, "mainnet"))

	db, err := ensureDB()
	require.NoError(t, err)
	oldTime := time.Now().Add(-10 * 24 * time.Hour).UnixNano()
	_, err = db.Exec("UPDATE rpc_cache SET created_at = ? WHERE cache_key = 'old'", oldTime)
	require.NoError(t, err)

	require.NoError(t, SetWithTTLAndNetwork("fresh", "v", time.Hour, "mainnet"))

	removed, err := CleanByFilter(CleanFilter{OlderThan: 7 * 24 * time.Hour})
	require.NoError(t, err)
	require.Equal(t, 1, removed)

	_, found, err := Get("fresh")
	require.NoError(t, err)
	require.True(t, found)
}

func TestCleanByFilter_ByNetworkAndAge(t *testing.T) {
	setupCleanTestDB(t)

	require.NoError(t, SetWithTTLAndNetwork("old-main", "v", time.Hour, "mainnet"))
	require.NoError(t, SetWithTTLAndNetwork("old-test", "v", time.Hour, "testnet"))

	db, err := ensureDB()
	require.NoError(t, err)
	oldTime := time.Now().Add(-10 * 24 * time.Hour).UnixNano()
	_, err = db.Exec("UPDATE rpc_cache SET created_at = ? WHERE cache_key IN ('old-main','old-test')", oldTime)
	require.NoError(t, err)

	removed, err := CleanByFilter(CleanFilter{
		OlderThan: 7 * 24 * time.Hour,
		Network:   "mainnet",
	})
	require.NoError(t, err)
	require.Equal(t, 1, removed)

	_, found, err := Get("old-test")
	require.NoError(t, err)
	require.True(t, found)
}
