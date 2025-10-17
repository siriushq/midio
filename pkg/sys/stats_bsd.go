//go:build openbsd || freebsd || dragonfly
// +build openbsd freebsd dragonfly

package sys

import (
	"encoding/binary"
	"syscall"
)

func getHwPhysmem() (uint64, error) {
	totalString, err := syscall.Sysctl("hw.physmem")
	if err != nil {
		return 0, err
	}

	// syscall.sysctl() helpfully assumes the result is a null-terminated string and
	// removes the last byte of the result if it's 0 :/
	totalString += "\x00"

	total := uint64(binary.LittleEndian.Uint64([]byte(totalString)))

	return total, nil
}

// GetStats - return system statistics for bsd.
func GetStats() (stats Stats, err error) {
	stats.TotalRAM, err = getHwPhysmem()
	return stats, err
}
