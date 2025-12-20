// SPDX-License-Identifier: BSD-3-Clause AND MIT

//go:build arm64 && !linux && !darwin
// +build arm64,!linux,!darwin

package cpuid

import "runtime"

func detectOS(c *CPUInfo) bool {
	c.PhysicalCores = runtime.NumCPU()
	// For now assuming 1 thread per core...
	c.ThreadsPerCore = 1
	c.LogicalCores = c.PhysicalCores
	return false
}
