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

export default function PackageDetails({ params }: { params: { id: string } }) {
  return (
    <div className={styles.container}>
      <header className={styles.header}>
        <div className={styles.headerTitle}>
          <div className={styles.avatar}></div>
          <div>
            <h1 className="heading-gradient">{params.id || 'soroban-token'}</h1>
            <p className={styles.subtitle}>Verified Soroban Smart Contract • v1.0.4</p>
          </div>
        </div>
        <div className={styles.actions}>
          <Button variant="primary">Install Package</Button>
          <Button variant="secondary">View Source</Button>
        </div>
      </header>

      <div className={styles.grid}>
        <div className={styles.mainContent}>
          <section className={`glass-panel ${styles.panel}`}>
            <h2>Contract ABI (Functions)</h2>
            <div className={styles.abiList}>
              <div className={styles.abiItem}>
                <span className={styles.abiName}>initialize</span>
                <span className={styles.abiArgs}>(admin: Address, decimal: u32, name: String)</span>
              </div>
              <div className={styles.abiItem}>
                <span className={styles.abiName}>mint</span>
                <span className={styles.abiArgs}>(to: Address, amount: i128)</span>
              </div>
              <div className={styles.abiItem}>
                <span className={styles.abiName}>transfer</span>
                <span className={styles.abiArgs}>(from: Address, to: Address, amount: i128)</span>
              </div>
            </div>
          </section>

          <section className={`glass-panel ${styles.panel}`}>
            <h2>Dependency Graph</h2>
            <div className={styles.graphPlaceholder}>
              <div className={styles.emptyIcon}>⑆</div>
              <p>Dependency graph visualization generated</p>
            </div>
          </section>
        </div>

        <div className={styles.sideContent}>
          <section className={`glass-card ${styles.metadataPanel}`}>
            <h3>Metadata</h3>
            <div className={styles.metaRow}>
              <span className={styles.metaLabel}>Contract ID</span>
              <span className={styles.metaValue}>CBA3...9F21</span>
            </div>
            <div className={styles.metaRow}>
              <span className={styles.metaLabel}>Publisher</span>
              <span className={styles.metaValue}>Stellar Development Foundation</span>
            </div>
            <div className={styles.metaRow}>
              <span className={styles.metaLabel}>License</span>
              <span className={styles.metaValue}>Apache-2.0</span>
            </div>
            <div className={styles.metaRow}>
              <span className={styles.metaLabel}>Total Invokes</span>
              <span className={styles.metaValue}>1.4M</span>
            </div>
          </section>
        </div>
      </div>
    </div>
  );
}
