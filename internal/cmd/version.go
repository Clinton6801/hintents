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

package cmd

import (
	"encoding/json"
	"fmt"
	"runtime/debug"
	"time"

	"github.com/dotandev/hintents/internal/version"
	"github.com/spf13/cobra"
)

type VersionInfo struct {
	Version   string `json:"version"`
	CommitSHA string `json:"commit_sha"`
	BuildDate string `json:"build_date"`
	GoVersion string `json:"go_version"`
}

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:     "version",
	GroupID: "utility",
	Short:   "Show version information",
	Long:    "Display detailed build information including version, commit hash, and build date",
	Args:    cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		jsonOutput, _ := cmd.Flags().GetBool("json")

		info := getVersionInfo()

		if jsonOutput {
			output, _ := json.MarshalIndent(info, "", "  ")
			fmt.Println(string(output))
		} else {
			fmt.Printf("Erst Version: %s
", info.Version)
			fmt.Printf("Commit SHA:   %s
", info.CommitSHA)
			fmt.Printf("Build Date:   %s
", info.BuildDate)
			fmt.Printf("Go Version:   %s
", info.GoVersion)
		}
		fmt.Printf("erst version %s
", version.Version)
	},
}

func getVersionInfo() VersionInfo {
	info := VersionInfo{
		Version:   version.Version,
		CommitSHA: version.CommitSHA,
		BuildDate: version.BuildDate,
		GoVersion: "unknown",
	}

	// Use runtime/debug as fallback
	if buildInfo, ok := debug.ReadBuildInfo(); ok {
		info.GoVersion = buildInfo.GoVersion

		// Try to get commit and build time from build info if not set by ldflags
		for _, setting := range buildInfo.Settings {
			switch setting.Key {
			case "vcs.revision":
				if info.CommitSHA == "unknown" {
					info.CommitSHA = setting.Value
				}
			case "vcs.time":
				if info.BuildDate == "unknown" {
					if t, err := time.Parse(time.RFC3339, setting.Value); err == nil {
						info.BuildDate = t.Format("2006-01-02 15:04:05 UTC")
					}
				}
			}
		}
	}

	return info
}

func init() {
	rootCmd.AddCommand(versionCmd)
	versionCmd.Flags().Bool("json", false, "Output version information in JSON format")
}
