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


import { NextResponse } from 'next/server';
import { exec } from 'child_process';
import { promisify } from 'util';
import path from 'path';

const execAsync = promisify(exec);

export async function POST(request: Request) {
  try {
    const { txHash, network } = await request.json();

    if (!txHash) {
      return NextResponse.json({ error: 'Transaction hash is required' }, { status: 400 });
    }

    const erstBinary = path.join(process.cwd(), '../erst');
    const netFlag = network || 'testnet';
    
    // Instead of launching the bubbletea UI which will hang the Node exec, 
    // we use a flag to output raw trace data or just get the standard debug output.
    // In a real scenario, this might be a different CLI subcommand that outputs JSON traces.
    const cmd = `${erstBinary} debug ${txHash} --network ${netFlag} --generate-trace --trace-output /dev/stdout`;
    
    console.log(`Executing: ${cmd}`);
    
    try {
      const { stdout, stderr } = await execAsync(cmd);
      
      let parsedOutput;
      try {
        parsedOutput = JSON.parse(stdout);
      } catch (e) {
        parsedOutput = stdout;
      }
      
      return NextResponse.json({
        success: true,
        traceData: parsedOutput,
        logs: stderr
      });
    } catch (execError: any) {
      return NextResponse.json({
        success: false,
        error: execError.message || 'Execution failed',
        stderr: execError.stderr || '',
        stdout: execError.stdout || ''
      }, { status: 500 });
    }
  } catch (err: any) {
    return NextResponse.json({ error: err.message }, { status: 500 });
  }
}
