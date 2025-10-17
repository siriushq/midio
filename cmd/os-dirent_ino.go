//go:build (linux || darwin) && !appengine
// +build linux darwin
// +build !appengine

package cmd

import "syscall"

func direntInode(dirent *syscall.Dirent) uint64 {
	return dirent.Ino
}
