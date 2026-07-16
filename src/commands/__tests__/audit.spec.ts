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



import { Command } from 'commander';
import { registerAuditCommands } from '../audit';
import * as fs from 'fs';
import { verifyAuditLog } from '../../audit/AuditVerifier';

jest.mock('../../audit/AuditVerifier');
jest.mock('fs');

describe('Audit Commands CLI', () => {
    let program: Command;

    beforeEach(() => {
        program = new Command();
        registerAuditCommands(program);
        jest.clearAllMocks();
    });

    describe('audit:verify', () => {
        it('should verify an audit log from a file', async () => {
            const mockLog = { trace: { foo: 'bar' }, signature: 'abc', publicKey: 'pub', hash: '123' };
            (fs.readFileSync as jest.Mock).mockReturnValue(JSON.stringify(mockLog));
            (verifyAuditLog as jest.Mock).mockReturnValue(true);

            const consoleLogSpy = jest.spyOn(console, 'log').mockImplementation();

            await program.parseAsync(['node', 'test', 'audit:verify', '--file', 'test.json']);

            expect(fs.readFileSync).toHaveBeenCalledWith('test.json', 'utf8');
            expect(verifyAuditLog).toHaveBeenCalledWith(mockLog);
            expect(consoleLogSpy).toHaveBeenCalledWith(expect.stringContaining('[OK] Verification successful'));

            consoleLogSpy.mockRestore();
        });

        it('should verify from individual components', async () => {
            const payload = JSON.stringify({ amount: 100 });
            const sig = 'deadbeef';
            const pubkey = 'pem-content';

            (verifyAuditLog as jest.Mock).mockReturnValue(true);
            const consoleLogSpy = jest.spyOn(console, 'log').mockImplementation();

            await program.parseAsync([
                'node', 'test', 'audit:verify',
                '--payload', payload,
                '--sig', sig,
                '--pubkey', pubkey
            ]);

            expect(verifyAuditLog).toHaveBeenCalledWith(expect.objectContaining({
                trace: { amount: 100 },
                signature: sig,
                publicKey: pubkey
            }));
            expect(consoleLogSpy).toHaveBeenCalledWith(expect.stringContaining('[OK] Verification successful'));

            consoleLogSpy.mockRestore();
        });

        it('should fail if signature is invalid', async () => {
            const mockLog = { trace: { foo: 'bar' }, signature: 'bad', publicKey: 'pub', hash: '123' };
            (fs.readFileSync as jest.Mock).mockReturnValue(JSON.stringify(mockLog));
            (verifyAuditLog as jest.Mock).mockReturnValue(false);

            const consoleErrorSpy = jest.spyOn(console, 'error').mockImplementation();
            const processExitSpy = jest.spyOn(process, 'exit').mockImplementation((() => { }) as any);

            await program.parseAsync(['node', 'test', 'audit:verify', '--file', 'test.json']);

            expect(verifyAuditLog).toHaveBeenCalled();
            expect(consoleErrorSpy).toHaveBeenCalledWith(expect.stringContaining('[FAIL] Verification failed'));
            expect(processExitSpy).toHaveBeenCalledWith(1);

            consoleErrorSpy.mockRestore();
            processExitSpy.mockRestore();
        });

        it('should throw error if missing arguments', async () => {
            const consoleErrorSpy = jest.spyOn(console, 'error').mockImplementation();
            const processExitSpy = jest.spyOn(process, 'exit').mockImplementation((() => { }) as any);

            await program.parseAsync(['node', 'test', 'audit:verify', '--payload', '{}']);

            expect(consoleErrorSpy).toHaveBeenCalledWith(expect.stringContaining('You must provide either --file or all of'));
            expect(processExitSpy).toHaveBeenCalledWith(1);

            consoleErrorSpy.mockRestore();
            processExitSpy.mockRestore();
        });
    });
});
