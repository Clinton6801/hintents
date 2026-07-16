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



// Example usage of generated TypeScript bindings
import * as StellarSdk from '@stellar/stellar-sdk';

// This would be imported from your generated bindings
// import { MyContractClient } from './generated';

interface ClientConfig {
  contractId: string;
  network: 'testnet' | 'mainnet' | 'futurenet';
  enableSimulation?: boolean;
}

async function main() {
  // Initialize client with contract details
  const config: ClientConfig = {
    contractId: 'CDLZFC3SYJYDZT7K67VZ75HPJVIEUVNIXF47ZG2FB2RMQQAHHAGCN4B2',
    network: 'testnet',
    enableSimulation: true,
  };

  // Create source keypair (in production, load from secure storage)
  const sourceKeypair = StellarSdk.Keypair.fromSecret(
    'SXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX'
  );

  console.log('Source account:', sourceKeypair.publicKey());

  // Example 1: Simulate before executing
  console.log('
=== Example 1: Simulation ===');
  // const simResult = await client.transfer(
  //   sourceKeypair,
  //   'GDQP2KPQGKIHYJGXNUIYOMHARUARCA7DJT5FO2FFOOKY3B2WSQHG4W37',
  //   BigInt(1000000),
  //   { simulate: true }
  // );
  // console.log('Simulation status:', simResult.simulation?.status);
  // console.log('CPU usage:', simResult.simulation?.budgetUsage?.cpuUsagePercent);

  // Example 2: Execute transaction
  console.log('
=== Example 2: Execute Transaction ===');
  // const result = await client.transfer(
  //   sourceKeypair,
  //   'GDQP2KPQGKIHYJGXNUIYOMHARUARCA7DJT5FO2FFOOKY3B2WSQHG4W37',
  //   BigInt(1000000)
  // );
  // console.log('Transaction hash:', result.transactionHash);

  // Example 3: Debug failed transaction
  console.log('
=== Example 3: Debug Failed Transaction ===');
}

  // const simulator = new ErstSimulator({
  //   network: 'testnet',
  //   rpcUrl: 'https://soroban-testnet.stellar.org',
  // });
  // const debugResult = await simulator.debugTransaction(
  //   '5c0a1234567890abcdef1234567890abcdef1234567890abcdef1234567890ab'
  // );
  // console.log('Error:', debugResult.error);
  // console.log('Events:', debugResult.events);
}

main().catch(console.error);
