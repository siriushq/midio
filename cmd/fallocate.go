//go:build !linux
// +build !linux

// SPDX-License-Identifier: BSD-3-Clause AND Apache-2.0
package cmd

// Fallocate is not POSIX and not supported under Windows
// Always return successful
func Fallocate(fd int, offset int64, len int64) error {
	return nil
}
