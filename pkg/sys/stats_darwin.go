//go:build darwin
// +build darwin

package sys

import (
	"encoding/binary"
	"syscall"
)

func getHwMemsize() (uint64, error) {
	totalString, err := syscall.Sysctl("hw.memsize")
	if err != nil {
		return 0, err
	}

	// syscall.sysctl() helpfully assumes the result is a null-terminated string and
	// removes the last byte of the result if it's 0 :/
	totalString += "\x00"

	return binary.LittleEndian.Uint64([]byte(totalString)), nil
}

// GetStats - return system statistics for macOS.
func GetStats() (stats Stats, err error) {
	stats.TotalRAM, err = getHwMemsize()
	return stats, err
}
