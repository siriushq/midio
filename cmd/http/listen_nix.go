//go:build linux || darwin || dragonfly || freebsd || netbsd || openbsd || rumprun
// +build linux darwin dragonfly freebsd netbsd openbsd rumprun

package http

import (
	"net"

	"github.com/valyala/tcplisten"
)

var cfg = &tcplisten.Config{
	DeferAccept: true,
	FastOpen:    true,
	// Bump up the soMaxConn value from 128 to 4096 to
	// handle large incoming concurrent requests.
	Backlog: 4096,
}

// Unix listener with special TCP options.
var listen = cfg.NewListener
var fallbackListen = net.Listen
