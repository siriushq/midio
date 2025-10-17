//go:build freebsd || dragonfly
// +build freebsd dragonfly

package sys

import "syscall"

// GetMaxMemoryLimit - returns the maximum size of the process's virtual memory (address space) in bytes.
func GetMaxMemoryLimit() (curLimit, maxLimit uint64, err error) {
	var rlimit syscall.Rlimit
	if err = syscall.Getrlimit(syscall.RLIMIT_DATA, &rlimit); err == nil {
		curLimit = uint64(rlimit.Cur)
		maxLimit = uint64(rlimit.Max)
	}

	return curLimit, maxLimit, err
}

// SetMaxMemoryLimit - sets the maximum size of the process's virtual memory (address space) in bytes.
func SetMaxMemoryLimit(curLimit, maxLimit uint64) error {
	rlimit := syscall.Rlimit{Cur: int64(curLimit), Max: int64(maxLimit)}
	return syscall.Setrlimit(syscall.RLIMIT_DATA, &rlimit)
}
