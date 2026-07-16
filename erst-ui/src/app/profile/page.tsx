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

export default function ProfilePage() {
  return (
    <div className={styles.container}>
      <header className={styles.header}>
        <h1 className="heading-gradient">Flamegraph Profiler</h1>
        <p className={styles.subtitle}>Visualize CPU and Memory consumption for contract execution</p>
      </header>

      <div className={styles.grid}>
        <div className={`glass-card ${styles.configPanel}`}>
          <h3>Profiling Target</h3>
          <div className={styles.inputGroup}>
            <label>Transaction Hash</label>
            <input type="text" className={styles.input} placeholder="e.g. 0x..." />
          </div>
          
          <div className={styles.inputGroup}>
            <label>Format</label>
            <select className={styles.input}>
              <option>Interactive HTML</option>
              <option>Raw SVG</option>
            </select>
          </div>
          
          <Button variant="primary">Generate Flamegraph</Button>
        </div>

        <div className={`glass-panel ${styles.viewer}`}>
          <div className={styles.emptyState}>
            <div className={styles.emptyIcon}>🔥</div>
            <p>Ready to generate Flamegraph</p>
          </div>
        </div>
      </div>
    </div>
  );
}
