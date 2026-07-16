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


import React from 'react';
import styles from './page.module.css';
import Button from '@/components/ui/Button';

export default function Home() {
  return (
    <div className={styles.dashboard}>
      <header className={styles.header}>
        <div>
          <h1 className="heading-gradient">Transactions Registry</h1>
          <p className={styles.subtitle}>High-Fidelity Soroban Replay Engine</p>
        </div>
        <div className={styles.actions}>
          <Button variant="primary">New Simulation</Button>
        </div>
      </header>

      <div className={styles.metricsGrid}>
        <div className="glass-card">
          <div className={styles.metricLabel}>Total Simulations</div>
          <div className={styles.metricValue}>1,248</div>
          <div className={styles.metricTrend}>+12% this week</div>
        </div>
        <div className="glass-card">
          <div className={styles.metricLabel}>Cache Hit Rate</div>
          <div className={styles.metricValue}>94.2%</div>
          <div className={styles.metricTrend}>LRU Eviction Active</div>
        </div>
        <div className="glass-card">
          <div className={styles.metricLabel}>Avg TTL</div>
          <div className={styles.metricValue}>14.2ms</div>
          <div className={styles.metricTrend}>Optimized</div>
        </div>
      </div>

      <section className={styles.recentTransactions}>
        <div className={styles.sectionHeader}>
          <h2 className="heading-gradient">Recent Traces</h2>
          <Button variant="ghost" size="sm">View All</Button>
        </div>
        
        <div className={`glass-panel ${styles.tableContainer}`}>
          <table className={styles.table}>
            <thead>
              <tr>
                <th>Tx Hash</th>
                <th>Contract</th>
                <th>Status</th>
                <th>Gas Cost</th>
                <th>Date</th>
              </tr>
            </thead>
            <tbody>
              {[
                { hash: 'e3b0c442...', contract: 'CCMZ...4F9A', status: 'Success', gas: '14,200', date: '2 mins ago' },
                { hash: '8d969eef...', contract: 'CDZ9...2L1X', status: 'Failed', gas: '42,100', date: '1 hour ago' },
                { hash: 'f2c97ggh...', contract: 'CCMZ...4F9A', status: 'Success', gas: '11,400', date: '3 hours ago' },
              ].map((tx, i) => (
                <tr key={i} className={styles.tableRow}>
                  <td className="text-mono">{tx.hash}</td>
                  <td className="text-mono">{tx.contract}</td>
                  <td>
                    <span className={`${styles.badge} ${tx.status === 'Success' ? styles.badgeSuccess : styles.badgeError}`}>
                      {tx.status}
                    </span>
                  </td>
                  <td className="text-mono">{tx.gas}</td>
                  <td className={styles.dateCell}>{tx.date}</td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>
      </section>
    </div>
  );
}
