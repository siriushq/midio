//go:build linux
// +build linux

// SPDX-License-Identifier: BSD-3-Clause AND Apache-2.0
package mountinfo

// mountInfo - This represents a single line in /proc/mounts.
type mountInfo struct {
	Device  string
	Path    string
	FSType  string
	Options []string
	Freq    string
	Pass    string
}

func (m mountInfo) String() string {
	return m.Path
}

// mountInfos - This represents the entire /proc/mounts.
type mountInfos []mountInfo
