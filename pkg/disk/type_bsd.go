//go:build darwin || freebsd || dragonfly || openbsd || solaris
// +build darwin freebsd dragonfly openbsd solaris

// SPDX-License-Identifier: BSD-3-Clause AND Apache-2.0
package disk

// getFSType returns the filesystem type of the underlying mounted filesystem
func getFSType(fstype []int8) string {
	b := make([]byte, len(fstype))
	for i, v := range fstype {
		b[i] = byte(v)
	}
	return string(b)
}
