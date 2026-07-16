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

export default function AuditPage() {
  const [payload, setPayload] = useState('');
  const [loading, setLoading] = useState(false);
  const [result, setResult] = useState<any>(null);

  const handleSign = async () => {
    if (!payload) return;
    setLoading(true);
    setResult(null);
    try {
      const res = await fetch('/api/audit', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ payload, privateKey: 'mock-key' })
      });
      const data = await res.json();
      setResult(data);
    } catch (err) {
      setResult({ error: 'Failed to execute request' });
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className={styles.container}>
      <header className={styles.header}>
        <h1 className="heading-gradient">Audit Log Signing</h1>
        <p className={styles.subtitle}>Generate deterministic, signed audit logs for transactions</p>
      </header>

      <div className={styles.grid}>
        <div className={`glass-card ${styles.configPanel}`}>
          <h3>Configuration</h3>
          
          <div className={styles.inputGroup}>
            <label>Mode</label>
            <select className={styles.input}>
              <option>Software (Ed25519)</option>
              <option>HSM (YubiKey/PKCS#11)</option>
              <option>AWS KMS</option>
            </select>
          </div>

          <div className={styles.inputGroup}>
            <label>Payload (JSON)</label>
            <textarea 
              className={`${styles.input} ${styles.textarea}`} 
              placeholder='{"event": "test"}'
              rows={6}
              value={payload}
              onChange={(e) => setPayload(e.target.value)}
            />
          </div>
          
          <Button variant="primary" onClick={handleSign} disabled={loading}>
            {loading ? 'Signing...' : 'Sign Payload'}
          </Button>
        </div>

        <div className={`glass-panel ${styles.viewer}`}>
          {!result ? (
            <div className={styles.emptyState}>
              <div className={styles.emptyIcon}>✍️</div>
              <p>Awaiting signature request</p>
            </div>
          ) : (
            <div style={{ padding: '1.5rem', wordBreak: 'break-all' }}>
              {result.success ? (
                <>
                  <h3 style={{ color: 'var(--accent-color)' }}>Signature Generated</h3>
                  <div style={{ marginTop: '1rem', padding: '1rem', background: 'rgba(0,0,0,0.2)', borderRadius: '8px' }}>
                    <strong>Signature:</strong>
                    <p style={{ color: 'var(--text-secondary)', fontFamily: 'monospace', marginTop: '0.5rem' }}>{result.signature}</p>
                  </div>
                  <div style={{ marginTop: '1rem', color: 'var(--success-color)' }}>
                    ✓ Deterministic JSON Verified
                  </div>
                </>
              ) : (
                <div style={{ color: 'var(--error-color)' }}>
                  <h3>Failed to Sign</h3>
                  <p>{result.error}</p>
                </div>
              )}
            </div>
          )}
        </div>
      </div>
    </div>
  );
}
