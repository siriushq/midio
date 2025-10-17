//go:build !windows
// +build !windows

package disk

import (
	"syscall"
)

// SameDisk reports whether di1 and di2 describe the same disk.
func SameDisk(disk1, disk2 string) (bool, error) {
	st1 := syscall.Stat_t{}
	st2 := syscall.Stat_t{}

	if err := syscall.Stat(disk1, &st1); err != nil {
		return false, err
	}

	if err := syscall.Stat(disk2, &st2); err != nil {
		return false, err
	}

	return st1.Dev == st2.Dev, nil
}
