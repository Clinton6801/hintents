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



import * as rpc from 'vscode-jsonrpc/node';
import * as net from 'net';

export interface TraceStep {
    step: number;
    timestamp: string;
    operation: string;
    contract_id?: string;
    function?: string;
    arguments?: any[];
    return_value?: any;
    error?: string;
    host_state?: any;
    memory?: any;
    cpu_delta?: number;
    memory_delta?: number;
}

export interface Trace {
    transaction_hash: string;
    start_time: string;
    states: TraceStep[];
}

export class ERSTClient {
    private connection: rpc.MessageConnection | undefined;

    constructor(private host: string = '127.0.0.1', private port: number = 8080) { }

    async connect(): Promise<void> {
        return new Promise((resolve, reject) => {
            const socket = net.createConnection({ host: this.host, port: this.port });
            socket.on('connect', () => {
                this.connection = rpc.createMessageConnection(
                    new rpc.StreamMessageReader(socket),
                    new rpc.StreamMessageWriter(socket)
                );
                this.connection.listen();
                resolve();
            });
            socket.on('error', (err) => {
                reject(err);
            });
        });
    }

    async debugTransaction(hash: string): Promise<any> {
        if (!this.connection) await this.connect();
        return this.connection!.sendRequest('DebugTransaction', { hash });
    }

    async getTrace(hash: string): Promise<Trace> {
        if (!this.connection) await this.connect();
        return this.connection!.sendRequest('GetTrace', { hash }) as Promise<Trace>;
    }

    dispose() {
        if (this.connection) {
            this.connection.dispose();
        }
    }
}
