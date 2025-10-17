//go:build !linux
// +build !linux

package cmd

// Fallocate is not POSIX and not supported under Windows
// Always return successful
func Fallocate(fd int, offset int64, len int64) error {
	return nil
}
