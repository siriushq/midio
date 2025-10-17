//go:build !linux
// +build !linux

package sys

import "errors"

// GetMaxThreads returns the maximum number of threads that the system can create.
func GetMaxThreads() (int, error) {
	return 0, errors.New("getting max threads is not supported")
}
