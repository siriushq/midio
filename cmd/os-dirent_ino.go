//go:build (linux || darwin) && !appengine
// +build linux darwin
// +build !appengine

// SPDX-License-Identifier: BSD-3-Clause AND Apache-2.0
package cmd

import "syscall"

func direntInode(dirent *syscall.Dirent) uint64 {
	return dirent.Ino
}
