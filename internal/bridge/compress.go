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

// Package bridge provides IPC compression helpers for ledger snapshot payloads.
// Ledger entries contain highly repetitive XDR bytes that compress extremely well
// with Zstd, reducing IPC payload size by 60-80% in practice.
package bridge

import (
	"bytes"
	"encoding/json"
	"fmt"
	"sync"

	"github.com/klauspost/compress/zstd"
)

var (
	encoder     *zstd.Encoder
	encoderErr  error
	decoder     *zstd.Decoder
	decoderErr  error
	encoderOnce sync.Once
	decoderOnce sync.Once
)

func getEncoder() (*zstd.Encoder, error) {
	encoderOnce.Do(func() {
		encoder, encoderErr = zstd.NewWriter(nil, zstd.WithEncoderLevel(zstd.SpeedDefault))
	})
	return encoder, encoderErr
}

func getDecoder() (*zstd.Decoder, error) {
	decoderOnce.Do(func() {
		decoder, decoderErr = zstd.NewReader(nil)
	})
	return decoder, decoderErr
}

// CompressLedgerEntries serialises entries to JSON and compresses with Zstd.
// Returns the raw compressed bytes.
func CompressLedgerEntries(entries map[string]string) ([]byte, error) {
	raw, err := json.Marshal(entries)
	if err != nil {
		return nil, fmt.Errorf("bridge: marshal ledger entries: %w", err)
	}
	enc, err := getEncoder()
	if err != nil {
		return nil, fmt.Errorf("bridge: init zstd encoder: %w", err)
	}
	return enc.EncodeAll(raw, make([]byte, 0, len(raw)/4)), nil
}

// DecompressLedgerEntries decompresses a Zstd blob produced by CompressLedgerEntries.
func DecompressLedgerEntries(compressed []byte) (map[string]string, error) {
	dec, err := getDecoder()
	if err != nil {
		return nil, fmt.Errorf("bridge: init zstd decoder: %w", err)
	}
	raw, err := dec.DecodeAll(compressed, make([]byte, 0, len(compressed)*4))
	if err != nil {
		return nil, fmt.Errorf("bridge: decompress ledger entries: %w", err)
	}
	var entries map[string]string
	if err := json.NewDecoder(bytes.NewReader(raw)).Decode(&entries); err != nil {
		return nil, fmt.Errorf("bridge: unmarshal ledger entries: %w", err)
	}
	return entries, nil
}
