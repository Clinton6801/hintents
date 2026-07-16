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

package db

import (
	"database/sql"
	"fmt"
	"log/slog"
	"strings"
)

// DestructiveOp represents a type of destructive SQL operation.
type DestructiveOp string

const (
	OpDelete   DestructiveOp = "DELETE"
	OpDrop     DestructiveOp = "DROP"
	OpAlter    DestructiveOp = "ALTER"
	OpTruncate DestructiveOp = "TRUNCATE"
	OpUpdate   DestructiveOp = "UPDATE"
	OpSafe     DestructiveOp = ""
)

// DryRunResult holds the outcome of a dry-run analysis.
type DryRunResult struct {
	Query       string
	Args        []interface{}
	Operation   DestructiveOp
	Destructive bool
}

// ClassifySQL returns the destructive operation type for a SQL statement.
// Returns OpSafe if the statement is not destructive.
func ClassifySQL(query string) DestructiveOp {
	normalized := strings.ToUpper(strings.TrimSpace(query))
	switch {
	case strings.HasPrefix(normalized, "DELETE"):
		return OpDelete
	case strings.HasPrefix(normalized, "DROP"):
		return OpDrop
	case strings.HasPrefix(normalized, "ALTER"):
		return OpAlter
	case strings.HasPrefix(normalized, "TRUNCATE"):
		return OpTruncate
	case strings.HasPrefix(normalized, "UPDATE"):
		return OpUpdate
	default:
		return OpSafe
	}
}

// DryRunExec analyzes a SQL statement and logs a warning if it is destructive,
// without executing it. Returns a DryRunResult describing what would happen.
func DryRunExec(logger *slog.Logger, _ *sql.DB, query string, args ...interface{}) DryRunResult {
	op := ClassifySQL(query)
	result := DryRunResult{
		Query:       query,
		Args:        args,
		Operation:   op,
		Destructive: op != OpSafe,
	}

	if result.Destructive {
		logger.Warn("[DRY-RUN] destructive SQL detected",
			"operation", string(op),
			"query", query,
			"args", fmt.Sprintf("%v", args),
		)
	}

	return result
}
