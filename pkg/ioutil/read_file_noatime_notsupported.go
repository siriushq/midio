//go:build windows || darwin || freebsd
// +build windows darwin freebsd

// SPDX-License-Identifier: BSD-3-Clause AND Apache-2.0
package ioutil

import "os"

var readMode = os.O_RDONLY
