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
	"context"
	"testing"

	hProtocol "github.com/stellar/go-stellar-sdk/protocols/horizon"
	"github.com/stretchr/testify/assert"
)

const probeTestHash = "5c0a1234567890abcdef1234567890abcdef1234567890abcdef1234567890ab"

func txSuccessRoute(hash string) MockRoute {
	return SuccessRoute(hProtocol.Transaction{Hash: hash})
}

func TestResolveNetwork_FoundOnMainnet(t *testing.T) {
	mainnet := NewMockServer(map[string]MockRoute{
		"/transactions/" + probeTestHash: txSuccessRoute(probeTestHash),
	})
	defer mainnet.Close()

	empty := NewMockServer(map[string]MockRoute{})
	defer empty.Close()

	overrides := map[Network]string{
		Mainnet:   mainnet.URL(),
		Testnet:   empty.URL(),
		Futurenet: empty.URL(),
	}

	net, err := resolveNetwork(context.Background(), probeTestHash, "", overrides)
	assert.NoError(t, err)
	assert.Equal(t, Mainnet, net)
}

func TestResolveNetwork_FoundOnTestnet(t *testing.T) {
	testnet := NewMockServer(map[string]MockRoute{
		"/transactions/" + probeTestHash: txSuccessRoute(probeTestHash),
	})
	defer testnet.Close()

	empty := NewMockServer(map[string]MockRoute{})
	defer empty.Close()

	overrides := map[Network]string{
		Mainnet:   empty.URL(),
		Testnet:   testnet.URL(),
		Futurenet: empty.URL(),
	}

	net, err := resolveNetwork(context.Background(), probeTestHash, "", overrides)
	assert.NoError(t, err)
	assert.Equal(t, Testnet, net)
}

func TestResolveNetwork_FoundOnFuturenet(t *testing.T) {
	futurenet := NewMockServer(map[string]MockRoute{
		"/transactions/" + probeTestHash: txSuccessRoute(probeTestHash),
	})
	defer futurenet.Close()

	empty := NewMockServer(map[string]MockRoute{})
	defer empty.Close()

	overrides := map[Network]string{
		Mainnet:   empty.URL(),
		Testnet:   empty.URL(),
		Futurenet: futurenet.URL(),
	}

	net, err := resolveNetwork(context.Background(), probeTestHash, "", overrides)
	assert.NoError(t, err)
	assert.Equal(t, Futurenet, net)
}

func TestResolveNetwork_NotFound(t *testing.T) {
	empty := NewMockServer(map[string]MockRoute{})
	defer empty.Close()

	overrides := map[Network]string{
		Mainnet:   empty.URL(),
		Testnet:   empty.URL(),
		Futurenet: empty.URL(),
	}

	_, err := resolveNetwork(context.Background(), probeTestHash, "", overrides)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "not found")
}

func TestResolveNetwork_ContextAlreadyCancelled(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	empty := NewMockServer(map[string]MockRoute{})
	defer empty.Close()

	overrides := map[Network]string{
		Mainnet:   empty.URL(),
		Testnet:   empty.URL(),
		Futurenet: empty.URL(),
	}

	_, err := resolveNetwork(ctx, probeTestHash, "", overrides)
	assert.Error(t, err)
}

func TestResolveNetwork_NotFoundErrorMentionsAllNetworks(t *testing.T) {
	empty := NewMockServer(map[string]MockRoute{})
	defer empty.Close()

	overrides := map[Network]string{
		Mainnet:   empty.URL(),
		Testnet:   empty.URL(),
		Futurenet: empty.URL(),
	}

	_, err := resolveNetwork(context.Background(), probeTestHash, "", overrides)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), string(Mainnet))
	assert.Contains(t, err.Error(), string(Testnet))
	assert.Contains(t, err.Error(), string(Futurenet))
}
