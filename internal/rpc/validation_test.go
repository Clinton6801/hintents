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
	"testing"
)

func TestValidateTransactionHash(t *testing.T) {
	tests := []struct {
		name    string
		hash    string
		wantErr bool
	}{
		{
			name:    "valid lowercase hash",
			hash:    "5c0a1234567890abcdef1234567890abcdef1234567890abcdef1234567890ab",
			wantErr: false,
		},
		{
			name:    "valid uppercase hash",
			hash:    "5C0A1234567890ABCDEF1234567890ABCDEF1234567890ABCDEF1234567890AB",
			wantErr: false,
		},
		{
			name:    "valid mixed case hash",
			hash:    "5c0a1234567890ABCDEF1234567890abcdef1234567890ABCDEF1234567890ab",
			wantErr: false,
		},
		{
			name:    "invalid length - too short",
			hash:    "123",
			wantErr: true,
		},
		{
			name:    "invalid length - too long",
			hash:    "5c0a1234567890abcdef1234567890abcdef1234567890abcdef1234567890ab12",
			wantErr: true,
		},
		{
			name:    "invalid characters",
			hash:    "5c0a1234567890abcdef1234567890abcdef1234567890abcdef1234567890gz", // 'g' and 'z' are invalid
			wantErr: true,
		},
		{
			name:    "empty string",
			hash:    "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateTransactionHash(tt.hash)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateTransactionHash() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
