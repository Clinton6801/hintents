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


import { XDRDecoder } from '../../xdr/decoder';

describe('Streaming XDR decoder memory usage', () => {
    it('should stream decode 1000 ledger entries with low peak memory', async () => {
        // Provide a mock decode function for the streaming decoder
            const mockDecodeFn = () => ({ dummy: true });
            // Generate 1000 fake base64-encoded LedgerEntry XDRs (simulate real entries)
            const fakeEntry = Buffer.from('AAAAAA==', 'base64').toString('base64'); // Minimal valid base64
            const entries = Array(1000).fill(fakeEntry).join('
');

            const input = Buffer.from(entries, 'utf8');
            const initialMemory = process.memoryUsage().heapUsed;
            const stream = XDRDecoder.streamLedgerEntries(input, mockDecodeFn);
            let count = 0;
            for await (const _ of stream) {
                count++;
            }
            const finalMemory = process.memoryUsage().heapUsed;
            const memoryIncrease = (finalMemory - initialMemory) / 1024 / 1024;
            console.log(`Streamed ${count} entries, memory increase: ${memoryIncrease.toFixed(2)}MB`);
            expect(count).toBe(1000);
            expect(memoryIncrease).toBeLessThan(50); // Should not exceed 50MB
        });
});
