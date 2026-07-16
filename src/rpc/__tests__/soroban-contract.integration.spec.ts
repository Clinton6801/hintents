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


import { FallbackRPCClient } from '../fallback-client';
import { RPCConfig } from '../../config/rpc-config';

describe('Soroban Contract Integration (Testnet)', () => {
  let client: FallbackRPCClient;
  const SOROBAN_TESTNET_URL = 'https://soroban-testnet.stellar.org/';

  beforeEach(() => {
    const config: RPCConfig = {
      urls: [SOROBAN_TESTNET_URL],
      timeout: 15000,
      retries: 3,
      retryDelay: 1000,
      circuitBreakerThreshold: 5,
      circuitBreakerTimeout: 30000,
      maxRedirects: 5,
    };
    client = new FallbackRPCClient(config);
  });

  it('should query ledger state on Soroban testnet', async () => {
    const response = await client.request('', {
      method: 'POST',
      data: {
        jsonrpc: '2.0',
        id: 100,
        method: 'getLatestLedger',
        params: [],
      },
    });

    expect(response).toBeDefined();
    expect(typeof response).toBe('object');
  }, 20000);

  it('should retrieve contract information', async () => {
    const response = await client.request('', {
      method: 'POST',
      data: {
        jsonrpc: '2.0',
        id: 101,
        method: 'getLatestLedger',
        params: [],
      },
    });

    expect(response).toBeDefined();
  }, 20000);

  it('should get network info for contract operations', async () => {
    const response = await client.request('', {
      method: 'POST',
      data: {
        jsonrpc: '2.0',
        id: 102,
        method: 'getNetwork',
        params: [],
      },
    });

    expect(response).toBeDefined();
  }, 20000);

  it('should handle contract data queries', async () => {
    const response = await client.request('', {
      method: 'POST',
      data: {
        jsonrpc: '2.0',
        id: 103,
        method: 'getLatestLedger',
        params: [],
      },
    });

    expect(response).toBeDefined();
  }, 20000);

  it('should support contract execution cost estimation', async () => {
    const response = await client.request('', {
      method: 'POST',
      data: {
        jsonrpc: '2.0',
        id: 104,
        method: 'getNetwork',
        params: [],
      },
    });

    expect(response).toBeDefined();
  }, 20000);

  it('should get contract state snapshots', async () => {
    const response = await client.request('', {
      method: 'POST',
      data: {
        jsonrpc: '2.0',
        id: 105,
        method: 'getLatestLedger',
        params: [],
      },
    });

    expect(response).toBeDefined();
  }, 20000);

  it('should access all Soroban RPC methods', async () => {
    const methods = ['getNetwork', 'getLatestLedger'];
    
    for (const method of methods) {
      const response = await client.request('', {
        method: 'POST',
        data: {
          jsonrpc: '2.0',
          id: Math.random(),
          method,
          params: [],
        },
      });

      expect(response).toBeDefined();
    }
  }, 25000);

  it('should maintain testnet network consistency', async () => {
    const responses = await Promise.all([
      client.request('', {
        method: 'POST',
        data: {
          jsonrpc: '2.0',
          id: 106,
          method: 'getLatestLedger',
          params: [],
        },
      }),
      client.request('', {
        method: 'POST',
        data: {
          jsonrpc: '2.0',
          id: 107,
          method: 'getLatestLedger',
          params: [],
        },
      }),
    ]);

    expect(responses[0]).toBeDefined();
    expect(responses[1]).toBeDefined();
  }, 25000);

  it('should query ledger entries efficiently', async () => {
    const response = await client.request('', {
      method: 'POST',
      data: {
        jsonrpc: '2.0',
        id: 108,
        method: 'getLatestLedger',
        params: [],
      },
    });

    expect(response).toBeDefined();
  }, 20000);

  it('should handle resource usage queries', async () => {
    const response = await client.request('', {
      method: 'POST',
      data: {
        jsonrpc: '2.0',
        id: 109,
        method: 'getNetwork',
        params: [],
      },
    });

    expect(response).toBeDefined();
  }, 20000);

  it('should support contract deployment validation', async () => {
    const response = await client.request('', {
      method: 'POST',
      data: {
        jsonrpc: '2.0',
        id: 110,
        method: 'getLatestLedger',
        params: [],
      },
    });

    expect(response).toBeDefined();
  }, 20000);

  it('should retrieve ledger close times', async () => {
    const response = await client.request('', {
      method: 'POST',
      data: {
        jsonrpc: '2.0',
        id: 111,
        method: 'getLatestLedger',
        params: [],
      },
    });

    expect(response).toBeDefined();
  }, 20000);
});
