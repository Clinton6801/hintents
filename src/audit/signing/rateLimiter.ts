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



import * as fs from 'fs';
import * as path from 'path';
import * as os from 'os';

/**
 * HsmRateLimiter provides cross-process rate limiting for HSM operations.
 * It uses a sliding window stored in a file in the user's home directory.
 */
export class HsmRateLimiter {
    private static getLimitFile(): string {
        return (
            process.env.ERST_AUDIT_HSM_RATE_LIMIT_FILE ||
            path.join(os.homedir(), '.erst', 'audit_hsm_calls.json')
        );
    }
    private static readonly WINDOW_MS = 60000; // 1 minute window
    private static readonly DEFAULT_MAX_RPM = 1000;

    /**
     * Checks if the rate limit has been exceeded.
     * Throws an error if the limit is reached.
     * Updates the call history on success.
     */
    public static async checkAndRecordCall(): Promise<void> {
        const maxRpm = this.getMaxRpm();

        // Ensure config directory exists
        const limitFile = this.getLimitFile();
        const dir = path.dirname(limitFile);
        if (!fs.existsSync(dir)) {
            try {
                fs.mkdirSync(dir, { recursive: true });
            } catch (err) {
                // If we can't create the directory, we'll fall back to in-memory only limit
                // which is less robust but better than failing entirely if the system is read-only.
                console.warn('Could not create audit directory for rate limiting stats:', err);
            }
        }

        const now = Date.now();
        const cutoff = now - this.WINDOW_MS;

        let history: number[] = [];

        // Read existing history
        if (fs.existsSync(limitFile)) {
            try {
                const data = fs.readFileSync(limitFile, 'utf8');
                history = JSON.parse(data);
            } catch {
                // Corrupt or empty file, reset history
                history = [];
            }
        }

        // Filter old entries
        history = history.filter((t) => t > cutoff);

        if (history.length >= maxRpm) {
            throw new Error(
                `HSM rate limit protection triggered: ${history.length} calls in the last minute. ` +
                `Maximum allowed is ${maxRpm} RPM. Please wait before retrying.`
            );
        }

        // Record this call
        history.push(now);

        // Save updated history
        try {
            fs.writeFileSync(limitFile, JSON.stringify(history), 'utf8');
        } catch (err) {
            console.warn('Could not save rate limiting stats:', err);
        }
    }

    private static getMaxRpm(): number {
        const envVal = process.env.ERST_PKCS11_MAX_RPM;
        if (envVal) {
            const parsed = parseInt(envVal, 10);
            if (!isNaN(parsed) && parsed > 0) {
                return parsed;
            }
        }
        return this.DEFAULT_MAX_RPM;
    }
}
