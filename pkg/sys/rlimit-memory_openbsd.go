//go:build openbsd
// +build openbsd

package sys

import "syscall"

// GetMaxMemoryLimit - returns the maximum size of the process's virtual memory (address space) in bytes.
func GetMaxMemoryLimit() (curLimit, maxLimit uint64, err error) {
	var rlimit syscall.Rlimit
	if err = syscall.Getrlimit(syscall.RLIMIT_DATA, &rlimit); err == nil {
		curLimit = rlimit.Cur
		maxLimit = rlimit.Max
	}

	return curLimit, maxLimit, err
}

// SetMaxMemoryLimit - sets the maximum size of the process's virtual memory (address space) in bytes.
func SetMaxMemoryLimit(curLimit, maxLimit uint64) error {
	rlimit := syscall.Rlimit{Cur: curLimit, Max: maxLimit}
	return syscall.Setrlimit(syscall.RLIMIT_DATA, &rlimit)
}
