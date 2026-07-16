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



import test from 'node:test';
import assert from 'node:assert/strict';
import { buildTraceTreeExport, renderStandaloneHtml } from './traceExport';
import { Trace } from './erstClient';

test('buildTraceTreeExport includes matches and expanded nodes', () => {
    const trace: Trace = {
        transaction_hash: 'tx-123',
        start_time: '2026-02-23T10:00:00Z',
        states: [
            {
                step: 1,
                timestamp: '2026-02-23T10:00:01Z',
                operation: 'invoke',
                contract_id: 'CABC123',
                function: 'transfer',
                arguments: ['alice', 'bob', 5],
                return_value: { ok: true }
            },
            {
                step: 2,
                timestamp: '2026-02-23T10:00:02Z',
                operation: 'event',
                error: 'insufficient balance'
            }
        ]
    };

    const payload = buildTraceTreeExport(trace, 'transfer');

    assert.equal(payload.transactionHash, 'tx-123');
    assert.equal(payload.searchQuery, 'transfer');
    assert.ok(payload.totalMatches > 0);
    assert.equal(payload.tree.length, 2);
    assert.ok(payload.tree[0].children.length >= 5);
});

test('renderStandaloneHtml renders metadata and payload json', () => {
    const trace: Trace = {
        transaction_hash: 'tx-html',
        start_time: '2026-02-23T10:00:00Z',
        states: [
            {
                step: 1,
                timestamp: '2026-02-23T10:00:01Z',
                operation: 'invoke'
            }
        ]
    };

    const payload = buildTraceTreeExport(trace, '');
    const html = renderStandaloneHtml(payload);

    assert.ok(html.includes('<!doctype html>'));
    assert.ok(html.includes('ERST Trace Tree Export'));
    assert.ok(html.includes('tx-html'));
    assert.ok(html.includes('&quot;transactionHash&quot;: &quot;tx-html&quot;'));
});
