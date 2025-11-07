//go:build windows || darwin || freebsd
// +build windows darwin freebsd

// SPDX-License-Identifier: BSD-3-Clause AND Apache-2.0
package cmd

import (
	"os"
)

var (
	// No special option for reads on windows
	readMode = os.O_RDONLY

	// Write with sync no buffering only used only for `xl.meta` writes
	writeMode = os.O_SYNC
)
