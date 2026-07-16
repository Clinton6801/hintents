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



export interface ContractDependency {
  name: string;
  version?: string;
}

export interface ContractDeployment {
  name: string;
  wasm: string;
  salt?: string;
  initArgs?: string[];
  dependencies?: string[];
}

export interface DeploymentManifest {
  version: string;
  network: string;
  contracts: ContractDeployment[];
}

export interface DeployedContract {
  name: string;
  id: string;
  address: string;
}

export interface DeploymentResult {
  contractName: string;
  contractId: string;
  success: boolean;
  error?: string;
}

export interface DeploymentPlan {
  order: ContractDeployment[];
  resolved: Map<string, string>;
}
