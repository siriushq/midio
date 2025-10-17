//go:build linux
// +build linux

package certs

import "github.com/rjeczalik/notify"

var (
	// eventWrite contains the notify events that will cause a write
	eventWrite = []notify.Event{notify.InCloseWrite}
)
