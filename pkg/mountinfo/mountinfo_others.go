//go:build !linux && !windows
// +build !linux,!windows

// SPDX-License-Identifier: BSD-3-Clause AND Apache-2.0
package mountinfo

// CheckCrossDevice - check if any input path has multiple sub-mounts.
// this is a dummy function and returns nil for now.
func CheckCrossDevice(paths []string) error {
	return nil
}

// IsLikelyMountPoint determines if a directory is a mountpoint.
func IsLikelyMountPoint(file string) bool {
	return false
}
