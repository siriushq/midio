//go:build darwin || freebsd || openbsd || netbsd
// +build darwin freebsd openbsd netbsd

// SPDX-License-Identifier: BSD-3-Clause AND Apache-2.0
package cmd

import "syscall"

func direntNamlen(dirent *syscall.Dirent) (uint64, error) {
	return uint64(dirent.Namlen), nil
}
