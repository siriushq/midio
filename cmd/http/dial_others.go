//go:build !linux
// +build !linux

package http

import (
	"context"
	"net"
	"syscall"
	"time"
)

// TODO: if possible implement for non-linux platforms, not a priority at the moment
//
//nolint:deadcode
func setInternalTCPParameters(c syscall.RawConn) error {
	return nil
}

// DialContext is a function to make custom Dial for internode communications
type DialContext func(ctx context.Context, network, address string) (net.Conn, error)

// NewInternodeDialContext setups a custom dialer for internode communication
var NewInternodeDialContext = NewCustomDialContext

// NewCustomDialContext configures a custom dialer for internode communications
func NewCustomDialContext(dialTimeout time.Duration) DialContext {
	return func(ctx context.Context, network, addr string) (net.Conn, error) {
		dialer := &net.Dialer{
			Timeout: dialTimeout,
		}
		return dialer.DialContext(ctx, network, addr)
	}
}
