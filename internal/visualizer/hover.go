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

package visualizer

import (
	"bytes"
	_ "embed"
	"strings"
	"text/template"
)

//go:embed hover.tmpl
var hoverTemplateSource string

var hoverTemplate = template.Must(template.New("hover").Parse(hoverTemplateSource))

// hoverContentData is the data passed to hover.tmpl when rendering hover content.
type hoverContentData struct {
	Name        string
	Description string
}

// HostFunctionHoverContent builds markdown-friendly hover content for the given host function.
//
// The presentation (markdown structure) lives in hover.tmpl, embedded at build time so
// documentation formatting can be updated without touching Go code.
func HostFunctionHoverContent(name string) string {
	data := hoverContentData{
		Name:        name,
		Description: DescribeHostFunction(name),
	}

	var buf bytes.Buffer
	if err := hoverTemplate.Execute(&buf, data); err != nil {
		return "**" + name + "**

" + data.Description
	}

	return strings.TrimRight(buf.String(), "
")
}
