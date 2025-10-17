//go:build windows
// +build windows

package sys

// GetMaxMemoryLimit - returns the maximum size of the process's virtual memory (address space) in bytes.
func GetMaxMemoryLimit() (curLimit, maxLimit uint64, err error) {
	// Nothing to do for windows.
	return curLimit, maxLimit, err
}

// SetMaxMemoryLimit - sets the maximum size of the process's virtual memory (address space) in bytes.
func SetMaxMemoryLimit(curLimit, maxLimit uint64) error {
	// Nothing to do for windows.
	return nil
}
