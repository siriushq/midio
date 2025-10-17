//go:build !linux && !windows
// +build !linux,!windows

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
