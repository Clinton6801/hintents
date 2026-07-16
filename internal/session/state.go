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

package session

import (
	"sync"
)

// State represents the current session data
type State map[string]interface{}

// Action defines a state change request
type Action struct {
	Type    string
	Payload interface{}
}

// Dispatcher is the function type that processes an Action
type Dispatcher func(action Action)

// Middleware wraps a Dispatcher to allow custom logic injection
type Middleware func(next Dispatcher) Dispatcher

// StateStore manages the session state with injectable middleware
type StateStore struct {
	mu         sync.RWMutex
	state      State
	dispatch   Dispatcher
	middleware []Middleware
}

// NewStateStore initializes the store [Issue #589]
func NewStateStore() *StateStore {
	s := &StateStore{
		state: make(State),
	}
	// The base dispatcher updates the actual state map
	s.dispatch = s.baseDispatch
	return s
}

// Use injects custom middleware into the state management pipeline
func (s *StateStore) Use(mw Middleware) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.middleware = append(s.middleware, mw)

	// Re-chain the middleware (the "Optimize" operation)
	// We wrap the base dispatch with each middleware in reverse order
	composed := s.baseDispatch
	for i := len(s.middleware) - 1; i >= 0; i-- {
		composed = s.middleware[i](composed)
	}
	s.dispatch = composed
}

func (s *StateStore) baseDispatch(action Action) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.state[action.Type] = action.Payload
}

// Dispatch triggers a state change through the middleware chain
func (s *StateStore) Dispatch(action Action) {
	s.dispatch(action)
}

// Get safely retrieves session data
func (s *StateStore) Get(key string) (interface{}, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	val, ok := s.state[key]
	return val, ok
}
