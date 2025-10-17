//go:build windows
// +build windows

package disk

// SameDisk reports whether di1 and di2 describe the same disk.
func SameDisk(disk1, disk2 string) (bool, error) {
	return false, nil
}
