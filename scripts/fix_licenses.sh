#!/bin/bash

# Standard Apache 2.0 Header
HEADER="// SPDX-License-Identifier: Apache-2.0
// Copyright (c) 2026 dotandev
//
// Licensed under the Apache License, Version 2.0 (the \"License\");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an \"AS IS\" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License."

echo "Starting Apache 2.0 license fix..."

# Find all relevant source files excluding node_modules and target directories
find . -type f \( -name "*.go" -o -name "*.ts" -o -name "*.tsx" -o -name "*.rs" \) \
  -not -path "*/node_modules/*" -not -path "*/.next/*" -not -path "*/target/*" | while read -r FILE; do
  
  if ! grep -q "SPDX-License-Identifier: Apache-2.0" "$FILE"; then
      echo "Adding Apache 2.0 header to $FILE"
      
      # Remove any existing MIT or old SPDX headers at the top
      sed -i.bak '/SPDX-License-Identifier:/d' "$FILE"
      sed -i.bak '/Copyright (c)/d' "$FILE"
      sed -i.bak '/Licensed under the Apache License/d' "$FILE"
      sed -i.bak '/http:\/\/www.apache.org\/licenses\/LICENSE-2.0/d' "$FILE"
      sed -i.bak '/Unless required by applicable law/d' "$FILE"
      sed -i.bak '/distributed under the License/d' "$FILE"
      sed -i.bak '/WITHOUT WARRANTIES OR CONDITIONS/d' "$FILE"
      sed -i.bak '/See the License for the specific language/d' "$FILE"
      sed -i.bak '/limitations under the License./d' "$FILE"
      rm -f "${FILE}.bak"

      # Prepend the new header
      echo -e "$HEADER\n\n$(cat "$FILE")" > "$FILE"
  fi
done

echo "All files processed for Apache 2.0 license compliance."