//go:build !linux
// +build !linux

// SPDX-License-Identifier: BSD-3-Clause AND Apache-2.0
package certs

import "github.com/rjeczalik/notify"

var (
	// eventWrite contains the notify events that will cause a write
	eventWrite = []notify.Event{notify.Create, notify.Write}
)
