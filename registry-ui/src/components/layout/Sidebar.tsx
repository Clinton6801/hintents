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
import styles from './Sidebar.module.css';

export default function Sidebar() {
  return (
    <aside className={`glass-panel ${styles.sidebar}`}>
      <div className={styles.logoContainer}>
        <div className={styles.logo}></div>
        <h1 className="heading-gradient">Erst<span className={styles.version}>v2.0</span></h1>
      </div>
      
      <nav className={styles.nav}>
        <ul className={styles.navList}>
          <li className={styles.navItem}>
            <a href="/" className={`${styles.navLink} ${styles.active}`}>
              <span className={styles.icon}>⌘</span>
              Registry
            </a>
          </li>
          <li className={styles.navItem}>
            <a href="/debug" className={styles.navLink}>
              <span className={styles.icon}>◇</span>
              Debugger
            </a>
          </li>
          <li className={styles.navItem}>
            <a href="/trace" className={styles.navLink}>
              <span className={styles.icon}>⑆</span>
              Trace Viewer
            </a>
          </li>
          <li className={styles.navItem}>
            <a href="/profile" className={styles.navLink}>
              <span className={styles.icon}>🔥</span>
              Flamegraphs
            </a>
          </li>
          <li className={styles.navItem}>
            <a href="/audit" className={styles.navLink}>
              <span className={styles.icon}>✍️</span>
              Audit Logs
            </a>
          </li>
        </ul>
      </nav>
      
      <div className={styles.footer}>
        <div className={styles.statusDot}></div>
        <span className="text-mono">Daemon Connected</span>
      </div>
    </aside>
  );
}
