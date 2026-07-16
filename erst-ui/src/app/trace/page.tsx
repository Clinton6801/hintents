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


"use client";

import React, { useState } from 'react';
import styles from './page.module.css';
import Button from '@/components/ui/Button';

export default function TracePage() {
  const [txHash, setTxHash] = useState('');
  const [loading, setLoading] = useState(false);
  const [result, setResult] = useState<any>(null);
  const [search, setSearch] = useState('');

  const handleTrace = async () => {
    if (!txHash) return;
    
    // Client-side validation for transaction hash (64 hex characters)
    if (txHash.length !== 64 || !/^[0-9a-fA-F]+$/.test(txHash)) {
      setResult({ 
        success: false, 
        error: 'Invalid Transaction Hash', 
        stderr: 'Transaction hash must be exactly 64 hex characters long.' 
      });
      return;
    }

    setLoading(true);
    setResult(null);
    try {
      const res = await fetch('/api/trace', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ txHash, network: 'testnet' })
      });
      const data = await res.json();
      setResult(data);
    } catch (err) {
      setResult({ success: false, error: 'Failed to execute request' });
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className={styles.container}>
      <header className={styles.header}>
        <h1 className="heading-gradient">Interactive Trace Viewer</h1>
        <p className={styles.subtitle}>Explore Soroban execution trees with deep search</p>
      </header>

      <div className={`glass-panel ${styles.viewerPanel}`}>
        <div className={styles.toolbar} style={{ display: 'flex', gap: '1rem', flexWrap: 'wrap' }}>
          <div className={styles.searchBox}>
            <input 
              type="text" 
              placeholder="Tx Hash to trace..." 
              value={txHash}
              onChange={(e) => setTxHash(e.target.value)}
              style={{ width: '300px' }}
            />
          </div>
          <Button variant="primary" onClick={handleTrace} disabled={loading}>
            {loading ? 'Tracing...' : 'Generate Trace'}
          </Button>

          <div className={styles.searchBox} style={{ marginLeft: 'auto' }}>
            <span>🔍</span>
            <input 
              type="text" 
              placeholder="Search trace (e.g. error, CDLZ...)" 
              value={search}
              onChange={(e) => setSearch(e.target.value)}
            />
          </div>
        </div>

        <div className={styles.traceArea} style={{ padding: '1rem', overflowY: 'auto', maxHeight: '600px' }}>
          {!result ? (
            <div className={styles.emptyState}>
              <div className={styles.emptyIcon}>⑆</div>
              <p>Run a simulation in the Debugger to generate a trace tree.</p>
            </div>
          ) : (
            <div style={{ fontSize: '0.85rem' }}>
              {result.success ? (
                <pre style={{ whiteSpace: 'pre-wrap', color: 'var(--text-primary)' }}>
                  {JSON.stringify(result.traceData, null, 2)}
                  {'

Logs:
'}
                  <span style={{ color: 'var(--text-secondary)' }}>{result.logs}</span>
                </pre>
              ) : (
                <>
                  <h3 style={{ color: 'var(--error-color)', marginBottom: '1rem' }}>Trace Failed</h3>
                  <p><strong>Error:</strong> {result.error}</p>
                  <pre style={{ color: '#ff6b6b', marginTop: '1rem', whiteSpace: 'pre-wrap' }}>
                    {result.stderr}
                  </pre>
                </>
              )}
            </div>
          )}
        </div>
      </div>
    </div>
  );
}
