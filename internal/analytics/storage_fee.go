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


// Copyright 2026 Erst Users

package analytics

type StorageFeeModel struct {
	FeePerByte uint64 // protocol-defined
}

func CalculateStorageFee(deltaBytes int64, model StorageFeeModel) int64 {
	if deltaBytes <= 0 {
		return 0
	}
	return deltaBytes * int64(model.FeePerByte)
}
