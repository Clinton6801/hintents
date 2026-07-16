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
import CodeBlock from '@/components/CodeBlock';
import styles from './page.module.css';

export default function Home() {
  return (
    <div style={{ position: 'relative', minHeight: '100vh', overflow: 'hidden' }}>
      {/* MVR Background Blurs */}
      <div style={{
        position: 'absolute',
        top: '-20%',
        left: '10%',
        width: '80%',
        height: '600px',
        background: 'var(--gradient-mvr-blur)',
        filter: 'blur(150px)',
        opacity: 0.15,
        zIndex: -1,
        borderRadius: '50%',
        pointerEvents: 'none'
      }}></div>

      <div style={{
        position: 'absolute',
        bottom: '-10%',
        right: '-10%',
        width: '600px',
        height: '600px',
        background: 'var(--gradient-mvr-blur)',
        filter: 'blur(200px)',
        opacity: 0.2,
        zIndex: -1,
        borderRadius: '50%',
        pointerEvents: 'none'
      }}></div>

      {/* Header & Search */}
      <div style={{ maxWidth: '800px', margin: '0 auto', paddingTop: '8rem', textAlign: 'center', padding: '8rem 2rem 4rem 2rem' }}>
        <h1 style={{ 
          fontSize: '3.5rem', 
          fontWeight: 800, 
          letterSpacing: '-0.02em',
          marginBottom: '1rem',
          lineHeight: 1.1
        }}>
          Move Package Registry
        </h1>
        <p style={{ 
          fontSize: '1.25rem', 
          color: 'var(--text-secondary)',
          lineHeight: 1.5,
          marginBottom: '3rem'
        }}>
          MVR is the central hub for discovering, sharing, and managing Move packages on the Soroban blockchain. Build secure, scalable, and innovative decentralized applications with the power of Soroban.
        </p>

        {/* Search Bar */}
        <div style={{
          position: 'relative',
          padding: '2px',
          background: 'var(--gradient-mvr)',
          borderRadius: '9999px',
          zIndex: 30
        }}>
          <div style={{
            display: 'flex',
            alignItems: 'center',
            background: 'var(--bg-surface)',
            borderRadius: '9999px',
            padding: '0.75rem 1.5rem',
            width: '100%'
          }}>
            <span style={{ opacity: 0.5, marginRight: '0.75rem' }}>🔍</span>
            <input 
              type="text"
              placeholder="Search package..."
              style={{
                background: 'transparent',
                border: 'none',
                outline: 'none',
                color: 'var(--text-primary)',
                fontSize: '1rem',
                width: '100%',
                padding: '0.5rem 0'
              }}
            />
          </div>
        </div>
      </div>

      {/* Docs Banner */}
      <div style={{ maxWidth: '1200px', margin: '0 auto', padding: '0 2rem' }}>
        <div className="glass-panel" style={{ 
          display: 'flex', 
          justifyContent: 'space-between', 
          alignItems: 'center',
          padding: '2.5rem',
          position: 'relative',
          overflow: 'hidden'
        }}>
          <div style={{ position: 'relative', zIndex: 10 }}>
            <h2 style={{ fontSize: '2.25rem', fontWeight: 700, marginBottom: '0.5rem' }}>Move onto MVR</h2>
            <p style={{ fontSize: '1.25rem', color: 'var(--text-secondary)' }}>Bring your package to the future of Move.</p>
          </div>
          <div style={{ display: 'flex', gap: '1rem', position: 'relative', zIndex: 10 }}>
            <button style={{
              background: 'rgba(255,255,255,0.1)',
              border: 'none',
              padding: '0.75rem 1.5rem',
              color: 'white',
              borderRadius: '6px',
              fontWeight: 500,
              cursor: 'pointer'
            }}>Register your app</button>
            <button style={{
              background: 'var(--accent-purple)',
              border: 'none',
              padding: '0.75rem 1.5rem',
              color: 'white',
              borderRadius: '6px',
              fontWeight: 500,
              cursor: 'pointer'
            }}>View MVR Docs</button>
          </div>
        </div>
      </div>

      {/* Getting Started Steps */}
      <div style={{ maxWidth: '1000px', margin: '4rem auto', padding: '0 2rem' }}>
        <div className="glass-panel" style={{ padding: '3rem' }}>
          <div style={{ textAlign: 'center', marginBottom: '2rem' }}>
            <h2 style={{ fontSize: '1.5rem', fontWeight: 700, color: 'var(--accent-blue)' }}>Share your package on MVR</h2>
            <p style={{ color: 'var(--text-secondary)' }}>Ready to dive in? Follow these steps to start building on MVR today.</p>
          </div>

          <hr style={{ borderColor: 'var(--border-subtle)', margin: '2rem 0' }} />

          <div style={{ display: 'grid', gridTemplateColumns: '1fr 1fr', gap: '2rem', alignItems: 'center' }}>
            <div>
              <h3 style={{ fontSize: '1.25rem', fontWeight: 700, marginBottom: '0.5rem' }}>Set up MVR CLI</h3>
              <p style={{ color: 'var(--text-secondary)' }}>Install the MVR command line tool to interact with MVR.</p>
            </div>
            <CodeBlock code="cargo install --locked --git https://github.com/mystenlabs/mvr --branch release mvr" />
          </div>

          <div style={{ display: 'grid', gridTemplateColumns: '1fr 1fr', gap: '2rem', alignItems: 'center', marginTop: '2rem' }}>
            <div>
              <h3 style={{ fontSize: '1.25rem', fontWeight: 700, marginBottom: '0.5rem' }}>Resolve packages</h3>
              <p style={{ color: 'var(--text-secondary)' }}>Verify your configuration to ensure that your packages are resolving properly.</p>
            </div>
            <CodeBlock code="mvr resolve @soroban/core" />
          </div>

          <div style={{ display: 'grid', gridTemplateColumns: '1fr 1fr', gap: '2rem', alignItems: 'center', marginTop: '2rem' }}>
            <div>
              <h3 style={{ fontSize: '1.25rem', fontWeight: 700, marginBottom: '0.5rem' }}>Add Dependencies</h3>
              <p style={{ color: 'var(--text-secondary)' }}>Unlock seamless dependency management with just one simple command!</p>
            </div>
            <CodeBlock code="mvr add @soroban/core" />
          </div>

        </div>
      </div>
    </div>
  );
}
