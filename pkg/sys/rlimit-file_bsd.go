//go:build freebsd || dragonfly
// +build freebsd dragonfly

package sys

import (
	"syscall"
)

// GetMaxOpenFileLimit - returns maximum file descriptor number that can be opened by this process.
func GetMaxOpenFileLimit() (curLimit, maxLimit uint64, err error) {
	var rlimit syscall.Rlimit
	if err = syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rlimit); err == nil {
		curLimit = uint64(rlimit.Cur)
		maxLimit = uint64(rlimit.Max)
	}

	return curLimit, maxLimit, err
}

// SetMaxOpenFileLimit - sets maximum file descriptor number that can be opened by this process.
func SetMaxOpenFileLimit(curLimit, maxLimit uint64) error {
	rlimit := syscall.Rlimit{Cur: int64(curLimit), Max: int64(curLimit)}
	return syscall.Setrlimit(syscall.RLIMIT_NOFILE, &rlimit)
}
