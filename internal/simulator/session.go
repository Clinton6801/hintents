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
	"database/sql"
	"os"
	"path/filepath"
	"time"

	_ "modernc.org/sqlite"
)

type Session struct {
	ID        int64     `json:"id"`
	TxHash    string    `json:"tx_hash"`
	Network   string    `json:"network"`
	Timestamp time.Time `json:"timestamp"`
	Error     string    `json:"error,omitempty"`
	Events    string    `json:"events,omitempty"`
	Logs      string    `json:"logs,omitempty"`
}

type DB struct {
	conn *sql.DB
}

type SearchFilters struct {
	Error    string
	Event    string
	Contract string
	UseRegex bool
}

func OpenDB() (*DB, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	dbPath := filepath.Join(home, ".erst", "sessions.db")

	if err := os.MkdirAll(filepath.Dir(dbPath), 0755); err != nil {
		return nil, err
	}

	conn, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return nil, err
	}

	db := &DB{conn: conn}
	if err := db.init(); err != nil {
		return nil, err
	}

	return db, nil
}

func (db *DB) init() error {
	query := `
	CREATE TABLE IF NOT EXISTS sessions (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		tx_hash TEXT NOT NULL,
		network TEXT NOT NULL,
		timestamp DATETIME NOT NULL,
		error TEXT,
		events TEXT,
		logs TEXT
	);
	CREATE INDEX IF NOT EXISTS idx_tx_hash ON sessions(tx_hash);
	CREATE INDEX IF NOT EXISTS idx_error ON sessions(error);
	`
	_, err := db.conn.Exec(query)
	return err
}

func (db *DB) SaveSession(s *Session) error {
	query := "INSERT INTO sessions (tx_hash, network, timestamp, error, events, logs) VALUES (?, ?, ?, ?, ?, ?)"
	_, err := db.conn.Exec(query, s.TxHash, s.Network, s.Timestamp, s.Error, s.Events, s.Logs)
	return err
}

func (db *DB) SearchSessions(filters SearchFilters) ([]Session, error) {
	query := "SELECT id, tx_hash, network, timestamp, error, events, logs FROM sessions WHERE 1=1"
	var args []interface{}

	if filters.Error != "" {
		if filters.UseRegex {
			query += " AND error REGEXP ?"
		} else {
			query += " AND error LIKE ?"
			filters.Error = "%" + filters.Error + "%"
		}
		args = append(args, filters.Error)
	}

	if filters.Event != "" {
		if filters.UseRegex {
			query += " AND events REGEXP ?"
		} else {
			query += " AND events LIKE ?"
			filters.Event = "%" + filters.Event + "%"
		}
		args = append(args, filters.Event)
	}

	if filters.Contract != "" {
		if filters.UseRegex {
			query += " AND (events REGEXP ? OR logs REGEXP ?)"
			args = append(args, filters.Contract, filters.Contract)
		} else {
			query += " AND (events LIKE ? OR logs LIKE ?)"
			match := "%" + filters.Contract + "%"
			args = append(args, match, match)
		}
	}

	query += " ORDER BY timestamp DESC"

	rows, err := db.conn.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var sessions []Session
	for rows.Next() {
		var s Session
		err := rows.Scan(&s.ID, &s.TxHash, &s.Network, &s.Timestamp, &s.Error, &s.Events, &s.Logs)
		if err != nil {
			return nil, err
		}
		sessions = append(sessions, s)
	}

	return sessions, nil
}
