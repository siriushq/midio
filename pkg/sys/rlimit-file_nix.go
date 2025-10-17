//go:build linux || darwin || openbsd || netbsd || solaris
// +build linux darwin openbsd netbsd solaris

package sys

import (
	"runtime"
	"syscall"
)

// GetMaxOpenFileLimit - returns maximum file descriptor number that can be opened by this process.
func GetMaxOpenFileLimit() (curLimit, maxLimit uint64, err error) {
	var rlimit syscall.Rlimit
	if err = syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rlimit); err == nil {
		curLimit = rlimit.Cur
		maxLimit = rlimit.Max
	}

	return curLimit, maxLimit, err
}

// SetMaxOpenFileLimit - sets maximum file descriptor number that can be opened by this process.
func SetMaxOpenFileLimit(curLimit, maxLimit uint64) error {
	if runtime.GOOS == "darwin" && curLimit > 10240 {
		// The max file limit is 10240, even though
		// the max returned by Getrlimit is 1<<63-1.
		// This is OPEN_MAX in sys/syslimits.h.
		// refer https://github.com/golang/go/issues/30401
		curLimit = 10240
	}
	rlimit := syscall.Rlimit{Cur: curLimit, Max: maxLimit}
	return syscall.Setrlimit(syscall.RLIMIT_NOFILE, &rlimit)
}
