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
	"context"
	"testing"
	"time"
)

func TestAsyncRunner_SubmitAndPoll(t *testing.T) {
	mock := NewMockRunner(func(ctx context.Context, req *SimulationRequest) (*SimulationResponse, error) {
		return &SimulationResponse{
			Status: "success",
			Events: []string{"event1"},
		}, nil
	})

	async := NewAsyncRunner(mock)
	jobID, err := async.Submit(&SimulationRequest{
		EnvelopeXdr:   "AAAA",
		ResultMetaXdr: "BBBB",
	})
	if err != nil {
		t.Fatalf("Submit failed: %v", err)
	}
	if jobID == "" {
		t.Fatal("expected non-empty job ID")
	}

	job, err := async.Wait(context.Background(), jobID, PollConfig{
		Interval: 10 * time.Millisecond,
		Timeout:  5 * time.Second,
	})
	if err != nil {
		t.Fatalf("Wait failed: %v", err)
	}
	if job.Status != JobStatusCompleted {
		t.Errorf("expected completed, got %s", job.Status)
	}
	if job.Response == nil {
		t.Fatal("expected non-nil response")
	}
	if job.Response.Status != "success" {
		t.Errorf("expected success, got %s", job.Response.Status)
	}
}

func TestAsyncRunner_SubmitFailure(t *testing.T) {
	mock := NewMockRunner(func(ctx context.Context, req *SimulationRequest) (*SimulationResponse, error) {
		return nil, &ValidationError{Field: "envelope_xdr", Message: "invalid"}
	})

	async := NewAsyncRunner(mock)
	jobID, err := async.Submit(&SimulationRequest{})
	if err != nil {
		t.Fatalf("Submit failed: %v", err)
	}

	job, err := async.Wait(context.Background(), jobID, PollConfig{
		Interval: 10 * time.Millisecond,
		Timeout:  5 * time.Second,
	})
	if err != nil {
		t.Fatalf("Wait failed: %v", err)
	}
	if job.Status != JobStatusFailed {
		t.Errorf("expected failed, got %s", job.Status)
	}
	if job.Error == "" {
		t.Error("expected non-empty error")
	}
}

func TestAsyncRunner_Timeout(t *testing.T) {
	mock := NewMockRunner(func(ctx context.Context, req *SimulationRequest) (*SimulationResponse, error) {
		time.Sleep(2 * time.Second)
		return &SimulationResponse{Status: "success"}, nil
	})

	async := NewAsyncRunner(mock)
	jobID, err := async.Submit(&SimulationRequest{
		EnvelopeXdr:   "AAAA",
		ResultMetaXdr: "BBBB",
	})
	if err != nil {
		t.Fatalf("Submit failed: %v", err)
	}

	_, err = async.Wait(context.Background(), jobID, PollConfig{
		Interval: 10 * time.Millisecond,
		Timeout:  50 * time.Millisecond,
	})
	if err == nil {
		t.Fatal("expected timeout error")
	}
}

func TestAsyncRunner_Cleanup(t *testing.T) {
	async := NewAsyncRunner(NewDefaultMockRunner())
	jobID, _ := async.Submit(&SimulationRequest{
		EnvelopeXdr:   "AAAA",
		ResultMetaXdr: "BBBB",
	})

	_, _ = async.Wait(context.Background(), jobID, PollConfig{
		Interval: 10 * time.Millisecond,
		Timeout:  5 * time.Second,
	})

	async.Cleanup(jobID)

	_, err := async.Poll(jobID)
	if err == nil {
		t.Error("expected error after cleanup")
	}
}

func TestAsyncRunner_PollNonExistent(t *testing.T) {
	async := NewAsyncRunner(NewDefaultMockRunner())
	_, err := async.Poll("nonexistent")
	if err == nil {
		t.Error("expected error for nonexistent job")
	}
}

func TestAsyncRunner_ContextCancel(t *testing.T) {
	mock := NewMockRunner(func(ctx context.Context, req *SimulationRequest) (*SimulationResponse, error) {
		time.Sleep(10 * time.Second)
		return &SimulationResponse{Status: "success"}, nil
	})

	async := NewAsyncRunner(mock)
	jobID, _ := async.Submit(&SimulationRequest{
		EnvelopeXdr:   "AAAA",
		ResultMetaXdr: "BBBB",
	})

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		time.Sleep(50 * time.Millisecond)
		cancel()
	}()

	_, err := async.Wait(ctx, jobID, PollConfig{
		Interval: 10 * time.Millisecond,
		Timeout:  30 * time.Second,
	})
	if err == nil {
		t.Fatal("expected context cancelled error")
	}
}
