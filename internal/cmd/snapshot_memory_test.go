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

package cmd

import (
	"bytes"
	"encoding/base64"
	"io"
	"os"
	"strings"
	"testing"
)

func TestExtractLinearMemoryBase64(t *testing.T) {
	enc := base64.StdEncoding.EncodeToString([]byte("abc"))

	got, err := extractLinearMemoryBase64(`{"linear_memory_base64":"` + enc + `"}`)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != enc {
		t.Fatalf("expected %q, got %q", enc, got)
	}

	got, err = extractLinearMemoryBase64(`{"linear_memory_dump":"` + enc + `"}`)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != enc {
		t.Fatalf("expected dump field %q, got %q", enc, got)
	}

	got, err = extractLinearMemoryBase64(`{"linear_memory":"` + enc + `"}`)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != enc {
		t.Fatalf("expected fallback memory field %q, got %q", enc, got)
	}
}

func TestPrintMemorySegment(t *testing.T) {
	orig := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("pipe error: %v", err)
	}
	os.Stdout = w

	printMemorySegment([]byte("ABCDEFGH"), 32)

	w.Close()
	os.Stdout = orig

	var buf bytes.Buffer
	if _, err := io.Copy(&buf, r); err != nil {
		t.Fatalf("copy error: %v", err)
	}
	out := buf.String()
	if !strings.Contains(out, "00000020") {
		t.Fatalf("expected offset in output, got %q", out)
	}
	if !strings.Contains(out, "|ABCD.EFGH|") {
		t.Fatalf("expected ascii segment in output, got %q", out)
	}
}
