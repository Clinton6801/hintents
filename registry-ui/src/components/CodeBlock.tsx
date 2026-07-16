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

export default function CodeBlock({ code }: { code: string }) {
  const [copied, setCopied] = useState(false);

  const copyToClipboard = () => {
    navigator.clipboard.writeText(code);
    setCopied(true);
    setTimeout(() => setCopied(false), 2000);
  };

  return (
    <div style={{
      position: 'relative',
      background: 'rgba(30, 27, 60, 0.4)',
      border: '1px solid var(--border-subtle)',
      borderRadius: '8px',
      padding: '1rem',
      display: 'flex',
      alignItems: 'center',
      justifyContent: 'space-between',
      margin: '1rem 0'
    }}>
      <pre style={{ margin: 0, fontFamily: 'var(--font-mono)', fontSize: '0.9rem', color: '#e2e8f0', overflowX: 'auto' }}>
        <code>{code}</code>
      </pre>
      <button 
        onClick={copyToClipboard}
        style={{
          background: 'transparent',
          border: 'none',
          color: 'var(--text-secondary)',
          cursor: 'pointer',
          padding: '0.5rem',
          marginLeft: '1rem'
        }}
        title="Copy to clipboard"
      >
        {copied ? '✓' : '📋'}
      </button>
    </div>
  );
}
