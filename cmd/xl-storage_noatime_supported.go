//go:build !windows && !darwin && !freebsd
// +build !windows,!darwin,!freebsd

package cmd

import (
	"os"
)

var (
	// Disallow updating access times
	readMode = os.O_RDONLY | 0x40000 // O_NOATIME

	// Write with data sync only used only for `xl.meta` writes
	writeMode = 0x1000 // O_DSYNC
)
