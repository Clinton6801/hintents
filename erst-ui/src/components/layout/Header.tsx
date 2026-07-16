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
import React from 'react';
import styles from './Header.module.css';

export default function Header() {
  return (
    <header className={`glass-panel ${styles.header}`}>
      <div className={styles.searchContainer}>
        <div className={styles.searchIcon}>🔍</div>
        <input 
          type="text" 
          className={styles.searchInput} 
          placeholder="Search by Tx Hash / Contract ID / Ledger Sequence..."
        />
        <div className={styles.searchShortcut}>⌘K</div>
      </div>
      
      <div className={styles.actions}>
        <div className={styles.networkSelector}>
          <span className={styles.networkIndicator}></span>
          Testnet
        </div>
        <div className={styles.profile}>
          <div className={styles.avatar}></div>
        </div>
      </div>
    </header>
  );
}
