//go:build windows
// +build windows

package sys

// GetMaxOpenFileLimit - returns maximum file descriptor number that can be opened by this process.
func GetMaxOpenFileLimit() (curLimit, maxLimit uint64, err error) {
	// Nothing to do for windows.
	return curLimit, maxLimit, err
}

// SetMaxOpenFileLimit - sets maximum file descriptor number that can be opened by this process.
func SetMaxOpenFileLimit(curLimit, maxLimit uint64) error {
	// Nothing to do for windows.
	return nil
}
