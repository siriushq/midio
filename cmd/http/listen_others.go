//go:build windows || plan9 || solaris
// +build windows plan9 solaris

// SPDX-License-Identifier: BSD-3-Clause AND Apache-2.0
package http

import "net"

// Windows, plan9 specific listener.
var listen = net.Listen
var fallbackListen = net.Listen
