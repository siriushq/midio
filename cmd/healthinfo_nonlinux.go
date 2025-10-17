//go:build !linux
// +build !linux

package cmd

import (
	"context"
	"net/http"
	"runtime"

	"github.com/siriushq/midio/pkg/madmin"
)

func getLocalDiskHwInfo(ctx context.Context, r *http.Request) madmin.ServerDiskHwInfo {
	addr := r.Host
	if globalIsDistErasure {
		addr = globalLocalNodeName
	}

	return madmin.ServerDiskHwInfo{
		Addr:  addr,
		Error: "unsupported platform: " + runtime.GOOS,
	}
}

func getLocalOsInfo(ctx context.Context, r *http.Request) madmin.ServerOsInfo {
	addr := r.Host
	if globalIsDistErasure {
		addr = globalLocalNodeName
	}

	return madmin.ServerOsInfo{
		Addr:  addr,
		Error: "unsupported platform: " + runtime.GOOS,
	}
}
