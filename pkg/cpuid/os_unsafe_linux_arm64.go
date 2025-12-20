// SPDX-License-Identifier: BSD-3-Clause AND MIT

//go:build !nounsafe
// +build !nounsafe

package cpuid

import _ "unsafe" // needed for go:linkname

//go:linkname hwcap internal/cpu.HWCap
var hwcap uint
